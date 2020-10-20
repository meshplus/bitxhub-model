package pb

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Receipt) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		m.Status,
		m.Ret,
		m.Events,
		m.TxHash,
		m.Version,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return *types.Bytes2Hash(data[:])
}

func (m *Receipt) IsSuccess() bool {
	return m.Status == Receipt_SUCCESS
}
