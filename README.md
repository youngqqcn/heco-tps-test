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