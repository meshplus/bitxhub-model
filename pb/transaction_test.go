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
		From:      types.NewAddressByStr(formalAccount),
		To:        types.NewAddressByStr(formalAccount),
		Timestamp: 1567345493,
		Nonce:     0,
	}
	ret := tx.SignHash()
	fmt.Printf("%s", ret.String())
	require.Equal(t, "0x9d51B5a879eb058A88419AE932189273B8AA10B6447806d3e5Cc59f85E09b08E", ret.String())
}

func TestTransaction_MarshalTo(t *testing.T) {
	tx := &Transaction{
		From: types.NewAddressByStr(formalAccount),
		To:   types.NewAddressByStr(formalAccount),
	}

	data, err := tx.Marshal()
	require.Nil(t, err)

	txx := &Transaction{}
	err = txx.Unmarshal(data)
	require.Nil(t, err)
	require.Equal(t, tx.SignHash(), txx.SignHash())
}
