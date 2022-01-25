package pb

import (
	"crypto/sha256"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"strconv"
	"strings"

	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxid"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) ServicePair() string {
	return fmt.Sprintf("%s-%s", m.From, m.To)
}

// ParseFrom should be called after CheckServiceID is called
func (m *IBTP) ParseFrom() (string, string, string) {
	if isDID, _ := m.CheckFormat(); isDID {
		methodID, serviceID, _ := ParseFullDID(m.From)
		return "did:bitxhub", methodID, serviceID
	}
	bxhID, chainID, serviceID, _ := ParseFullServiceID(m.From)
	return bxhID, chainID, serviceID
}

// ParseTo should be called after CheckServiceID is called
func (m *IBTP) ParseTo() (string, string, string) {
	if isDID, _ := m.CheckFormat(); isDID {
		methodID, serviceID, _ := ParseFullDID(m.To)
		return "did:bitxhub", methodID, serviceID
	}
	bxhID, chainID, serviceID, _ := ParseFullServiceID(m.To)
	return bxhID, chainID, serviceID
}

func (m *IBTP) CheckServiceID() error {
	isDID, err := m.CheckFormat()
	if err != nil {
		return err
	}
	if isDID {
		return m.CheckDID()
	}
	_, _, _, err = ParseFullServiceID(m.From)
	if err != nil {
		return err
	}
	_, _, _, err = ParseFullServiceID(m.To)
	return err
}

func (m *IBTP) CheckFormat() (bool, error) {
	fromSplits := strings.Split(m.From, ":")
	toSplits := strings.Split(m.To, ":")
	if len(fromSplits) != len(toSplits) {
		return false, fmt.Errorf("(%s) is not the same format to (%s)", m.From, m.To)
	}
	if len(fromSplits) == 4 {
		return true, nil
	}
	if len(fromSplits) == 3 {
		return false, nil
	}
	return false, fmt.Errorf("invalid format (from: %s, to: %s)", m.From, m.To)
}

func (m *IBTP) CheckDID() error {
	_, _, err := ParseFullDID(m.From)
	if err != nil {
		return err
	}
	_, _, err = ParseFullDID(m.To)
	return err
}

func (m *IBTP) SetExtra() error {
	pd := Payload{}
	if err := pd.Unmarshal(m.Payload); err != nil {
		return fmt.Errorf("unmarshal payload error: %w", err)
	}
	ct := Content{}
	if err := ct.Unmarshal(pd.Content); err != nil {
		return fmt.Errorf("unmarshal content error: %w", err)
	}
	if ct.GetFunc() == "interchainInstructionIssue" || ct.GetFunc() == "interchainRuleIssue" {
		m.Extra = ct.GetArgs()[1]
	}
	if ct.GetFunc() == "interchainUpload" {
		index := string(ct.GetArgs()[0])
		splits := strings.Split(index, ":")
		id := splits[0] + ":" + splits[1] + ":" + splits[2] + ":" + "result-" + strings.Split(splits[3], "-")[1]
		m.Extra = []byte(id)
	}

	return nil
}

func (m *IBTP) ParseDIDFrom() (string, string) {
	methodID, serviceID, _ := ParseFullDID(m.From)
	return methodID, serviceID
}

func (m *IBTP) ParseDIDTo() (string, string) {
	methodID, serviceID, _ := ParseFullDID(m.To)
	return methodID, serviceID
}

func ParseFullDID(id string) (string, string, error) {
	splits := strings.Split(id, ":")
	if len(splits) != 4 || splits[0] != "did" || splits[1] != "bitxhub" {
		return "", "", fmt.Errorf("invalid DID format: %s", id)
	}
	return splits[2], splits[3], nil
}

func GetAddrFromDoc(doc *bitxid.MethodDoc) (string, error) {
	type publicKeyInfo struct {
		Raw       asn1.RawContent
		Algorithm pkix.AlgorithmIdentifier
		PublicKey asn1.BitString
	}

	key := doc.PublicKey[0].PublicKeyPem
	keyType, err := crypto.CryptoNameToType(doc.PublicKey[0].Type)
	if err != nil {
		return "", fmt.Errorf("get key type failed: %w", err)
	}

	b, _ := pem.Decode([]byte(key))
	if b == nil {
		return "", fmt.Errorf("decode public key pem failed")
	}

	var pki publicKeyInfo
	if _, err := asn1.Unmarshal(b.Bytes, &pki); err != nil {
		return "", fmt.Errorf("asn1 unmarshal public key error: %w", err)
	}
	pub, err := ecdsa.UnmarshalPublicKey(pki.PublicKey.RightAlign(), keyType)
	if err != nil {
		return "", fmt.Errorf("unmarshal public key error: %w", err)
	}
	addr, err := pub.Address()
	if err != nil {
		return "", fmt.Errorf("get address from public key failed: %w", err)
	}

	return addr.String(), nil
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
		return "", "", 0, fmt.Errorf("invalid  IBTP ID: %s", id)
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
