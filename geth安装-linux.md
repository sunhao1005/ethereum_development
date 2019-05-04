# Ubuntu 16.04 快速搭建以太坊私有链开发环境

链接地址：<https://www.jianshu.com/p/ca7bfff540fc>

**安装git**

之后的安装都需要依赖Git

> sudo add-apt-repository ppa:git-core/ppa
>
> sudo apt-get update
>
> sudo apt-get install git

查看版本号

> git --version
>
> git version 2.18.0

1.安装go语言环境

（1）下载linux下go的安装包

> $ wget https://studygolang.com/dl/golang/go1.10.3.linux-amd64.tar.gz

（2）下载完后，进入到下载目录，将其解压到 /usr/local 文件夹下并在$HOME下新建一个文件夹go：

> $ sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
>
> $ mkdir go

   (3)配置环境变量

​    使用如下命令打开环境变量配置文件

> $ sudo vim ~/.bashrc

​    将下列代码复制粘贴到文件最后，wq命令保存退出

> export GOROOT=/usr/local/go
>
> export GOBIN=$GOROOT/bin
>
> export GOPATH=$HOME/go
>
> export PATH=$PATH:$GOBIN

​    输入如下命令使环境变量生效(注意：不用加sudo)

> $ source ~/.bashrc

​    至此，go语言环境已经安装完成，输入go version查看安装是否成功

> $ go version
>
> go version go1.10.3 linux/amd64

[2.安装以太坊客户端Ethereum]()

​    使用如下命令进行安装

> $ sudo apt-get install software-properties-common
>
> $ sudo add-apt-repository -y ppa:ethereum/ethereum
>
> $ sudo apt-get update
>
> $ sudo apt-get install ethereum

​    安装完后，输入geth version查看安装是否成功

> $ geth version
>
> Geth Version: 1.8.11-stable

[3.以太坊私有链搭建]()

（1）新建一个文件夹，例如mychain，用来保存私有链的数据。创建一个初始化的配置文件，genesis.json

> $ mkdir mychain
>
> $ vim genesis.json

然后将如下的配置信息写入到genesis.json文件中（注意：chainId不能为0）

> {
>
> "config": {
>
> ​        "chainId": 8,
>
> ​        "homesteadBlock": 0,
>
> ​        "eip155Block": 0,
>
> ​        "eip158Block": 0
>
> ​    },
>
> "alloc"      : {},
>
> "coinbase"   : "0x0000000000000000000000000000000000000000",
>
> "difficulty" : "0x200",
>
> "extraData"  : "",
>
> "gasLimit"   : "0x2fefd8",
>
> "nonce"      : "0x0000000000000042",
>
> "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
>
> "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
>
> "timestamp"  : "0x00"
>
> }

（2）初始化创世区块，并指定数据存放路径

> $ geth --datadir mychain init genesis.json

（3）启动私有链节点，并将日志输入到mychain.log

> $ geth --networkid 8 --datadir mychain --identity "mychain" --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi "admin,debug,eth,miner,net,personal,shh,txpool,web3" --port 30303 --rpccorsdomain "*" --nodiscover console 2>>mychain.log

