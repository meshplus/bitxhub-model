package pb

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/meshplus/bitxhub-kit/types"
)

var _ Transaction = (*BxhTransaction)(nil)

func init() {
	RegisterTxConstructor(0, func() Transaction {
		return &BxhTransaction{}
	})
}

func (m *BxhTransaction) Hash() *types.Hash {
	tx := &BxhTransaction{
		From:      m.From,
		To:        m.To,
		Timestamp: m.Timestamp,
		Payload:   m.Payload,
		IBTP:      m.IBTP,
		Nonce:     m.Nonce,
		Amount:    m.Amount,
		Signature: m.Signature,
	}

	body, err := tx.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.NewHash(data[:])
}

func (m *BxhTransaction) SignHash() *types.Hash {
	tx := &BxhTransaction{
		From:      m.From,
		To:        m.To,
		Timestamp: m.Timestamp,
		Payload:   m.Payload,
		IBTP:      m.IBTP,
		Nonce:     m.Nonce,
		Amount:    m.Amount,
	}

	body, err := tx.Marshal()
	if err != nil {
		panic(err)
	}

	ret := sha256.Sum256(body)

	return types.NewHash(ret[:])
}

func (m *BxhTransaction) Sign(key crypto.PrivateKey) error {
	sign, err := key.Sign(m.SignHash().Bytes())
	if err != nil {
		return err
	}

	m.Signature = sign

	return nil
}

func (m *BxhTransaction) GetCrosschainExtra() (*CrosschainTransactionExtra, error) {
	extra := &CrosschainTransactionExtra{}
	if err := proto.Unmarshal(m.Extra, extra); err != nil {
		return nil, err
	}

	return extra, nil
}

func (m *BxhTransaction) IsIBTP() bool {
	return m.IBTP != nil
}

func (m *BxhTransaction) Account() string {
	return m.From.String()
}

func (m *BxhTransaction) GetFrom() *types.Address {
	return m.From
}

func (m *BxhTransaction) GetTo() *types.Address {
	return m.To
}

func (m *BxhTransaction) GetTimeStamp() int64 {
	return m.Timestamp
}

func (m *BxhTransaction) GetHash() *types.Hash {
	if m.TransactionHash == nil {
		m.TransactionHash = m.Hash()
	}

	return m.TransactionHash
}

func (m *BxhTransaction) GetGas() uint64 {
	return 0
}

func (m *BxhTransaction) GetGasPrice() *big.Int {
	return big.NewInt(0)
}

func (m *BxhTransaction) GetChainID() *big.Int {
	return big.NewInt(0)
}

func (m *BxhTransaction) MarshalWithFlag() ([]byte, error) {
	data, err := m.Marshal()
	if err != nil {
		return nil, err
	}

	txData := append([]byte{0}, data...)

	return txData, nil
}

func (m *BxhTransaction) SizeWithFlag() int {
	return m.Size() + 1
}

func (m *BxhTransaction) GetSignHash() *types.Hash {
	return m.SignHash()
}

// RawSignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (m *BxhTransaction) GetRawSignature() (v, r, s *big.Int) {
	if len(m.Signature) != 65 {
		return nil, nil, nil
	}

	r = &big.Int{}
	r.SetBytes(m.Signature[:32])
	s = &big.Int{}
	s.SetBytes(m.Signature[32:64])
	v = &big.Int{}
	v.SetBytes(m.Signature[64:])

	return v, r, s
}

// RawSignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (m *BxhTransaction) GetType() byte {
	return BxhTxType
}

func (m *BxhTransaction) VerifySignature() error {
	ok, err := asym.Verify(crypto.Secp256k1, m.GetSignature(), m.GetSignHash().Bytes(), *m.GetFrom())
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

func (m *BxhTransaction) GetValue() *big.Int {
	if m.Amount == "" {
		return big.NewInt(0)
	}

	if val, ok := new(big.Int).SetString(m.Amount, 10); ok {
		return val
	}

	return big.NewInt(0)
}
