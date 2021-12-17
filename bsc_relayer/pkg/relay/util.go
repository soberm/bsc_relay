package relay

import (
	"bytes"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

func EncodeUnsignedHeaderToRLP(header *types.Header, chainId *big.Int) ([]byte, error) {
	buffer := new(bytes.Buffer)

	err := rlp.Encode(buffer, []interface{}{
		chainId,
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65],
		header.MixDigest,
		header.Nonce,
	})

	return buffer.Bytes(), err
}