具体参数含义请参考这个博客：[geth命令](https://www.cnblogs.com/tinyxiong/p/7918706.html)

启动成功后进入控制台界面，如下所示:

![1556425037402](C:\Users\孙浩\AppData\Roaming\Typora\typora-user-images\1556425037402.png)

[4.测试]()

（1）创建新账户

> personal.newAccount("密码")

​    创建成功将返回账户的地址

> \> personal.newAccount("123")
>
> "0xe6319357003ff9889b7a849f5ec66b2903d08289"
>
> \> user1=eth.accounts[0]
>
> "0x3e46c05151dcce7c994aaeb6f7366fc39ddfc694"
>
> \>user2=eth.accounts[1]
>
> "0x611b1c6a1a859da53ed6154aeeec797ea03b9289"

（2）查询账户余额

> \>eth.getBalance("账户地址")
>
> 0
>
> \>eth.getBalance(user1)  //两种查询方式都可以
>
> 0

（3）开启/停止挖矿，第一次开起需要等待一段时间。

> miner.start()

第一次要等一阵时间，大约5分钟左右

开始都如第一行所示，出现加粗则是挖矿成功了

> INFO [07-26|18:15:21.550] Generating DAG in progress               epoch=1 percentage=16 elapsed=3m25.200s
>
> INFO [07-26|18:15:33.338] Generating DAG in progress               epoch=1 percentage=17 elapsed=3m36.988s
>
> INFO [07-26|18:15:45.307] Generating DAG in progress               epoch=1 percentage=18 elapsed=3m48.956s
>
> INFO [07-26|18:15:57.158] Generating DAG in progress               epoch=1 percentage=19 elapsed=4m0.808s
>
> INFO [07-26|18:15:59.828] Successfully sealed new block            number=30 hash=2b5e96…217cae
>
> INFO [07-26|18:15:59.828]  block reached canonical chain          number=25 hash=37bbe5…bcba21
>
> INFO [07-26|18:15:59.829]  mined potential block                  number=30 hash=2b5e96…217cae
>
> INFO [07-26|18:16:20.992] Generating DAG in progress               epoch=1 percentage=21 elapsed=4m24.642s
>
> INFO [07-26|18:16:28.468] Successfully sealed new block            number=34 hash=bac86c…464986
>
> INFO [07-26|18:16:28.468]  block reached canonical chain          number=29 hash=d874de…8cfc4c
>
> INFO [07-26|18:16:28.469]  mined potential block                  number=34 hash=bac86c…464986
>
> INFO [07-26|18:16:28.489] Commit new mining work                   number=35 txs=0 uncles=0 elapsed=176.283µs
>
> INFO [07-26|18:16:28.541] Successfully sealed new block            number=35 hash=421888…c52d31
>
> **INFO [07-26|18:16:28.541]  block reached canonical chain          number=30 hash=2b5e96…217cae**
>
> **INFO [07-26|18:16:28.542]  mined potential block                  number=35 hash=421888…c52d31**
>
> **INFO [07-26|18:16:28.543] Commit new mining work                   number=36 txs=0 uncles=0 elapsed=500.117µs**
>
> INFO [07-26|18:16:33.099] Generating DAG in progress               epoch=1 percentage=22 elapsed=4m36.749s
>
> **INFO [07-26|18:16:34.837] Successfully sealed new block            number=36 hash=f9ec16…7feafa**
>
> **INFO [07-26|18:16:34.838]  block reached canonical chain          number=31 hash=0ef145…b841bb**
>
> **INFO [07-26|18:16:34.838]  mined potential block                  number=36 hash=f9ec16…7feafa**
>
> INFO [07-26|18:16:34.840] Commit new mining work                   number=37 txs=0 uncles=0 elapsed=498.032µs
>
> INFO [07-26|18:16:40.194] Successfully sealed new block            number=37 hash=549c29…93f046



> miner.stop()

​    查看挖矿日志，可以新开启一个终端，输入下面命令进行查看

> $ tail -f mychain.log

（4）转账

​    挖矿成功后，默认会将以太币给第一个账户，查看余额

> \>eth.getBalance(eth.accounts[0])
>
> 135000000000000000000

​    eth.accounts：以数组形式返回为账户列表，当前只有一个账户，可用eth.accounts[0]表示

​    新建一个账户，查看余额

> \>personal.newAccount("111")
>
> "0x866e46b71923d54e6fee26f42ececc216ca76c1d"
>
> \>eth.getBalance(eth.accounts[1])
>
> 0

​    从accounts[0]向accounts[1]发送1个以太币，首先需要解锁accounts[0]，使用如下命令

> \>personal.unlockAccount(eth.accounts[0],"123",0)
>
> true

​    三个参数分别为accounts[0]的地址，创建时的密码及解锁时间（毫秒为单位）。0 代表长时间解锁

转账操作

> \>eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[1],value:web3.toWei(1,'ether')})

​    开启挖矿等待区块验证，转账成功后，查询accounts[1]余额，默认单位是Wei

> \>eth.getBalance(eth.accounts[1])
>
> 100000000000000000

可以使用web3.fromWei()命令将wei转为ether。



# [以太坊客户端Geth命令用法-参数详解](https://www.cnblogs.com/tinyxiong/p/7918706.html)

## 命令用法

