## ganache-cli命令行参数说明

ganache-cli是以太坊节点仿真器软件ganache的命令行版本，可以方便开发者快速进行以太坊DApp的开发与测试。 本文将详细介绍ganache-cli的安装及命令行参数的作用。



## 安装

```
npm install -g ganache-cli
```



## 环境配置

```
sudo npm install -g ganache-cli                            查看安装位置

如果文件安装在   /usr/local/node-v8.9.3-linux-x64/bin 
修改环境变量     sudo vim /etc/profile

export GANACHE=/usr/local/node-v8.9.3-linux-x64/bin
export GOROOT=/usr/local/go
export GOPATH=$GOROOT/bin
export PATH=$PATH:$GOPATH:$GANACHE

在profile文件第一行添加运行文件地址
在最后一行添加$GANACHE

source /etc/profile   执行配置并退出
 
```



## 启动

```
~$ ganache-cli                                                         基础命令
~$ ganache-cli -a -10 -e -100 -p 7545 -h 0.0.0.0 -db /root/ganache -debug  
~$ nohup ganache-cli -a -10 -e -100 -p 7545 -h 0.0.0.0 -db /root/ganache -debug >>/root/ganachenoubh.out >2&1 &                                         挂载服务
```



## 启动选项

- -a 或 –accounts： 指定启动时要创建的测试账户数量。
- -e 或 –defaultBalanceEther： 分配给每个测试账户的ether数量，默认值为100。
- -b 或r –blockTime： 指定自动挖矿的blockTime，以秒为单位。默认值为0，表示不进行自动挖矿。
- -d 或 –deterministic： 基于预定的助记词（`mnemonic`）生成固定的测试账户地址。
- -n 或 –secure： 默认锁定所有测试账户，有利于进行第三方交易签名。
- -m 或 –mnemonic： 用于生成测试账户地址的助记词。
- -p 或 –port： 设置监听端口，默认值为8545。
- -h 或 –hostname： 设置监听主机，默认值同NodeJS的`server.listen()`。
- -s 或 –seed： 设置生成助记词的种子。.
- -g 或 –gasPrice： 设定Gas价格，默认值为20000000000。
- -l 或 –gasLimit： 设定Gas上限，默认值为90000。
- -f 或 –fork： 从一个运行中的以太坊节点客户端软件的指定区块分叉。输入值应当是该节点旳HTTP地址和端口，例如`http://localhost:8545`。 可选使用@标记来指定具体区块，例如：`http://localhost:8545@1599200`。
- -i 或 –networkId：指定网络id。默认值为当前时间，或使用所分叉链的网络id。
- –db： 设置保存链数据的目录。如果该路径中已经有链数据，ganache-cli将用它初始化链而不是重新创建。
- –debug：输出VM操作码，用于调试。
- –mem：输出ganache-cli内存使用统计信息，这将替代标准的输出信息。
- –noVMErrorsOnRPCResponse：不把失败的交易作为RCP错误发送。开启这个标志使错误报告方式兼容其他的节点客户端，例如geth和Parity。

## 特殊选项

- –account： 指定账户私钥和账户余额来创建初始测试账户。可多次设置：

  ```
  $ ganache-cli --account="<privatekey>,balance" [--account="<privatekey>,balance"]
  ```

注意私钥长度为64字符，必须使用0x前缀的16进制字符串。账户余额可以是整数，也可以是0x前缀的17进制字符串，单位为wei。

使用–account选项时，不会自动创建HD钱包。

- -u 或 –unlock： 解锁指定账户，或解锁指定序号的账户。可以设置多次。当与–secure选项同时使用时，这个选项将改变指定账户的锁定状态：

```
$ ganache-cli --secure --unlock "0x1234..." --unlock "0xabcd..."
```

也可以指定一个数字，按序号解锁账号：

```
$ ganache-cli --secure -u 0 -u 1
```