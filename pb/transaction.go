package pb

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"

	mt "github.com/cbergoon/merkletree"
	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/merkle/merkletree"
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

// Verify that MerkleWrapper is valid: Related transactions in MerkleWrapper can be verified by merkle existence
func (m *MerkleWrapper) VerifyMerkleWrapper() (bool, error) {
	HashesForMerkle := make([]interface{}, 0, len(m.TransactionHashes))
	for _, txHash := range m.TransactionHashes {
		HashesForMerkle = append(HashesForMerkle, TransactionHash(txHash.Bytes()))
	}

	txTree := merkletree.NewMerkleTree()
	err := txTree.InitMerkleTree(HashesForMerkle)

	if err != nil {
		return false, err
	}
	for _, tx := range m.Transactions {
		exist, err := txTree.VerifyContent(TransactionHash(tx.TransactionHash.Bytes()))
		if err != nil || !exist {
			return false, nil
		}
	}
	return true, nil
}

func (m *Transaction) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		m.From,
		m.To,
		m.Timestamp,
		m.Nonce,
		m.Signature,
		m.Data,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (m *Transaction) SignHash() types.Hash {
	var data []byte
	if m.Data != nil {
		b, err := m.Data.Marshal()
		if err != nil {
			panic(err)
		}
		data = b
	}

	c := fmt.Sprintf("from=%s&to=%s&timestamp=%d&nonce=%d&data=%x",
		m.From.Hex(),
		m.To.Hex(),
		m.Timestamp,
		m.Nonce,
		data,
	)

	ret := sha256.Sum256([]byte(c))

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
