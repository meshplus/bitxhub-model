package pb

import (
	"bytes"
	"fmt"
	io "io"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/meshplus/bitxhub-kit/types"
	ethtypes "github.com/meshplus/eth-kit/types"
	"golang.org/x/crypto/sha3"
)

var _ Transaction = (*EthTransaction)(nil)

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

// deriveBufferPool holds temporary encoder buffers for DeriveSha and TX encoding.
var encodeBufferPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

var signer EIP155Signer

type EIP155Signer struct {
	chainId, chainIdMul *big.Int
}

func InitEIP155Signer(chainId *big.Int) {
	if chainId == nil {
		chainId = new(big.Int)
	}
	signer = EIP155Signer{
		chainId:    chainId,
		chainIdMul: new(big.Int).Mul(chainId, big.NewInt(2)),
	}
}

// Transaction is an Ethereum transaction.
type EthTransaction struct {
	inner ethtypes.TxData // Consensus contents of a transaction
	time  time.Time       // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type writeCounter int

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

func (e *EthTransaction) GetVersion() []byte {
	return nil
}

func (e *EthTransaction) GetInner() ethtypes.TxData {
	return e.inner
}

// Protected says whether the transaction is replay-protected.
func (e *EthTransaction) Protected() bool {
	switch tx := e.inner.(type) {
	case *ethtypes.LegacyTx:
		return tx.V != nil && isProtectedV(tx.V)
	default:
		return true
	}
}

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28 && v != 1 && v != 0
	}
	// anything not 27 or 28 is considered protected
	return true
}

func recoverPlain(sighash *types.Hash, R, S, Vb *big.Int, homestead bool) (*types.Address, error) {
	addr, err := ethtypes.RecoverPlain(sighash.Bytes(), R, S, Vb, homestead)
	if err != nil {
		return nil, err
	}
	return types.NewAddress(addr), nil
}

func (e *EthTransaction) GetFrom() *types.Address {
	if addr := e.from.Load(); addr != nil {
		return addr.(*types.Address)
	}

	addr, err := e.sender()
	if err != nil {
		return nil
	}
	e.from.Store(addr)

	return addr
}

func (e *EthTransaction) sender() (*types.Address, error) {
	V, R, S := e.GetRawSignature()
	switch e.GetType() {
	case ethtypes.LegacyTxType:
		if !e.Protected() {
			hash := ethtypes.RlpHash([]interface{}{
				e.GetNonce(),
				e.GetGasPrice(),
				e.GetGas(),
				e.inner.GetTo(),
				e.inner.GetValue(),
				e.GetPayload(),
			})
			addr, err := recoverPlain(types.NewHash(hash.Bytes()), R, S, V, true)
			if err != nil {
				return nil, fmt.Errorf("invalid signature")
			}
			return addr, nil
		}
		V = new(big.Int).Sub(V, signer.chainIdMul)
		V.Sub(V, big.NewInt(8))
	case ethtypes.AccessListTxType:
		// ACL txs are defined to use 0 and 1 as their recovery id, add
		// 27 to become equivalent to unprotected Homestead signatures.
		V = new(big.Int).Add(V, big.NewInt(27))
	default:
		return nil, fmt.Errorf("unknown tx type")
	}
	if e.GetChainID().Cmp(signer.chainId) != 0 {
		return nil, fmt.Errorf("invalid chain id")
	}
	return recoverPlain(e.GetSignHash(), R, S, V, true)
}

func (e *EthTransaction) GetTo() *types.Address {
	if e.inner.GetTo() == nil {
		return nil
	}
	return types.NewAddress(e.inner.GetTo().Bytes())
}

func (e *EthTransaction) GetPayload() []byte {
	return e.inner.GetData()
}

func (e *EthTransaction) GetNonce() uint64 {
	return e.inner.GetNonce()
}

func (e *EthTransaction) GetAmount() uint64 {
	return e.inner.GetValue().Uint64()
}

func (e *EthTransaction) GetTimeStamp() int64 {
	return e.time.UnixNano()
}

