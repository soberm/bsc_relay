package relay

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type relayer struct {
	ethClient        *ethclient.Client
	bscClient        *ethclient.Client
	bscRelayContract *BSCRelayContract
	privateKey       *ecdsa.PrivateKey
	batchSize        uint
	ethChainID       *big.Int
	bscChainID       *big.Int
}

func NewRelayer(c Config) (*relayer, error) {
	ethClient, err := ethclient.Dial(c.EthereumHost)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %w", err)
	}

	bscClient, err := ethclient.Dial(c.BSCHost)
	if err != nil {
		return nil, fmt.Errorf("dial bsc client: %w", err)
	}

	bscRelayContract, err := NewBSCRelayContract(common.HexToAddress(c.ContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("new relay contract: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("private key hex to ecdsa: %w", err)
	}

	ethChainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get source chain id: %w", err)
	}

	bscChainID, err := bscClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get source chain id: %w", err)
	}

	return &relayer{
		ethClient:        ethClient,
		bscClient:        bscClient,
		bscRelayContract: bscRelayContract,
		privateKey:       privateKey,
		batchSize:        c.BatchSize,
		ethChainID:       ethChainID,
		bscChainID:       bscChainID,
	}, nil
}

func (r *relayer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := r.relay(ctx); err != nil {
				log.Errorf("relay: %v", err)
			}
		}
	}
}

func (r *relayer) relay(ctx context.Context) error {
	currentHeight, err := r.bscRelayContract.GetCurrentHeight(nil)
	if err != nil {
		return fmt.Errorf("get current height: %w", err)
	}

	blockNumber, err := r.bscClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("blocknumber: %w", err)
	}

	unsignedHeaders := make([][]byte, 0)
	signedHeaders := make([][]byte, 0)

	diff := blockNumber - currentHeight.Uint64()
	if diff > uint64(r.batchSize) {
		diff = uint64(r.batchSize)
	}
	targetHeight := currentHeight.Uint64() + diff

	for i := currentHeight.Uint64(); i < targetHeight; i++ {
		header, err := r.bscClient.HeaderByNumber(ctx, big.NewInt(int64(i+1)))
		if err != nil {
			return fmt.Errorf("header by number: %w", err)
		}

		signedHeader := new(bytes.Buffer)
		err = rlp.Encode(signedHeader, header)
		if err != nil {
			return fmt.Errorf("rlp encode signed header: %w", err)
		}
		signedHeaders = append(signedHeaders, signedHeader.Bytes())

		unsignedHeader, err := EncodeUnsignedHeaderToRLP(header, r.bscChainID)
		if err != nil {
			return fmt.Errorf("rlp encode unsigned header: %w", err)
		}
		unsignedHeaders = append(unsignedHeaders, unsignedHeader)
	}

	log.Infof("Submitting headers from height %d to %d", currentHeight.Int64(), targetHeight)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(r.privateKey, r.ethChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}

	tx, err := r.bscRelayContract.SubmitBlockHeaderBatch(transactOpts, unsignedHeaders, signedHeaders)
	if err != nil {
		return fmt.Errorf("submit block header: %w", err)
	}

	_, err = bind.WaitMined(ctx, r.ethClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}

	return nil
}
