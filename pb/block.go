package pb

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Block) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		m.BlockHeader.Number,
		m.BlockHeader.ParentHash,
		m.BlockHeader.StateRoot,
		m.BlockHeader.TxRoot,
		m.BlockHeader.ReceiptRoot,
		m.BlockHeader.Version,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return *types.Bytes2Hash(data[:])
}

func (m *Block) Height() uint64 {
	if m == nil || m.BlockHeader == nil {
		return 0
	}

	return m.BlockHeader.Number
}
