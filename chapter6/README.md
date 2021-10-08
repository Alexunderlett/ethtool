# 过滤器与发布订阅

使用过滤器和发布订阅机制来检测以太坊区块链上感兴趣的事件。

在运行预置代码之前，请首先在1#终端启动节点仿真器：

```
~$ ganache-cli -d
```

## 拉取模式监听新块的生成

在2#终端进入`~/repo/chapter6/block-monitor-pull`并运行程序：

```
~$ cd ~/repo/chapter6/block-monitor-pull
~/repo/chapter6/block-monitor-pull$ go run main.go
```

## 推送模式监听新块的生成

在2#终端进入`~/repo/chapter6/block-monitor-push`并运行程序：

```
~$ cd ~/repo/chapter6/block-monitor-push
~/repo/chapter6/block-monitor-push$ go run main.go
```

## 拉取模式监听新交易的生成

在2#终端进入`~/repo/chapter6/tx-monitor-pull`并运行程序：

```
~$ cd ~/repo/chapter6/tx-monitor-pull
~/repo/chapter6/tx-monitor-pull$ go run main.go
```

## 推送模式监听新交易的生成

在2#终端进入`~/repo/chapter6/tx-monitor-push`并运行程序：

```
~$ cd ~/repo/chapter6/tx-monitor-push
~/repo/chapter6/tx-monitor-push$ go run main.go
```

## 拉取模式监听待定交易的生成

在2#终端进入`~/repo/chapter6/pending-tx-monitor-pull`并运行程序：

```
~$ cd ~/repo/chapter6/pending-tx-monitor-pull
~/repo/chapter6/pending-tx-monitor-pull$ go run main.go
```

## 推送模式监听待定交易的生成

在2#终端进入`~/repo/chapter6/pending-tx-monitor-push`并运行程序：

```
~$ cd ~/repo/chapter6/pending-tx-monitor-push
~/repo/chapter6/pending-tx-monitor-push$ go run main.go
```

## 拉取模式监听合约日志

如果还没有部署合约，那么首先在2#终端执行如下命令编译并部署代币合约：

```
~/repo/chapter6$ ./build-contract.sh
~/repo/chapter6$ cd deploy-contract
~/repo/chapter6/deploy-contract$ go run main.go
```

在2#终端进入`~/repo/chapter6/log-monitor-pull`并运行程序：

```
~$ cd ~/repo/chapter6/log-monitor-pull
~/repo/chapter6/log-monitor-pull$ go run main.go
```

## 推送模式监听合约日志

如果还没有部署合约，那么首先在2#终端执行如下命令编译并部署代币合约：

```
~/repo/chapter6$ ./build-contract.sh
~/repo/chapter6$ cd deploy-contract
~/repo/chapter6/deploy-contract$ go run main.go
```

在2#终端进入`~/repo/chapter6/log-monitor-push`并运行程序：

```
~$ cd ~/repo/chapter6/log-monitor-push
~/repo/chapter6/log-monitor-push$ go run main.go
```

## 使用Go封装包监听合约日志

如果还没有部署合约，那么首先在2#终端执行如下命令编译并部署代币合约：

```
~/repo/chapter6$ ./build-contract.sh
~/repo/chapter6$ cd deploy-contract
~/repo/chapter6/deploy-contract$ go run main.go
```

在2#终端进入`~/repo/chapter6/log-monitor-bind`并运行程序：

```
~$ cd ~/repo/chapter6/log-monitor-bind
~/repo/chapter6/log-monitor-bind$ go run main.go
```

