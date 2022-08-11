package consvr

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	contract "github.com/youngqqcn/heco-tps-test/contractx"
	"github.com/youngqqcn/heco-tps-test/ethsdk"
)

type ContractSvc struct {
	ethSDK           *ethsdk.EthSDK
	contractAddress  string
	contractInstance *contract.Contract
	mu               *sync.Mutex
	counter          int64
}

func NewContractSvc(ethsdk *ethsdk.EthSDK, contractAddress string) *ContractSvc {
	return &ContractSvc{
		ethSDK:          ethsdk,
		contractAddress: contractAddress,
		mu:              new(sync.Mutex),
		counter:         0,
	}
}

func (c *ContractSvc) Load() error {
	ins, err := c.ethSDK.GetCollectibleInstance(c.contractAddress)
	if err != nil {
		return err
	}
	c.contractInstance = ins

	return nil
}

// 仅一次性测试，很多变量直接写死了。
// 如果要多次测试，需要先将节点数据清空重启，然后跑测试
func (c *ContractSvc) Mint(batchCount int) (string, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	// to := "0xf513e4e5Ded9B510780D016c482fC158209DE9AA"
	// privateKey, err := crypto.HexToECDSA("5ea30eea9ba9500f3601f7659f0ccace819c562456e2f745fb2555918ab32277")
	to := "0x8284B6412ef6eFA75adDEa85f07E7de5f8F8ec48"
	privateKey, err := crypto.HexToECDSA("cfe945f87d61aa82e903804bcc32bacdf130ae47268a2f6d7a3d877cbf028ff6")
	if err != nil {
		return "", err
	}
	chainId := big.NewInt(2285)
	caller, err := c.ethSDK.ContractCaller(privateKey, chainId)

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce_pending, err := c.ethSDK.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	data := ""
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	txs := make([]string, 0)
	nonce := int64(nonce_pending)

	for i := 0; i < batchCount; i++ {
		caller.Nonce = big.NewInt(nonce + int64(i))
		tokenId := c.counter

		toAddr := common.HexToAddress(to)
		tx, err := c.contractInstance.Mint(caller, toAddr, big.NewInt(tokenId), dataBytes)
		if err != nil {
			return "", err
		}
		txs = append(txs, tx.Hash().Hex())
		c.counter += 1
	}
	fmt.Printf("len(txs) is %d\n", len(txs))

	return "===finished==", nil
}

// func (c *ContractSvc) PurchaseCollectible(owner *model.Account, purchaser *model.Account, tokenId int64, price *big.Int) error {
// 	// transfer eth
// 	purchaserPrivate, err := crypto.HexToECDSA(purchaser.PrivateKey[2:])
// 	if err != nil {
// 		return err
// 	}
// 	ttx, err := c.ethSDK.Transfer(purchaserPrivate, owner.Address, price.Int64())
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.sdb.TxLogRecord(&model.TxLog{
// 		Tx:      ttx,
// 		Type:    model.TxLogType_Purchase,
// 		From:    purchaser.Address,
// 		To:      owner.Address,
// 		Created: time.Now().Unix(),
// 	}); err != nil {
// 		return err
// 	}
// 	log.Logger.Infof("Transfer tx: %s\n", ttx)

// 	// transfer token with platform(approver)
// 	ownerPK, err := util.GetPrivateKeyByHash(owner.PrivateKey)
// 	if err != nil {
// 		return err
// 	}
// 	_ = ownerPK

// 	auth, err := c.ethSDK.ContractCaller(c.platform)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := c.contractInstance.SafeTransferFrom(
// 		auth,
// 		common.HexToAddress(owner.Address),
// 		common.HexToAddress(purchaser.Address),
// 		big.NewInt(tokenId),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if err := c.sdb.TxLogRecord(&model.TxLog{
// 		Tx:      tx.Hash().Hex(),
// 		From:    owner.Address,
// 		To:      purchaser.Address,
// 		Type:    model.TxLogType_Transfer,
// 		Created: time.Now().Unix(),
// 	}); err != nil {
// 		return err
// 	}
// 	log.Logger.Infof("Purchase tx: %s\n", tx.Hash().Hex())

// 	// update database
// 	coll, err := c.sdb.CollectibleGet(tokenId)
// 	if err != nil {
// 		return err
// 	}
// 	coll.Onsale = false
// 	coll.Price = ""
// 	if err := c.sdb.CollectibleUpdate(coll); err != nil {
// 		return err
// 	}
// 	cowner := &model.CollectibleOwner{
// 		TokenId: tokenId,
// 		Owner:   purchaser.Address,
// 	}
// 	err = c.sdb.CollectibleOwnerUpdate(cowner)
// 	return err
// }

// // read collecible token counter
// func (c *ContractSvc) cCounter() (int64, error) {
// 	return c.sdb.CollectibleCounterGet()
// }
