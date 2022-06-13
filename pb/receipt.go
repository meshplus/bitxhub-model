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
	if l.TransactionHash != nil {
		enc.TxHash = *l.TransactionHash
	}
	enc.TxIndex = hexutil.Uint(l.TransactionIndex)
	if l.BlockHash != nil {
		enc.BlockHash = *l.BlockHash
	}
	enc.Index = hexutil.Uint(l.LogIndex)
	enc.Removed = l.Removed

	return json.Marshal(&enc)
}

func (e *Event) IsAuditEvent() bool {
	switch e.EventType {
	case Event_AUDIT_PROPOSAL, Event_AUDIT_APPCHAIN, Event_AUDIT_RULE, Event_AUDIT_SERVICE, Event_AUDIT_NODE, Event_AUDIT_ROLE, Event_AUDIT_INTERCHAIN, Event_AUDIT_DAPP:
		return true
	default:
		return false
	}
}
