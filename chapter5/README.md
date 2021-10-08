# 智能合约的开发与交互

学习ERC20代币智能合约的设计并使用solidity开发语言实现，然后使用
Go程序进行部署与交互。

在运行预置代码之前，请首先在1#终端启动节点仿真器：

```
~$ ganache-cli -d
```

## 编译合约并生成Go封装包

在2#终端进入`~/repo/chapter5`目录并执行以下命令：

```
~/repo/chapter5$ ./build-contract.sh
```

## 部署合约原理实现

在2#终端进入`~/repo/chapter5/deploy-contract-theory`目录并运行程序：

```
~$ cd ~/repo/chapter5/deploy-contract-theory
~/repo/chapter5/deploy-contract-theory$ go run main.go
```

## 访问合约原理实现

在2#终端进入`~/repo/chapter5/access-contract-theory`目录并运行程序：

```
~$ cd ~/repo/chapter5/access-contract-theory
~/repo/chapter5/access-contract-theory$ go run main.go
```

## 用Go封装包部署合约

在2#终端进入`~/repo/chapter5/deploy-contract-bind`目录并运行程序：

```
~$ cd ~/repo/chapter5/deploy-contract-bind
~/repo/chapter5/deploy-contract-bind$ go run main.go
```

## 用Go封装包访问合约

在2#终端进入`~/repo/chapter5/access-contract-bind`目录并运行程序：

```
~$ cd ~/repo/chapter5/access-contract-bind
~/repo/chapter5/access-contract-bind$ go run main.go
```
