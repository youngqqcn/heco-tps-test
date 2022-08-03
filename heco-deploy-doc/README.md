搭建2个节点的heco测试链，使用congress共识




# 编译heco代码

> 参考：https://learnblockchain.cn/article/3844


使用`v1.2.2`版本

```shell
git checkout v1.2.2

# 编译geth
make 

#创建geth软连接
ln -s xxxx/build/bin/geth /usr/bin/geth

# 执行以下命令，安装bootnode
env GOBIN= go install ./cmd/bootnode


ln -s $GOBIN/bootnode /usr/bin/bootnode

```


# 部署

使用 `myheco` 中的`Makefile`

```shell
make clean
make init

# 在终端1执行
make bootnode

# 在终端2执行
make start-node1

# 在终端3执行
make start-node2

```

