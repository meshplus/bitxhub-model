package pb

import (
	"encoding/json"

	"github.com/meshplus/eth-kit/types"
)

// MarshalJSON marshals as JSON.
func (l EvmLog) MarshalJSON() ([]byte, error) {
	var enc types.Log

	if l.Address != nil {
		enc.Address = types.BytesToAddress(l.Address.Bytes())
	}
	for _, topic := range l.Topics {
		enc.Topics = append(enc.Topics, types.BytesToHash(topic.Bytes()))
	}
	enc.Data = l.Data
	enc.BlockNumber = types.Uint64(l.BlockNumber)
	if l.TxHash != nil {
		enc.TxHash = types.BytesToHash(l.TxHash.Bytes())
	}
	enc.TxIndex = types.Uint(l.TxIndex)
	if l.BlockHash != nil {
		enc.BlockHash = types.BytesToHash(l.BlockHash.Bytes())
	}
	enc.Index = types.Uint(l.Index)
	enc.Removed = l.Removed

	return json.Marshal(&enc)
}
