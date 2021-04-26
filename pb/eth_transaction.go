package pb

import (
	"bytes"
	"fmt"
	io "io"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/meshplus/bitxhub-kit/types"
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

type sigCache struct {
	signer EIP155Signer
	from   common.Address
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
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

// TxData is the underlying data of a transaction.
//
// This is implemented by LegacyTx and AccessListTx.
type TxData interface {
	txType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	chainID() *big.Int
	accessList() types2.AccessList
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	value() *big.Int
	nonce() uint64
	to() *common.Address

	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(chainID, v, r, s *big.Int)
}

// AccessListTx is the data of EIP-2930 access list transactions.
type AccessListTx struct {
	ChainID    *big.Int          // destination chain ID
	Nonce      uint64            // nonce of sender account
	GasPrice   *big.Int          // wei per gas
	Gas        uint64            // gas limit
	To         *common.Address   `rlp:"nil"` // nil means contract creation
	Value      *big.Int          // wei amount
	Data       []byte            // contract invocation input data
	AccessList types2.AccessList // EIP-2930 access list
	V, R, S    *big.Int          // signature values
}

// LegacyTx is the transaction data of regular Ethereum transactions.
type LegacyTx struct {
	Nonce    uint64          // nonce of sender account
	GasPrice *big.Int        // wei per gas
	Gas      uint64          // gas limit
	To       *common.Address `rlp:"nil"` // nil means contract creation
	Value    *big.Int        // wei amount
	Data     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *AccessListTx) copy() TxData {
	cpy := &AccessListTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		AccessList: make(types2.AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.

func (tx *AccessListTx) txType() byte                  { return types2.AccessListTxType }
func (tx *AccessListTx) chainID() *big.Int             { return tx.ChainID }
func (tx *AccessListTx) protected() bool               { return true }
func (tx *AccessListTx) accessList() types2.AccessList { return tx.AccessList }
func (tx *AccessListTx) data() []byte                  { return tx.Data }
func (tx *AccessListTx) gas() uint64                   { return tx.Gas }
func (tx *AccessListTx) gasPrice() *big.Int            { return tx.GasPrice }
func (tx *AccessListTx) value() *big.Int               { return tx.Value }
func (tx *AccessListTx) nonce() uint64                 { return tx.Nonce }
func (tx *AccessListTx) to() *common.Address           { return tx.To }

func (tx *AccessListTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *AccessListTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *LegacyTx) copy() TxData {
	cpy := &LegacyTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.

func (tx *LegacyTx) txType() byte                  { return types2.LegacyTxType }
func (tx *LegacyTx) chainID() *big.Int             { return deriveChainId(tx.V) }
func (tx *LegacyTx) accessList() types2.AccessList { return nil }
func (tx *LegacyTx) data() []byte                  { return tx.Data }
func (tx *LegacyTx) gas() uint64                   { return tx.Gas }
func (tx *LegacyTx) gasPrice() *big.Int            { return tx.GasPrice }
func (tx *LegacyTx) value() *big.Int               { return tx.Value }
func (tx *LegacyTx) nonce() uint64                 { return tx.Nonce }
func (tx *LegacyTx) to() *common.Address           { return tx.To }

func (tx *LegacyTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *LegacyTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}

type writeCounter int

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

// deriveChainId derives the chain id from the given v parameter
func deriveChainId(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

func (e *EthTransaction) GetVersion() []byte {
	return nil
}

// Protected says whether the transaction is replay-protected.
func (e *EthTransaction) Protected() bool {
	switch tx := e.inner.(type) {
	case *LegacyTx:
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
	if Vb.BitLen() > 8 {
		return nil, fmt.Errorf("invalid signature")
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return nil, fmt.Errorf("invalid signature")
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash.Bytes(), sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, fmt.Errorf("invalid public key")
	}
	return types.NewAddress(crypto.Keccak256(pub[1:])[12:]), nil
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
	case types2.LegacyTxType:
		if !e.Protected() {
			hash := rlpHash([]interface{}{
				e.GetNonce(),
				e.GetGasPrice(),
				e.GetGas(),
				e.inner.to(),
				e.inner.value(),
				e.GetPayload(),
			})
			addr, err := recoverPlain(hash, R, S, V, true)
			if err != nil {
				return nil, fmt.Errorf("invalid signature")
			}
			return addr, nil
		}
		V = new(big.Int).Sub(V, signer.chainIdMul)
		V.Sub(V, big.NewInt(8))
	case types2.AccessListTxType:
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
	if e.inner.to() == nil {
		return nil
	}
	return types.NewAddress(e.inner.to().Bytes())
}

func (e *EthTransaction) GetPayload() []byte {
	return e.inner.data()
}

func (e *EthTransaction) GetNonce() uint64 {
	return e.inner.nonce()
}

func (e *EthTransaction) GetAmount() uint64 {
	return e.inner.value().Uint64()
}

func (e *EthTransaction) GetTimeStamp() int64 {
	return e.time.UnixNano()
}

func (e *EthTransaction) GetHash() *types.Hash {
	if hash := e.hash.Load(); hash != nil {
		return hash.(*types.Hash)
	}

	var h *types.Hash
	if e.GetType() == types2.LegacyTxType {
		h = rlpHash(e.inner)
	} else {
		h = prefixedRlpHash(e.GetType(), e.inner)
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
	return e.inner.gas()
}

func (e *EthTransaction) GetGasPrice() *big.Int {
	return e.inner.gasPrice()
}

func (e *EthTransaction) GetChainID() *big.Int {
	return e.inner.chainID()
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
	rlp.Encode(&c, &e.inner)
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
	return e.inner.txType()
}

func (e *EthTransaction) SizeWithFlag() int {
	return e.Size() + 1
}

func (e *EthTransaction) GetSignature() []byte {
	var sig []byte
	v, r, s := e.inner.rawSignatureValues()
	sig = append(sig, r.Bytes()...)
	sig = append(sig, s.Bytes()...)
	sig = append(sig, v.Bytes()...)

	return sig
}

func (e *EthTransaction) GetSignHash() *types.Hash {
	switch e.GetType() {
	case types2.LegacyTxType:
		return rlpHash([]interface{}{
			e.GetNonce(),
			e.GetGasPrice(),
			e.GetGas(),
			e.inner.to(),
			e.inner.value(),
			e.GetPayload(),
			signer.chainId, uint(0), uint(0),
		})
	case types2.AccessListTxType:
		return prefixedRlpHash(
			e.GetType(),
			[]interface{}{
				signer.chainId,
				e.GetNonce(),
				e.GetGasPrice(),
				e.GetGas(),
				e.inner.to(),
				e.inner.value(),
				e.GetPayload(),
				e.AccessList(),
			})
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
	return e.inner.rawSignatureValues()
}

func (e *EthTransaction) VerifySignature() error {
	if e.GetFrom() == nil {
		return fmt.Errorf("verify signature failed")
	}

	return nil
}

// AccessList returns the access list of the transaction.
func (e *EthTransaction) AccessList() types2.AccessList {
	return e.inner.accessList()
}

// EncodeRLP implements rlp.Encoder
func (tx *EthTransaction) EncodeRLP(w io.Writer) error {
	if tx.GetType() == types2.LegacyTxType {
		return rlp.Encode(w, tx.inner)
	}
	// It's an EIP-2718 typed TX envelope.
	buf := encodeBufferPool.Get().(*bytes.Buffer)
	defer encodeBufferPool.Put(buf)
	buf.Reset()
	if err := tx.encodeTyped(buf); err != nil {
		return err
	}
	return rlp.Encode(w, buf.Bytes())
}

// encodeTyped writes the canonical encoding of a typed transaction to w.
func (tx *EthTransaction) encodeTyped(w *bytes.Buffer) error {
	w.WriteByte(tx.GetType())
	return rlp.Encode(w, tx.inner)
}

// MarshalBinary returns the canonical encoding of the transaction.
// For legacy transactions, it returns the RLP encoding. For EIP-2718 typed
// transactions, it returns the type and payload.
func (tx *EthTransaction) MarshalBinary() ([]byte, error) {
	if tx.GetType() == types2.LegacyTxType {
		return rlp.EncodeToBytes(tx.inner)
	}
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}

// DecodeRLP implements rlp.Decoder
func (tx *EthTransaction) DecodeRLP(s *rlp.Stream) error {
	kind, size, err := s.Kind()
	switch {
	case err != nil:
		return err
	case kind == rlp.List:
		// It's a legacy transaction.
		var inner LegacyTx
		err := s.Decode(&inner)
		if err == nil {
			tx.setDecoded(&inner, int(rlp.ListSize(size)))
		}
		return err
	case kind == rlp.String:
		// It's an EIP-2718 typed TX envelope.
		var b []byte
		if b, err = s.Bytes(); err != nil {
			return err
		}
		inner, err := tx.decodeTyped(b)
		if err == nil {
			tx.setDecoded(inner, len(b))
		}
		return err
	default:
		return rlp.ErrExpectedList
	}
}

// UnmarshalBinary decodes the canonical encoding of transactions.
// It supports legacy RLP transactions and EIP2718 typed transactions.
func (tx *EthTransaction) UnmarshalBinary(b []byte) error {
	if len(b) > 0 && b[0] > 0x7f {
		// It's a legacy transaction.
		var data LegacyTx
		err := rlp.DecodeBytes(b, &data)
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
func (tx *EthTransaction) decodeTyped(b []byte) (TxData, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("empty tx type")
	}
	switch b[0] {
	case types2.AccessListTxType:
		var inner AccessListTx
		err := rlp.DecodeBytes(b[1:], &inner)
		return &inner, err
	default:
		return nil, fmt.Errorf("unsupported tx type")
	}
}

// setDecoded sets the inner transaction and size after decoding.
func (tx *EthTransaction) setDecoded(inner TxData, size int) {
	tx.inner = inner
	tx.time = time.Now()
	if size > 0 {
		tx.size.Store(size)
	}
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) *types.Hash {
	var h types.Hash
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h.RawHash[:])
	return &h
}

// prefixedRlpHash writes the prefix into the hasher before rlp-encoding x.
// It's used for typed transactions.
func prefixedRlpHash(prefix byte, x interface{}) *types.Hash {
	var h types.Hash
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	sha.Write([]byte{prefix})
	rlp.Encode(sha, x)
	sha.Read(h.RawHash[:])
	return &h
}
