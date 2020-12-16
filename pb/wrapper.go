package pb

import (
	"github.com/meshplus/bitxhub-kit/types"
	"sort"
)

type InterchainMeta struct {
	Counter map[string]*Uint64Slice
	L2Roots []types.Hash
}

type Interchain struct {
	ID                   string
	InterchainCounter    map[string]uint64
	ReceiptCounter       map[string]uint64
	SourceReceiptCounter map[string]uint64
}

func (m *InterchainMeta) Marshal() ([]byte, error) {
	ims := &InterchainMetaS{
		Counter: stringUint64SliceMapToSlice(m.Counter),
		L2Roots: m.L2Roots,
	}

	return ims.Marshal()
}

func (m *InterchainMeta) Unmarshal(data []byte) error {
	ims := &InterchainMetaS{}
	if err := ims.Unmarshal(data); err != nil {
		return err
	}

	if ims.Counter != nil {
		m.Counter = ims.Counter.toMap()
	}
	m.L2Roots = ims.L2Roots

	return nil
}

func (m *Interchain) Marshal() ([]byte, error) {
	ics := &InterchainS{
		ID:                   m.ID,
		InterchainCounter:    stringUint64MapToSlice(m.InterchainCounter),
		ReceiptCounter:       stringUint64MapToSlice(m.ReceiptCounter),
		SourceReceiptCounter: stringUint64MapToSlice(m.SourceReceiptCounter),
	}

	return ics.Marshal()
}

func (m *Interchain) Unmarshal(data []byte) error {
	ics := &InterchainS{}
	if err := ics.Unmarshal(data); err != nil {
		return err
	}

	m.ID = ics.ID

	if ics.InterchainCounter != nil {
		m.InterchainCounter = ics.InterchainCounter.toMap()
	} else {
		m.InterchainCounter = make(map[string]uint64)
	}

	if ics.ReceiptCounter != nil {
		m.ReceiptCounter = ics.ReceiptCounter.toMap()
	} else {
		m.ReceiptCounter = make(map[string]uint64)
	}

	if ics.SourceReceiptCounter != nil {
		m.SourceReceiptCounter = ics.SourceReceiptCounter.toMap()
	} else {
		m.SourceReceiptCounter = make(map[string]uint64)
	}

	return nil
}

func (sum *StringUint64Map) toMap() map[string]uint64 {
	m := make(map[string]uint64)

	for i := range sum.Keys {
		m[sum.Keys[i]] = sum.Vals[i]
	}

	return m
}

func (sum *StringUint64SliceMap) toMap() map[string]*Uint64Slice {
	m := make(map[string]*Uint64Slice)

	for i := range sum.Keys {
		m[sum.Keys[i]] = sum.Vals[i]
	}

	return m
}

func stringUint64MapToSlice(m map[string]uint64) *StringUint64Map {
	sum := &StringUint64Map{}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		sum.Keys = append(sum.Keys, k)
		sum.Vals = append(sum.Vals, m[k])
	}

	return sum
}

func stringUint64SliceMapToSlice(m map[string]*Uint64Slice) *StringUint64SliceMap {
	sum := &StringUint64SliceMap{}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		sum.Keys = append(sum.Keys, k)
		sum.Vals = append(sum.Vals, m[k])
	}

	return sum
}
