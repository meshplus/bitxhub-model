package pb

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	mt "github.com/cbergoon/merkletree"
	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/types"
)

type TransactionHash []byte

// CalculateHash hashes the values of a TestContent
func (t TransactionHash) CalculateHash() ([]byte, error) {
	return t, nil
}

// Equals tests for equality of two Contents
func (t TransactionHash) Equals(other mt.Content) (bool, error) {
	tOther, ok := other.(TransactionHash)
	if !ok {
		return false, errors.New("parameter should be type TransactionHash")
	}
	return bytes.Equal(t, tOther), nil
}

func (m *Transaction) Hash() types.Hash {
	tx := &Transaction{
		From:      m.From,
		To:        m.To,
		Timestamp: m.Timestamp,
		Data:      m.Data,
		Nonce:     m.Nonce,
		Signature: m.Signature,
	}

	body, err := tx.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (m *Transaction) SignHash() types.Hash {
	tx := &Transaction{
		From:      m.From,
		To:        m.To,
		Timestamp: m.Timestamp,
		Data:      m.Data,
		Nonce:     m.Nonce,
	}

	body, err := tx.Marshal()
	if err != nil {
		panic(err)
	}

	ret := sha256.Sum256(body)

	return types.Bytes2Hash(ret[:])
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

func (m *Transaction) GetIBTP() (*IBTP, error) {
	if m.Data == nil {
		return nil, fmt.Errorf("empty transaction data")
	}

	if m.Data.Payload == nil {
		return nil, fmt.Errorf("empty transaction payload")
	}

	pl := &InvokePayload{}
	if err := pl.Unmarshal(m.Data.Payload); err != nil {
		return nil, fmt.Errorf("unmarshal payload: %w", err)
	}

	if len(pl.Args) == 0 {
		return nil, fmt.Errorf("ibtp in tx is empty")
	}

	ibtp := &IBTP{}
	if err := ibtp.Unmarshal(pl.Args[0].Value); err != nil {
		return nil, fmt.Errorf("unmarshal ibtp from tx :%w", err)
	}

	return ibtp, nil
}
