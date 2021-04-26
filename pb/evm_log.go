package pb

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MarshalJSON marshals as JSON.
func (l EvmLog) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address     common.Address `json:"address" gencodec:"required"`
		Topics      []common.Hash  `json:"topics" gencodec:"required"`
		Data        hexutil.Bytes  `json:"data" gencodec:"required"`
		BlockNumber hexutil.Uint64 `json:"blockNumber"`
		TxHash      common.Hash    `json:"transactionHash" gencodec:"required"`
		TxIndex     hexutil.Uint   `json:"transactionIndex"`
		BlockHash   common.Hash    `json:"blockHash"`
		Index       hexutil.Uint   `json:"logIndex"`
		Removed     bool           `json:"removed"`
	}
	var enc Log
	if l.Address != nil {
		enc.Address = common.BytesToAddress(l.Address.Bytes())
	}
	for _, topic := range l.Topics {
		enc.Topics = append(enc.Topics, common.BytesToHash(topic.Bytes()))
	}
	enc.Data = l.Data
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	if l.TxHash != nil {
		enc.TxHash = common.BytesToHash(l.TxHash.Bytes())
	}
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	if l.BlockHash != nil {
		enc.BlockHash = common.BytesToHash(l.BlockHash.Bytes())
	}
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}
