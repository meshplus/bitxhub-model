package pb

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/meshplus/bitxhub-kit/types"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) ServicePair() string {
	return fmt.Sprintf("%s-%s", m.From, m.To)
}

// ParseFrom should be called after CheckServiceID is called
func (m *IBTP) ParseFrom() (string, string, string) {
	bxhID, chainID, serviceID, _ := ParseFullServiceID(m.From)
	return bxhID, chainID, serviceID
}

// ParseTo should be called after CheckServiceID is called
func (m *IBTP) ParseTo() (string, string, string) {
	bxhID, chainID, serviceID, _ := ParseFullServiceID(m.To)
	return bxhID, chainID, serviceID
}

func (m *IBTP) CheckServiceID() error {
	_, _, _, err := ParseFullServiceID(m.From)
	if err != nil {
		return err
	}
	_, _, _, err = ParseFullServiceID(m.To)
	return err
}

func ParseFullServiceID(id string) (string, string, string, error) {
	splits := strings.Split(id, ":")
	if len(splits) != 3 {
		return "", "", "", fmt.Errorf("invalid full service ID: %s", id)
	}
	return splits[0], splits[1], splits[2], nil
}

func ParseServicePair(servicePair string) (string, string, error) {
	splits := strings.Split(servicePair, "-")
	if len(splits) != 2 {
		return "", "", fmt.Errorf("invalid service pair: %s", servicePair)
	}
	return splits[0], splits[1], nil
}

func ParseIBTPID(id string) (string, string, uint64, error) {
	splits := strings.Split(id, "-")
	if len(splits) != 3 {
		return "", "", 0, fmt.Errorf("invalid IBTP ID: %s", id)
	}

	index, err := strconv.Atoi(splits[2])
	if err != nil {
		return "", "", 0, fmt.Errorf("invalid IBTP ID: %s", id)
	}

	return splits[0], splits[1], uint64(index), nil
}

func GenServicePair(from, to string) string {
	return fmt.Sprintf("%s-%s", from, to)
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
	case IBTP_INTERCHAIN:
		return IBTP_REQUEST
	case IBTP_RECEIPT_SUCCESS, IBTP_RECEIPT_FAILURE, IBTP_RECEIPT_ROLLBACK:
		return IBTP_RESPONSE
	}
	return IBTP_UNKNOWN
}
