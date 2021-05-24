package pb

import (
	"encoding/binary"
	"fmt"
	"math/big"

	proto "github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-kit/types"
)

const BxhTxType = 2

type TxConstructer func() Transaction

var supportedTx map[uint64]TxConstructer

func RegisterTxConstructor(flag uint64, txConstructer TxConstructer) {
	if supportedTx == nil {
		supportedTx = make(map[uint64]TxConstructer)
	}

	supportedTx[flag] = txConstructer
}

type Transaction interface {
	GetVersion() []byte
	GetFrom() *types.Address
	GetTo() *types.Address
	GetPayload() []byte
	GetNonce() uint64
	GetAmount() uint64
	GetTimeStamp() int64
	GetHash() *types.Hash
	GetIBTP() *IBTP
	GetExtra() []byte
	GetGas() uint64
	GetGasPrice() *big.Int
	GetChainID() *big.Int
	GetSignature() []byte
	GetRawSignature() (*big.Int, *big.Int, *big.Int)
	GetSignHash() *types.Hash
	GetType() byte
	IsIBTP() bool
	MarshalWithFlag() ([]byte, error)
	Size() int
	SizeWithFlag() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	VerifySignature() error
}

type Transactions struct {
	Transactions []Transaction
}

func (txs *Transactions) Reset()         { *txs = Transactions{} }
func (txs *Transactions) String() string { return proto.CompactTextString(txs) }
func (txs *Transactions) ProtoMessage()  {}

func (txs *Transactions) MarshalTo(data []byte) (int, error) {
	txsData, err := txs.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, txsData)

	return len(txsData), nil
}

func (txs *Transactions) Marshal() ([]byte, error) {
	var txsData []byte

	for _, tx := range txs.Transactions {
		size := tx.SizeWithFlag()
		sizeByte := make([]byte, 8)
		binary.LittleEndian.PutUint64(sizeByte, uint64(size))
		txsData = append(txsData, sizeByte...)

		txData, err := tx.MarshalWithFlag()
		if err != nil {
			return nil, err
		}
		txsData = append(txsData, txData...)
	}

	return txsData, nil
}

func (txs *Transactions) Size() int {
	size := 0

	for _, tx := range txs.Transactions {
		size += tx.SizeWithFlag() + 8
	}

	return size
}

func (txs *Transactions) Unmarshal(data []byte) error {
	l := len(data)
	index := 0

	for index < l {
		if l-index < 9 {
			return fmt.Errorf("unmarshal txs failed, invalid data size: l %d, index %d", l, index)
		}

		size := int(binary.LittleEndian.Uint64(data[index : index+8]))
		txData := data[index+8 : index+8+size]
		tx, err := UnmarshalTx(txData)
		if err != nil {
			return err
		}

		txs.Transactions = append(txs.Transactions, tx)

		index += size + 8
	}

	return nil
}

func UnmarshalTx(data []byte) (Transaction, error) {
	if len(data) == 0 {
		return nil, nil
	}

	constrcutor, ok := supportedTx[uint64(data[0])]
	if !ok {
		return nil, fmt.Errorf("unexpected tx type: %d", data[0])
	}

	tx := constrcutor()
	if err := tx.Unmarshal(data[1:]); err != nil {
		return nil, err
	}

	return tx, nil
}
