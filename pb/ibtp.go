package pb

import (
	"crypto/sha256"
	"fmt"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) ServicePair() string {
	return fmt.Sprintf("%s-%s", m.From, m.To)
}

func (m *IBTP) Hash() *types.Hash {
	ibtp := &IBTP{
		From:          m.From,
		To:            m.To,
		Index:         m.Index,
		Type:          m.Type,
		TimeoutHeight: m.TimeoutHeight,
		Payload:       m.Payload,
		Extra:         m.Extra,
	}
	body, err := ibtp.Marshal()
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.NewHash(data[:])
}

func (m *IBTP) Category() IBTP_Category {
	switch m.Type {
	case IBTP_INTERCHAIN, IBTP_ROLLBACK:
		return IBTP_REQUEST
	case IBTP_RECEIPT_SUCCESS, IBTP_RECEIPT_FAILURE, IBTP_RECEIPT_ROLLBACK:
		return IBTP_RESPONSE
	}
	return IBTP_UNKNOWN
}
