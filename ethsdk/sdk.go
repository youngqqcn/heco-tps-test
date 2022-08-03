package ethsdk

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/youngqqcn/heco-tps-test/contractx"
)

type EthSDK struct {
	Url    string
	Client *ethclient.Client
}

var defaultGasLimit = 21000

func NewEthSDK(url string) (*EthSDK, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	sdk := &EthSDK{
		Url:    url,
		Client: client,
	}
	return sdk, nil
}

func (e *EthSDK) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := e.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}

func (e *EthSDK) GetNet() (uint64, uint64, error) {
	cid, err := e.Client.ChainID(context.Background())
	if err != nil {
		return 0, 0, err
	}
	nid, err := e.Client.NetworkID(context.Background())
	if err != nil {
		return 0, 0, err
	}
	return cid.Uint64(), nid.Uint64(), nil
}

func (e *EthSDK) GetBlockNum() (uint64, error) {
	return e.Client.BlockNumber(context.Background())
}

func (e *EthSDK) GetPeerCount() (uint64, error) {
	return e.Client.PeerCount(context.Background())
}

func (e *EthSDK) GetBlockHeaderByNumber(num uint64) (*types.Header, error) {
	return e.Client.HeaderByNumber(context.Background(), big.NewInt(int64(num)))
}

func (e *EthSDK) GetBlockByNumber(num uint64) (*types.Block, error) {
	return e.Client.BlockByNumber(context.Background(), big.NewInt(int64(num)))
}

func (e *EthSDK) GetBlockByTxHash(txHash string) (*types.Block, error) {
	tx, _, err := e.Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	_ = tx
	re, err := e.Client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	return e.Client.BlockByNumber(context.Background(), re.BlockNumber)
}

func (e *EthSDK) GetBlockByHash(hash string) (*types.Block, error) {
	return e.Client.BlockByHash(context.Background(), common.HexToHash(hash))
}

func (e *EthSDK) GetCollectibleInstance(contractAddress string) (*contract.Contract, error) {
	address := common.HexToAddress(contractAddress)
	return contract.NewContract(address, e.Client)
}

func (e *EthSDK) ContractCaller(privateKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	publicKeyECDSA := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKeyECDSA)
	nonce, err := e.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