func (e *EthTransaction) GetHash() *types.Hash {
	if hash := e.hash.Load(); hash != nil {
		return hash.(*types.Hash)
	}

	var h *types.Hash
	if e.GetType() == ethtypes.LegacyTxType {
		hash := ethtypes.RlpHash(e.inner)
		h = types.NewHash(hash.Bytes())
	} else {
		hash := ethtypes.PrefixedRlpHash(e.GetType(), e.inner)
		h = types.NewHash(hash.Bytes())
	}
	e.hash.Store(h)
	return h
}

func (e *EthTransaction) GetIBTP() *IBTP {
	return nil
}

func (e *EthTransaction) GetExtra() []byte {
	return nil
}

func (e *EthTransaction) GetGas() uint64 {
	return e.inner.GetGas()
}

func (e *EthTransaction) GetGasPrice() *big.Int {
	return e.inner.GetGasPrice()
}

func (e *EthTransaction) GetChainID() *big.Int {
	return e.inner.GetChainID()
}

func (e *EthTransaction) MarshalWithFlag() ([]byte, error) {
	data, err := e.MarshalBinary()
	if err != nil {
		return nil, err
	}

	txData := append([]byte{1}, data...)

	return txData, nil
}

func (e *EthTransaction) Size() int {
	if size := e.size.Load(); size != nil {
		return size.(int)
	}
	c := writeCounter(0)
	ethtypes.RlpEncode(&c, &e.inner)
	e.size.Store(int(c))
	return int(c)
}

func (e *EthTransaction) MarshalTo(buf []byte) (int, error) {
	data, err := e.MarshalBinary()
	if err != nil {
		return 0, err
	}

	copy(buf, data)

	return len(data), nil
}

func (e *EthTransaction) Unmarshal(buf []byte) error {
	return e.UnmarshalBinary(buf)
}

// Type returns the transaction type.
func (e *EthTransaction) GetType() byte {
	return e.inner.TxType()
}

func (e *EthTransaction) SizeWithFlag() int {
	return e.Size() + 1
}

func (e *EthTransaction) GetSignature() []byte {
	var sig []byte
	v, r, s := e.inner.RawSignatureValues()
	sig = append(sig, r.Bytes()...)
	sig = append(sig, s.Bytes()...)
	sig = append(sig, v.Bytes()...)

	return sig
}

func (e *EthTransaction) GetSignHash() *types.Hash {
	switch e.GetType() {
	case ethtypes.LegacyTxType:
		hash := ethtypes.RlpHash([]interface{}{
			e.GetNonce(),
			e.GetGasPrice(),
			e.GetGas(),
			e.inner.GetTo(),
			e.inner.GetValue(),
			e.GetPayload(),
			signer.chainId, uint(0), uint(0),
		})

		return types.NewHash(hash.Bytes())
	case ethtypes.AccessListTxType:
		hash := ethtypes.PrefixedRlpHash(
			e.GetType(),
			[]interface{}{
				signer.chainId,
				e.GetNonce(),
				e.GetGasPrice(),
				e.GetGas(),
				e.inner.GetTo(),
				e.inner.GetValue(),
				e.GetPayload(),
				e.inner.GetAccessList(),
			})

		return types.NewHash(hash.Bytes())
	default:
		// This _should_ not happen, but in case someone sends in a bad
		// json struct via RPC, it's probably more prudent to return an
		// empty hash instead of killing the node with a panic
		//panic("Unsupported transaction type: %d", tx.typ)
		return nil
	}
}

func (e *EthTransaction) IsIBTP() bool {
	return false
}

// RawSignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (e *EthTransaction) GetRawSignature() (v, r, s *big.Int) {
	return e.inner.RawSignatureValues()
}

func (e *EthTransaction) VerifySignature() error {
	if e.GetFrom() == nil {
		return fmt.Errorf("verify signature failed")
	}

	return nil
}

//// AccessList returns the access list of the transaction.
//func (e *EthTransaction) AccessList() types2.AccessList {
//	return e.inner.GetAccessList()
//}

// EncodeRLP implements rlp.Encoder
func (tx *EthTransaction) EncodeRLP(w io.Writer) error {
	if tx.GetType() == ethtypes.LegacyTxType {
		return ethtypes.RlpEncode(w, tx.inner)
	}
	// It's an EIP-2718 typed TX envelope.
	buf := encodeBufferPool.Get().(*bytes.Buffer)
	defer encodeBufferPool.Put(buf)
	buf.Reset()
	if err := tx.encodeTyped(buf); err != nil {
		return err
	}
	return ethtypes.RlpEncode(w, buf.Bytes())
}

