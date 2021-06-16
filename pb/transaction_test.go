package pb

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
	"github.com/meshplus/bitxhub-kit/hexutil"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/stretchr/testify/assert"
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

func TestBxhTransaction_VerifySignature(t *testing.T) {
	addrStr := "0x7a0b4Fa7fe66074C3cf19002D414c754261c54DE"
	msg := "data"
	sig := "0x73b0f712892e6014fe33e1291ae392525c5404a680f71f4f0e1ee6757d0e0e9f3b4d5b00334c5f73b436bfb29216f8e6ba69d204b3cbf94503c2879c37ac9ea41b"

	msgWithPrefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	msgHash := ecdsa.Keccak256([]byte(msgWithPrefix))

	sigBytes := hexutil.Decode(sig)
	r := &big.Int{}
	r.SetBytes(sigBytes[:32])
	s := &big.Int{}
	s.SetBytes(sigBytes[32:64])
	v := &big.Int{}
	v.SetBytes(sigBytes[64:])

	addr, err := ecdsa.RecoverPlain(msgHash, r, s, v, true)
	assert.Nil(t, err)
	assert.Equal(t, addrStr, types.NewAddress(addr).String())
}
