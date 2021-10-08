# 理解状态与交易

以太坊是一个巨大的分布式状态机，交易则是驱动这个状态机的力量。

在运行预置代码之前，请首先在1#终端启动节点仿真器：

```
~$ ganache-cli -d
```

## 查看账户余额

在2#终端进入`~/repo/chapter4/check-balance`目录并运行程序：

```
~$ cd ~/repo/chapter4/check-balance
~/repo/chapter4/check-balance$ go run main.go
```

## 进行单位换算

在2#终端进入`~/repo/chapter4/unit-conversion`目录并运行程序：

```
~$ cd ~/repo/chapter4/unit-conversion
~/repo/chapter4/unit-conversion$ go run main.go
```

## 执行普通交易

在2#终端进入`~/repo/chapter4/transaction-only`目录并运行程序：

```
~$ cd ~/repo/chapter4/transaction-only
~/repo/chapter4/transaction-only$ go run main.go
```

## 执行普通交易并读取回执

在2#终端进入`~/repo/chapter4/transaction-receipt`目录并运行程序：

```
~$ cd ~/repo/chapter4/transaction-receipt
~/repo/chapter4/transaction-receipt$ go run main.go
```

## 执行裸交易

在2#终端进入`~/repo/chapter4/raw-transaction`目录并运行程序：

```
~$ cd ~/repo/chapter4/raw-transaction
~/repo/chapter4/raw-transaction$ go run main.go
```