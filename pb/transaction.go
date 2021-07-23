package pb

import (
	"crypto/sha256"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"
	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Transaction) Hash() *types.Hash {
	tx := &Transaction{
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

func (m *Transaction) SignHash() *types.Hash {
	tx := &Transaction{
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

func (m *Transaction) Sign(key crypto.PrivateKey) error {
	sign, err := asym.SignWithType(key, m.SignHash().Bytes())
	if err != nil {
		return err
	}

	m.Signature = sign

	return nil
}

func (m *Transaction) GetCrosschainExtra() (*CrosschainTransactionExtra, error) {
	extra := &CrosschainTransactionExtra{}
	if err := proto.Unmarshal(m.Extra, extra); err != nil {
		return nil, err
	}

	return extra, nil
}

func (m *Transaction) IsIBTP() bool {
	return m.IBTP != nil
}

func (m *Transaction) Account() string {
	return m.From.String()
}
