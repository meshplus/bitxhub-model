package pb

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) Account() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Category())
}

func (m *IBTP) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		m.From,
		m.To,
		m.Index,
		m.Type,
		m.Timestamp,
		m.Payload,
		m.Extra,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (m *IBTP) Category() IBTP_Category {
	switch m.Type {
	case IBTP_INTERCHAIN, IBTP_ASSET_EXCHANGE_INIT, IBTP_ASSET_EXCHANGE_REDEEM, IBTP_ASSET_EXCHANGE_REFUND:
		return IBTP_REQUEST
	case IBTP_RECEIPT_SUCCESS, IBTP_RECEIPT_FAILURE, IBTP_ASSET_EXCHANGE_RECEIPT:
		return IBTP_RESPONSE
	}
	return IBTP_UNKNOWN
}
