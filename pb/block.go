package pb

import (
	"crypto/sha256"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Block) Hash() *types.Hash {
	blockheader := &BlockHeader{
		Number:      m.BlockHeader.Number,
		ParentHash:  m.BlockHeader.ParentHash,
		StateRoot:   m.BlockHeader.StateRoot,
		TxRoot:      m.BlockHeader.TxRoot,
		ReceiptRoot: m.BlockHeader.ReceiptRoot,
		Version:     m.BlockHeader.Version,
	}
	body, err := blockheader.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (m *Block) Height() uint64 {
	if m == nil || m.BlockHeader == nil {
		return 0
	}

	return m.BlockHeader.Number
}