```
geth [选项] 命令 [命令选项] [参数…]
```

### 版本：

```
1.7.3-stable
```

### 命令:

```
account    管理账户
attach     启动交互式JavaScript环境（连接到节点）
bug        上报bug Issues
console    启动交互式JavaScript环境
copydb     从文件夹创建本地链
dump       Dump（分析）一个特定的块存储
dumpconfig 显示配置值
export     导出区块链到文件
import     导入一个区块链文件
init       启动并初始化一个新的创世纪块
js         执行指定的JavaScript文件(多个)
license    显示许可信息
makecache  生成ethash验证缓存(用于测试)
makedag    生成ethash 挖矿DAG(用于测试)
monitor    监控和可视化节点指标
removedb   删除区块链和状态数据库
version    打印版本号
wallet     管理Ethereum预售钱包
help,h     显示一个命令或帮助一个命令列表
```

### ETHEREUM选项:

```
--config value          TOML 配置文件
--datadir “xxx”         数据库和keystore密钥的数据目录
--keystore              keystore存放目录(默认在datadir内)
--nousb                 禁用监控和管理USB硬件钱包
--networkid value       网络标识符(整型, 1=Frontier, 2=Morden (弃用), 3=Ropsten, 4=Rinkeby) (默认: 1)
--testnet               Ropsten网络:预先配置的POW(proof-of-work)测试网络
--rinkeby               Rinkeby网络: 预先配置的POA(proof-of-authority)测试网络
--syncmode "fast"       同步模式 ("fast", "full", or "light")
--ethstats value        上报ethstats service  URL (nodename:secret@host:port)
--identity value        自定义节点名
--lightserv value       允许LES请求时间最大百分比(0 – 90)(默认值:0) 
--lightpeers value      最大LES client peers数量(默认值:20)
--lightkdf              在KDF强度消费时降低key-derivation RAM&CPU使用
```

### 开发者（模式）选项:

```
--dev               使用POA共识网络，默认预分配一个开发者账户并且会自动开启挖矿。
--dev.period value  开发者模式下挖矿周期 (0 = 仅在交易时) (默认: 0)
```

### ETHASH 选项:

```
--ethash.cachedir                        ethash验证缓存目录(默认 = datadir目录内)
--ethash.cachesinmem value               在内存保存的最近的ethash缓存个数  (每个缓存16MB ) (默认: 2)
--ethash.cachesondisk value              在磁盘保存的最近的ethash缓存个数 (每个缓存16MB) (默认: 3)
--ethash.dagdir ""                       存ethash DAGs目录 (默认 = 用户hom目录)
--ethash.dagsinmem value                 在内存保存的最近的ethash DAGs 个数 (每个1GB以上) (默认: 1)
--ethash.dagsondisk value                在磁盘保存的最近的ethash DAGs 个数 (每个1GB以上) (默认: 2)
```

### 交易池选项:

```
--txpool.nolocals            为本地提交交易禁用价格豁免
--txpool.journal value       本地交易的磁盘日志：用于节点重启 (默认: "transactions.rlp")
--txpool.rejournal value     重新生成本地交易日志的时间间隔 (默认: 1小时)
--txpool.pricelimit value    加入交易池的最小的gas价格限制(默认: 1)
--txpool.pricebump value     价格波动百分比（相对之前已有交易） (默认: 10)
--txpool.accountslots value  每个帐户保证可执行的最少交易槽数量  (默认: 16)
--txpool.globalslots value   所有帐户可执行的最大交易槽数量 (默认: 4096)
--txpool.accountqueue value  每个帐户允许的最多非可执行交易槽数量 (默认: 64)
--txpool.globalqueue value   所有帐户非可执行交易最大槽数量  (默认: 1024)
--txpool.lifetime value      非可执行交易最大入队时间(默认: 3小时)
```

### 性能调优的选项:

```
--cache value                分配给内部缓存的内存MB数量，缓存值(最低16 mb /数据库强制要求)(默认:128)
--trie-cache-gens value      保持在内存中产生的trie node数量(默认:120)
```

### 帐户选项:

```
--unlock value              需解锁账户用逗号分隔
--password value            用于非交互式密码输入的密码文件
```

### API和控制台选项:

