package pb

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) ServicePair() string {
	return fmt.Sprintf("%s-%s", m.From, m.To)
}

// SrcChainID should be called after CheckServiceID is called
func (m *IBTP) SrcChainID() string {
	_, chainID, _, _ := parseFullServiceID(m.From)
	return chainID
}

// DstChainID should be called after CheckServiceID is called
func (m *IBTP) DstChainID() string {
	_, chainID, _, _ := parseFullServiceID(m.To)
	return chainID
}

func (m *IBTP) CheckServiceID() error {
	_, _, _, err := parseFullServiceID(m.From)
	if err != nil {
		return err
	}
	_, _, _, err = parseFullServiceID(m.To)
	return err
}

func parseFullServiceID(id string) (string, string, string, error) {
	splits := strings.Split(id, ":")
	if len(splits) != 3 {
		return "", "", "", fmt.Errorf("invalid full service ID: %s", id)
	}
	return splits[0], splits[1], splits[2], nil
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
