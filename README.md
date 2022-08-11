# 测试Heco的TPS



执行流程

1、进入`heco-deploy-doc`, 根据README.md，执行makefile中的命令，执行命令，启动2个节点的测试链

2、进入`hardhat`，根据README.md，执行makefile中的命令，部署智能合约

3、在`heco-tps-test`目录，执行`make` 和 `make run`命令进行测试

4、观察节点日志输出，每个区块包含的交易数，就可以计算出TPS



目前只做了简单ERC721的合约的mint交易测试，TPS在150左右。

如果是生产环境的ERC721合约，逻辑一般比较复杂，单笔交易的gas更高，一个区块能够包含的交易数更少。


---

注意： golang测试代码中，nonce 和 tokenid都是写死了，如果要多次测试，需要先清空节点数据，从头来。



---

主账户：0xf513e4e5Ded9B510780D016c482fC158209DE9AA



```
 curl  -H "Content-Type: application/json"  -X POST --data '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xf860038082520894153c0a8e316e43c3069537b8d1bdd20e88bf2d9c80808211fe9f2f01943f1b55e0569aac61ebf462356a6e8eab8884bc17e37bb28ab0478bd5a04632f40f084d7b3e1b9b82343d5f22075d8aea559f0b3a53698c710c681712a8"],"id":1}' 192.168.110.37:8545
{"jsonrpc":"2.0","id":1,"result":"0x4cc327e21f97d5f7b26593643fed803467de0b63b1bd53f32d67724f52a136c0"}

```