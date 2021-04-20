package pb

import (
	"crypto/sha256"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *Block) Hash() *types.Hash {
	return m.BlockHeader.Hash()
}

func (m *Block) Height() uint64 {
	if m == nil || m.BlockHeader == nil {
		return 0
	}

	return m.BlockHeader.Number
}

func (m *BlockHeader) Hash() *types.Hash {
	blockheader := &BlockHeader{
		Number:      m.Number,
		ParentHash:  m.ParentHash,
		StateRoot:   m.StateRoot,
		TxRoot:      m.TxRoot,
		ReceiptRoot: m.ReceiptRoot,
		Version:     m.Version,
	}
	body, err := blockheader.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.NewHash(data[:])
}
