package pb

import (
	"crypto/sha256"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Transaction) Hash() types.Hash {
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

	return *types.Bytes2Hash(data[:])
}

func (m *Transaction) SignHash() types.Hash {
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

	return *types.Bytes2Hash(ret[:])
}

func (m *Transaction) Sign(key crypto.PrivateKey) error {
	sign, err := key.Sign(m.SignHash().Bytes())
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
	if m.IsIBTP() {
		return fmt.Sprintf("%s-%s-%d", m.IBTP.From, m.IBTP.To, m.IBTP.Category())
	}
	return m.From.Hex()
}
