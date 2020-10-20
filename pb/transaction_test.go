package pb

import (
	"fmt"
	"testing"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/stretchr/testify/require"
)

const (
	hash          = "0x9f41dd84524bf8a42f8ab58ecfca6e1752d6fd93fe8dc00af4c71963c97db59f"
	formalAccount = "0x929545f44692178EDb7FA468B44c5351596184Ba"
)

func TestTransaction_SignHash(t *testing.T) {
	tx := &Transaction{
		From:      *types.String2Address(formalAccount),
		To:        *types.String2Address(formalAccount),
		Timestamp: 1567345493,
		Nonce:     0,
	}
	ret := tx.SignHash()
	fmt.Printf("%s", ret.Hex())
	require.Equal(t, "0x1E5E61c0e80ccF4A543A0EA0Ac6A0cdBc5f55E596eCd6eB8558ed6B94fE7bF86", ret.Hex())
}

func TestTransaction_MarshalTo(t *testing.T) {
	tx := &Transaction{
		From: *types.String2Address(formalAccount),
		To:   *types.String2Address(formalAccount),
	}

	data, err := tx.Marshal()
	require.Nil(t, err)

	txx := &Transaction{}
	err = txx.Unmarshal(data)
	require.Nil(t, err)
	require.Equal(t, tx.SignHash(), txx.SignHash())
}
