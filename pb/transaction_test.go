package pb

import (
	"testing"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/stretchr/testify/require"
)

func TestTransaction_SignHash(t *testing.T) {
	tx := &Transaction{
		From:      types.String2Address("002233"),
		To:        types.String2Address("223344"),
		Timestamp: 1567345493,
		Data:      nil,
		Nonce:     0,
	}
	ret := tx.SignHash()
	require.Equal(t, "0x3ec3562d6b9330f06ebd46a0f4bfa0db562d172e979ed6e73cab98eb6b1da5ad", ret.Hex())
}

func TestTransaction_MarshalTo(t *testing.T) {
	tx := &Transaction{
		From: types.String2Address("002233"),
		To:   types.String2Address("223344"),
	}

	data, err := tx.Marshal()
	require.Nil(t, err)

	txx := &Transaction{}
	err = txx.Unmarshal(data)
	require.Nil(t, err)
	require.Equal(t, tx.SignHash(), txx.SignHash())
}