// encodeTyped writes the canonical encoding of a typed transaction to w.
func (tx *EthTransaction) encodeTyped(w *bytes.Buffer) error {
	w.WriteByte(tx.GetType())
	return ethtypes.RlpEncode(w, tx.inner)
}

// MarshalBinary returns the canonical encoding of the transaction.
// For legacy transactions, it returns the RLP encoding. For EIP-2718 typed
// transactions, it returns the type and payload.
func (tx *EthTransaction) MarshalBinary() ([]byte, error) {
	if tx.GetType() == ethtypes.LegacyTxType {
		return ethtypes.RlpEncodeToBytes(tx.inner)
	}
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}

// DecodeRLP implements rlp.Decoder
//func (tx *EthTransaction) DecodeRLP(s *rlp.Stream) error {
//	kind, size, err := s.Kind()
//	switch {
//	case err != nil:
//		return err
//	case kind == rlp.List:
//		// It's a legacy transaction.
//		var inner ethtypes.LegacyTx
//		err := s.Decode(&inner)
//		if err == nil {
//			tx.setDecoded(&inner, int(rlp.ListSize(size)))
//		}
//		return err
//	case kind == rlp.String:
//		// It's an EIP-2718 typed TX envelope.
//		var b []byte
//		if b, err = s.Bytes(); err != nil {
//			return err
//		}
//		inner, err := tx.decodeTyped(b)
//		if err == nil {
//			tx.setDecoded(inner, len(b))
//		}
//		return err
//	default:
//		return rlp.ErrExpectedList
//	}
//}

// UnmarshalBinary decodes the canonical encoding of transactions.
// It supports legacy RLP transactions and EIP2718 typed transactions.
func (tx *EthTransaction) UnmarshalBinary(b []byte) error {
	if len(b) > 0 && b[0] > 0x7f {
		// It's a legacy transaction.
		var data ethtypes.LegacyTx
		err := ethtypes.RlpDecodeBytes(b, &data)
		if err != nil {
			return err
		}
		tx.setDecoded(&data, len(b))
		return nil
	}
	// It's an EIP2718 typed transaction envelope.
	inner, err := tx.decodeTyped(b)
	if err != nil {
		return err
	}
	tx.setDecoded(inner, len(b))
	return nil
}

// decodeTyped decodes a typed transaction from the canonical format.
func (tx *EthTransaction) decodeTyped(b []byte) (ethtypes.TxData, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("empty tx type")
	}
	switch b[0] {
	case ethtypes.AccessListTxType:
		var inner ethtypes.AccessListTx
		err := ethtypes.RlpDecodeBytes(b[1:], &inner)
		return &inner, err
	default:
		return nil, fmt.Errorf("unsupported tx type")
	}
}

// setDecoded sets the inner transaction and size after decoding.
func (tx *EthTransaction) setDecoded(inner ethtypes.TxData, size int) {
	tx.inner = inner
	tx.time = time.Now()
	if size > 0 {
		tx.size.Store(size)
	}
}

func (e *EthTransaction) FromCallArgs(callArgs ethtypes.CallArgs) {
	if callArgs.From != nil {
		e.from.Store(types.NewAddress(callArgs.From.Bytes()))
	}

	inner := &ethtypes.AccessListTx{
		GasPrice: (*big.Int)(callArgs.GasPrice),
		To:       callArgs.To,
		Value:    (*big.Int)(callArgs.Value),
	}

	if callArgs.Gas != nil {
		inner.Gas = (uint64)(*callArgs.Gas)
	}

	if callArgs.GasPrice == nil {
		inner.GasPrice = big.NewInt(0)
	}

	if callArgs.Value == nil {
		inner.Value = big.NewInt(0)
	}

	if callArgs.Data != nil {
		inner.Data = ([]byte)(*callArgs.Data)
	}

	if callArgs.AccessList != nil {
		inner.AccessList = *callArgs.AccessList
	}

	e.inner = inner
}
