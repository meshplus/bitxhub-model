package pb

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/meshplus/bitxhub-kit/hexutil"
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

	return types.NewHash(data[:])
}

func (m *Receipt) IsSuccess() bool {
	return m.Status == Receipt_SUCCESS
}

type Log struct {
	Address     types.Address  `json:"address" gencodec:"required"`
	Topics      []types.Hash   `json:"topics" gencodec:"required"`
	Data        hexutil.Bytes  `json:"data" gencodec:"required"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	TxHash      types.Hash     `json:"transactionHash" gencodec:"required"`
	TxIndex     hexutil.Uint   `json:"transactionIndex"`
	BlockHash   types.Hash     `json:"blockHash"`
	Index       hexutil.Uint   `json:"logIndex"`
	Removed     bool           `json:"removed"`
}

//MarshalJSON marshals as JSON.
func (l EvmLog) MarshalJSON() ([]byte, error) {
	var enc Log

	if l.Address != nil {
		enc.Address = *l.Address
	}
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	if l.TxHash != nil {
		enc.TxHash = *l.TxHash
	}
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	if l.BlockHash != nil {
		enc.BlockHash = *l.BlockHash
	}
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed

	return json.Marshal(&enc)
}
