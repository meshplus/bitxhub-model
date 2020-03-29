package pb

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/meshplus/bitxhub-kit/types"
)

func (r *Receipt) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		r.Status,
		r.Ret,
		r.Events,
		r.TxHash,
		r.Version,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}

func (r *Receipt) IsSuccess() bool {
	return r.Status == Receipt_SUCCESS
}
