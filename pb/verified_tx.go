package pb

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/cbergoon/merkletree"
)

var _ merkletree.Content = (*VerifiedTx)(nil)

func (m *VerifiedTx) CalculateHash() ([]byte, error) {
	validByte := byte(0)
	if m.Valid {
		validByte = 1
	}

	data := append(m.Tx.Hash().Bytes(), validByte)
	hash := sha256.Sum256(data)

	return hash[:], nil
}

func (m *VerifiedTx) Equals(other merkletree.Content) (bool, error) {
	tOther, ok := other.(*VerifiedTx)
	if !ok {
		return false, errors.New("parameter should be type VerifiedTx")
	}

	if !bytes.Equal(m.Tx.Hash().RawHash[:], tOther.Tx.Hash().RawHash[:]) {
		return false, errors.New("hash fields in 2 verifiedTx are not equal")
	}

	if m.Valid != tOther.Valid {
		return false, errors.New("valid fields in 2 verifiedTx are not equal")
	}

	return true, nil
}
