package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/youngqqcn/heco-tps-test/consvr"
	"github.com/youngqqcn/heco-tps-test/ethsdk"
)

type EthSDK struct {
	Url    string
	client *ethclient.Client
}

var defaultGasLimit = 21000

var CHAIN_ID = big.NewInt(2285)

// var CHAIN_ID = big.NewInt(1337)
// var CHAIN_ID = big.NewInt(128)

func NewEthSDK(url string) (*EthSDK, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	sdk := &EthSDK{
		Url:    url,
		client: client,
	}
	return sdk, nil
}

func (e *EthSDK) Transfer(from *ecdsa.PrivateKey, to string, amount int64) (string, error) {
	fromAddress := crypto.PubkeyToAddress(from.PublicKey)
	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	value := big.NewInt(amount)
	gasLimit := uint64(defaultGasLimit)
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	toAddress := common.HexToAddress(to)

	tx := types.NewTransaction(
		nonce,
		toAddress,
		value,
		gasLimit,
		gasPrice,
		nil,
	)

	// chainId, err := e.client.NetworkID(context.Background())
	// if err != nil {
	// 	return "", err
	// }

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(CHAIN_ID), from)
	if err != nil {
		return "", err
	}

	if err := e.client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (e *EthSDK) BatchTransfer(from *ecdsa.PrivateKey, batchCount int, to string, amount int64) ([]string, error) {

	txs := make([]string, 0)

	fromAddress := crypto.PubkeyToAddress(from.PublicKey)
	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return txs, err
	}
	value := big.NewInt(amount)
	gasLimit := uint64(defaultGasLimit)
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return txs, err
	}

	toAddress := common.HexToAddress(to)

	// chainId, err := e.client.NetworkID(context.Background())
	// if err != nil {
	// 	return txs, err
	// }

	for i := 0; i < batchCount; i++ {

		tx := types.NewTransaction(
			nonce+uint64(i),
			toAddress,
			value,
			gasLimit,
			gasPrice,
			nil,
		)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(CHAIN_ID), from)
		if err != nil {
			return []string{}, err
		}

		if err := e.client.SendTransaction(context.Background(), signedTx); err != nil {
			return []string{}, err
		}
		// fmt.Printf("%v\n", signedTx)
		txs = append(txs, signedTx.Hash().Hex())
	}
	return txs, nil
}

func (e *EthSDK) BatchTransfer2(from *ecdsa.PrivateKey, batchCount int, tos []string, amount int64) ([]string, error) {

	txs := make([]string, 0)

	fromAddress := crypto.PubkeyToAddress(from.PublicKey)
	nonce, err := e.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return txs, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(defaultGasLimit)
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return txs, err
	}
	gasPrice = big.NewInt(0)

	// chainId, err := e.client.NetworkID(context.Background())
	// if err != nil {
	// 	return txs, err
	// }

	for i := 0; i < batchCount; i++ {

		toAddress := common.HexToAddress(tos[i])

		tx := types.NewTransaction(
			nonce+uint64(i),
			toAddress,
			value,
			gasLimit,
			gasPrice,
			nil,
		)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(CHAIN_ID), from)
		if err != nil {
			return []string{}, err
		}
		// stx, _ := signedTx.MarshalBinary()
		// fmt.Printf("signedtx=====>%v\n", hexutil.Encode(stx))
		// return []string{}, err

		if err := e.client.SendTransaction(context.Background(), signedTx); err != nil {
			return []string{}, err
		}
		// fmt.Printf("%v\n", signedTx)
		txs = append(txs, signedTx.Hash().Hex())
	}
	fmt.Printf("last tx : %s\n", txs[len(txs)-1])
	return txs, nil
}

func getAdddress(count int) map[string]*ecdsa.PrivateKey {
	ret := make(map[string]*ecdsa.PrivateKey, 0)

	for i := 0; i < count; i++ {
		k, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(k.PublicKey).Hex()
		ret[addr] = k
	}

	return ret
}

func toWei(ether int64) int64 {
	ret := big.NewInt(1)
	ret.Mul(big.NewInt(ether), math.BigPow(10, 18))
	return ret.Int64()
}

func ethTps() {

	// init eth sdk
	// sdk, err := NewEthSDK("http://192.168.110.95:8545")
	sdk, err := NewEthSDK("http://localhost:8545")
	// sdk, err := NewEthSDK("https://http-mainnet.hecochain.com/")
	if err != nil {
		fmt.Printf("Init sdk failed: %s", err)
		return
	}

	// sender is 0xf513e4e5Ded9B510780D016c482fC158209DE9AA
	// privkey, err := crypto.HexToECDSA("5ea30eea9ba9500f3601f7659f0ccace819c562456e2f745fb2555918ab32277")
	// sender is 0x8284B6412ef6eFA75adDEa85f07E7de5f8F8ec48
	privkey, err := crypto.HexToECDSA("cfe945f87d61aa82e903804bcc32bacdf130ae47268a2f6d7a3d877cbf028ff6")
	if err != nil {
		fmt.Printf("priv kery error: %v\n", err)
		return
	}

	subAddrs := getAdddress(1)
	tos := make([]string, 0)
	for k := range subAddrs {
		tos = append(tos, k)
	}

	txs, err := sdk.BatchTransfer2(privkey, len(tos), tos, toWei(10000))
	if err != nil {
		fmt.Printf("BatchTransfer2 error: %v\n", err)
		return
	}
	fmt.Printf("BatchTransfer, ret: %d\n", len(txs))

	return

	time.Sleep(10 * time.Second)
	// txhash, err := sdk.Transfer(privkey, "0x8284B6412ef6eFA75adDEa85f07E7de5f8F8ec48", 1000)
	// if err != nil {
	// 	fmt.Printf("transfer error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("tx is %s\n", txhash)

	////////////////////////////////////////////////

	wg := &sync.WaitGroup{}
	wg.Add(len(subAddrs))

	for i := 0; i < len(subAddrs); i++ {
		go func(pk *ecdsa.PrivateKey) {
			defer wg.Done()
			txs, err := sdk.BatchTransfer(pk, 5000, "0x8284B6412ef6eFA75adDEa85f07E7de5f8F8ec48", 1000)
			if err != nil {
				fmt.Printf("transfer error: %v\n", err)
				return
			}
			fmt.Printf("txs len is %d\n", len(txs))

		}(subAddrs[tos[i]])
	}

	wg.Wait()
	fmt.Printf("all goroutine finished\n")
}

func nftTps() {

	// init eth sdk
	// sdk, err := ethsdk.NewEthSDK("http://192.168.110.95:8545")
	// sdk, err := NewEthSDK("http://localhost:8545")
	sdk, err := ethsdk.NewEthSDK("http://localhost:8545")
	if err != nil {
		fmt.Printf("Init sdk failed: %s", err)
		return
	}

	// privkey, err := crypto.HexToECDSA("5ea30eea9ba9500f3601f7659f0ccace819c562456e2f745fb2555918ab32277")
	// if err != nil {
	// 	fmt.Printf("priv kery error: %v\n", err)
	// 	return
	// }

	contractAddress := "0x8dc3b1010dcc7a1b9dbdbb4423445627e9de5919"
	csvr := consvr.NewContractSvc(sdk, contractAddress)
	csvr.Load()

	h, err := csvr.Mint(100)
	if err != nil {
		fmt.Printf("mint error : %s\n", err)
		return
	}
	fmt.Printf("hash: %s\n", h)
}

func main() {

	ethTps()
	// nftTps()
}
