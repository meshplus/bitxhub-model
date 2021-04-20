package pb

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym"

	"github.com/meshplus/bitxhub-kit/hexutil"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/stretchr/testify/require"
)

const (
	hash          = "0x9f41dd84524bf8a42f8ab58ecfca6e1752d6fd93fe8dc00af4c71963c97db59f"
	formalAccount = "0x929545f44692178EDb7FA468B44c5351596184Ba"
)

func TestTransaction_SignHash(t *testing.T) {
	tx := &BxhTransaction{
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
	tx := &BxhTransaction{
		From: types.NewAddressByStr(formalAccount),
		To:   types.NewAddressByStr(formalAccount),
	}

	data, err := tx.Marshal()
	require.Nil(t, err)

	txx := &BxhTransaction{}
	err = txx.Unmarshal(data)
	require.Nil(t, err)
	require.Equal(t, tx.SignHash(), txx.SignHash())
}

func TestUnmarshalTx(t *testing.T) {
	tx := &BxhTransaction{
		From:      types.NewAddressByStr(formalAccount),
		To:        types.NewAddressByStr(formalAccount),
		Timestamp: 1567345493,
		Nonce:     0,
	}
	txs := &Transactions{Transactions: []Transaction{tx}}

	txsData, err := txs.Marshal()
	require.Nil(t, err)

	txs2 := &Transactions{}
	err = txs2.Unmarshal(txsData)
	require.Nil(t, err)

	require.Equal(t, 1, len(txs2.Transactions))
	tx2 := txs2.Transactions[0]

	tx2, ok := tx2.(*BxhTransaction)
	require.True(t, ok)
	txData, err := tx.Marshal()
	require.Nil(t, err)

	tx3 := tx2.(*BxhTransaction)
	tx3Data, err := tx3.Marshal()
	require.Nil(t, err)

	require.Equal(t, txData, tx3Data)
}

func TestEthTransaction_GetSignHash(t *testing.T) {
	rawTx := "0xf86c8085147d35700082520894f927bb571eaab8c9a361ab405c9e4891c5024380880de0b6b3a76400008025a00b8e3b66c1e7ae870802e3ef75f1ec741f19501774bd5083920ce181c2140b99a0040c122b7ebfb3d33813927246cbbad1c6bf210474f5d28053990abff0fd4f53"
	tx := &EthTransaction{}
	tx.Unmarshal(hexutil.Decode(rawTx))

	addr := "0xC63573cB77ec56e0A1cb40199bb85838D71e4dce"

	fmt.Println("tx hash:", tx.GetHash().String())

	fmt.Println(tx.GetRawSignature())
	InitEIP155Signer(big.NewInt(1))

	ok, err := asym.Verify(crypto.Secp256k1, tx.GetSignature(), tx.GetSignHash().Bytes(), *types.NewAddressByStr(addr))
	assert.Nil(t, err)
	assert.True(t, ok)
}