```
--rpc                       启用HTTP-RPC服务器
--rpcaddr value             HTTP-RPC服务器接口地址(默认值:“localhost”)
--rpcport value             HTTP-RPC服务器监听端口(默认值:8545)
--rpcapi value              基于HTTP-RPC接口提供的API
--ws                        启用WS-RPC服务器
--wsaddr value              WS-RPC服务器监听接口地址(默认值:“localhost”)
--wsport value              WS-RPC服务器监听端口(默认值:8546)
--wsapi  value              基于WS-RPC的接口提供的API
--wsorigins value           websockets请求允许的源
--ipcdisable                禁用IPC-RPC服务器
--ipcpath                   包含在datadir里的IPC socket/pipe文件名(转义过的显式路径)
--rpccorsdomain value       允许跨域请求的域名列表(逗号分隔)(浏览器强制)
--jspath loadScript         JavaScript加载脚本的根路径(默认值:“.”)
--exec value                执行JavaScript语句(只能结合console/attach使用)
--preload value             预加载到控制台的JavaScript文件列表(逗号分隔)
```

### 网络选项:

```
--bootnodes value    用于P2P发现引导的enode urls(逗号分隔)(对于light servers用v4+v5代替)
--bootnodesv4 value  用于P2P v4发现引导的enode urls(逗号分隔) (light server, 全节点)
--bootnodesv5 value  用于P2P v5发现引导的enode urls(逗号分隔) (light server, 轻节点)
--port value         网卡监听端口(默认值:30303)
--maxpeers value     最大的网络节点数量(如果设置为0，网络将被禁用)(默认值:25)
--maxpendpeers value    最大尝试连接的数量(如果设置为0，则将使用默认值)(默认值:0)
--nat value             NAT端口映射机制 (any|none|upnp|pmp|extip:<IP>) (默认: “any”)
--nodiscover            禁用节点发现机制(手动添加节点)
--v5disc                启用实验性的RLPx V5(Topic发现)机制
--nodekey value         P2P节点密钥文件
--nodekeyhex value      十六进制的P2P节点密钥(用于测试)
```

### 矿工选项:

```
--mine                  打开挖矿
--minerthreads value    挖矿使用的CPU线程数量(默认值:8)
--etherbase value       挖矿奖励地址(默认=第一个创建的帐户)(默认值:“0”)
--targetgaslimit value  目标gas限制：设置最低gas限制（低于这个不会被挖？） (默认值:“4712388”)
--gasprice value        挖矿接受交易的最低gas价格
--extradata value       矿工设置的额外块数据(默认=client version)
```

### GAS价格选项:

```
--gpoblocks value      用于检查gas价格的最近块的个数  (默认: 10)
--gpopercentile value  建议gas价参考最近交易的gas价的百分位数，(默认: 50)
```

### 虚拟机的选项:

```
--vmdebug        记录VM及合约调试信息
```

### 日志和调试选项:

```
--metrics            启用metrics收集和报告
--fakepow            禁用proof-of-work验证
--verbosity value    日志详细度:0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail (default: 3)
--vmodule value      每个模块详细度:以 <pattern>=<level>的逗号分隔列表 (比如 eth/*=6,p2p=5)
--backtrace value    请求特定日志记录堆栈跟踪 (比如 “block.go:271”)
--debug                     突出显示调用位置日志(文件名及行号)
--pprof                     启用pprof HTTP服务器
--pprofaddr value           pprof HTTP服务器监听接口(默认值:127.0.0.1)
--pprofport value           pprof HTTP服务器监听端口(默认值:6060)
--memprofilerate value      按指定频率打开memory profiling    (默认:524288)
--blockprofilerate value    按指定频率打开block profiling    (默认值:0)
--cpuprofile value          将CPU profile写入指定文件
--trace value               将execution trace写入指定文件
```

### WHISPER实验选项:

```
--shh                        启用Whisper
--shh.maxmessagesize value   可接受的最大的消息大小 (默认值: 1048576)
--shh.pow value              可接受的最小的POW (默认值: 0.2)
```

### 弃用选项：

```
--fast     开启快速同步
--light    启用轻客户端模式
```

### 其他选项:

```
–help, -h    显示帮助
```

​                                                                                                                                                                                                                                2019/4/28/12:33

