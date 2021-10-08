# Hi，以太坊

使用不同的方式接入以太坊。

在运行预置代码之前，请首先在1#终端启动节点仿真器并保持其运行：

```
~$ ganache-cli
```

## 使用curl命令行

在2#终端进入`~/repo/chapter2`目录并运行bash脚本：

```
~$ cd ~/repo/chapter2
~/repo/chapter2$ ./rpc-curl.sh
```

## 使用http代码

在2#终端进入`~/repo/chapter2/rpc-http`目录并运行Go应用：

```
~$ cd ~/repo/chapter2/rpc-http
~/repo/chapter2/rpc-http$ go run main.go
```

## 使用http + codec代码

在2#终端进入`~/repo/chapter2/rpc-http-codec`目录并运行Go应用：

```
~$ cd ~/repo/chapter2/rpc-http-codec
~/repo/chapter2/rpc-http-codec$ go run main.go
```

## 使用geth rpc客户端代码

在2#终端进入`~/repo/chapter2/rpc-geth`目录并运行Go应用：

```
~$ cd ~/repo/chapter2/rpc-geth
~/repo/chapter2/rpc-geth$ go run main.go
```

## 使用geth ethclient客户端代码

在2#终端进入`~/repo/chapter2/rpc-geth-ethclient`目录并运行Go应用：

```
~$ cd ~/repo/chapter2/rpc-geth-ethclient
~/repo/chapter2/rpc-geth-ethclient$ go run main.go
```

## 使用ethtool客户端代码

在2#终端进入`~/repo/chapter2/rpc-ethtool`目录并运行Go应用：

```
~$ cd ~/repo/chapter2/rpc-ethtool
~/repo/chapter2/rpc-ethtool$ go run main.go
```
