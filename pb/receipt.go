package pb

import (
	"crypto/sha256"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Receipt) Hash() *types.Hash {
	receipt := &Receipt{
		Status:  m.Status,
		Ret:     m.Ret,
		Events:  m.Events,
		TxHash:  m.TxHash,
		Version: m.Version,
	}
	body, err := receipt.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (m *Receipt) IsSuccess() bool {
	return m.Status == Receipt_SUCCESS
}
