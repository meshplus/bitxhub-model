package pb

import (
	"sort"

	"github.com/meshplus/bitxhub-kit/types"
)

type InterchainMeta struct {
	Counter        map[string]*VerifiedIndexSlice
	L2Roots        []types.Hash
	TimeoutCounter map[string]*StringSlice
	TimeoutL2Roots []types.Hash
}

type Interchain struct {
	ID                      string
	InterchainCounter       map[string]uint64
	ReceiptCounter          map[string]uint64
	SourceInterchainCounter map[string]uint64
	SourceReceiptCounter    map[string]uint64
}

func (m *InterchainMeta) Marshal() ([]byte, error) {
	ims := &InterchainMetaS{
		Counter:        stringVerifiedIndexSliceMapToSlice(m.Counter),
		L2Roots:        m.L2Roots,
		TimeoutCounter: stringStringSliceMapToSlice(m.TimeoutCounter),
		TimeoutL2Roots: m.TimeoutL2Roots,
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

	if ims.TimeoutCounter != nil {
		m.TimeoutCounter = ims.TimeoutCounter.toMap()
	}
	m.TimeoutL2Roots = ims.TimeoutL2Roots

	return nil
}

func (m *Interchain) Marshal() ([]byte, error) {
	ics := &InterchainS{
		ID:                      m.ID,
		InterchainCounter:       stringUint64MapToSlice(m.InterchainCounter),
		ReceiptCounter:          stringUint64MapToSlice(m.ReceiptCounter),
		SourceInterchainCounter: stringUint64MapToSlice(m.SourceInterchainCounter),
		SourceReceiptCounter:    stringUint64MapToSlice(m.SourceReceiptCounter),
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

	if ics.SourceInterchainCounter != nil {
		m.SourceInterchainCounter = ics.SourceInterchainCounter.toMap()
	} else {
		m.SourceInterchainCounter = make(map[string]uint64)
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

func (sum *StringVerifiedIndexMap) toMap() map[string]*VerifiedIndex {
	m := make(map[string]*VerifiedIndex)

	for i := range sum.Keys {
		m[sum.Keys[i]] = sum.Vals[i]
	}

	return m
}

func (sum *StringVerifiedIndexSliceMap) toMap() map[string]*VerifiedIndexSlice {
	m := make(map[string]*VerifiedIndexSlice)

	for i := range sum.Keys {
		m[sum.Keys[i]] = sum.Vals[i]
	}

	return m
}

func (sum *StringStringSliceMap) toMap() map[string]*StringSlice {
	m := make(map[string]*StringSlice)

	for i := range sum.Keys {
		m[sum.Keys[i]] = sum.Vals[i]
	}

	return m
}

func stringVerifiedIndexSliceMapToSlice(m map[string]*VerifiedIndexSlice) *StringVerifiedIndexSliceMap {
	sum := &StringVerifiedIndexSliceMap{}

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

func stringStringSliceMapToSlice(m map[string]*StringSlice) *StringStringSliceMap {
	sum := &StringStringSliceMap{}

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
