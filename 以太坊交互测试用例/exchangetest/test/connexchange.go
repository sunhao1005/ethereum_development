package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/zhouyasong/tbcexchangetest/ethereum"
	"math/big"
	"time"
)

//查询erc20余额
func GetBalanceErc20(addr string) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(CONTRACT_ADDRESS_ERC20), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}
	//
	balance, err := tokenme.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(addr))
	if err != nil {
		fmt.Println("获取余额失败：", err)
		return
	}
	fmt.Println("balance :", balance.Uint64())
}

//添加代币
func AddTokenAddress(addr string, name string, trandtype bool) {
	conn, err := ethereum.NewClient(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	tx, err := conn.ChangeToken(addr, name, trandtype)
	if err != nil {
		fmt.Println("添加代币失败:", err)
		return
	}

	fmt.Println("成功：", tx)
}

//授权代币
func ApproveTokenAmount(addr string, amount int64) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(CONTRACT_ADDRESS_ERC20), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}

	tx, err := tokenme.Approve(auth, common.HexToAddress(addr), big.NewInt(amount))
	if err != nil {
		fmt.Println("erc20合约授权失败：", err)
		return
	}
	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return
	}
	tx.MarshalJSON()
	fmt.Println("交易数据：", hexutil.Encode(data))
}

//组装卖单
func BuildOrderSell(tract string, orderid uint64, useraddress string, receiverAddress string, tradetype bool, sellamount int64, price int64, buyaddress string, createtime int64) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建交易所合约对象
	tokenme, err := NewExchange(common.HexToAddress(CONTRACT_ADDRESS_EXCHANGE), conn)
	if err != nil {
		fmt.Println("连接交易所合约合约失败：", err)
		return
	}

	tx, err := tokenme.DepositsSell(auth, orderid, common.HexToAddress(useraddress), common.HexToAddress(receiverAddress), tradetype, computeAmount(sellamount), big.NewInt(price), common.HexToAddress(buyaddress), big.NewInt(createtime))
	if err != nil {
		fmt.Println("提交卖单失败：", err)
		return
	}

	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return
	}
	tx.MarshalJSON()
	fmt.Println("交易数据：", hexutil.Encode(data))
}

//
func computeAmount(amount int64) *big.Int {
	a := new(big.Int)
	a.Mul(big.NewInt(amount), big.NewInt(1*10^18))
	a.Div(a, big.NewInt(1*10^8))
	return a
}

//组装撤单
func BuildWithable(orderid int64) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建交易所合约对象
	tokenme, err := NewExchange(common.HexToAddress(CONTRACT_ADDRESS_EXCHANGE), conn)
	if err != nil {
		fmt.Println("连接交易所合约合约失败：", err)
		return
	}

	tx, err := tokenme.Withdrawal(auth, uint64(orderid))
	if err != nil {
		fmt.Println("撤销订单失败：", err)
		return
	}

	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return
	}
	tx.MarshalJSON()
	fmt.Println("交易数据：", hexutil.Encode(data))
}

//直接发送组装好数据到链上

func SendTxToEth(str string, conn *ethereum.Client) {
	//conn, err := ethereum.NewClient(ETH_CONN_ADDRESS)
	//if err != nil {
	//	fmt.Println("new client error", err)
	//	return
	//}
	tx, err := conn.SendSignedTransaction(str)
	if err != nil {
		fmt.Println("conn.SendSignedTransaction", err)
		return
	}
	fmt.Println("交易成功：", tx)

}

//组装erc20 transfer

func BuildTransfer(addrTo string, amount int64, conn *ethclient.Client) string {
	//连接以太坊
	//conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	//if err != nil {
	//	fmt.Println("连接以太坊失败:", err)
	//	return ""
	//}
	//加载私钥，构造TransactOpts对象
	primale:="132a55323963ea7f77ff812daad8ba1efcf4a3abc0fb75bfdc8fdf8b1fc4a58e"
	piivatekey, err := crypto.HexToECDSA(primale)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return ""
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return ""
	}

	 addrlocal :="0x7bf33E89E47AC84bD502dB307E70B9e9d6944bbf"
	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(addrlocal), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return ""
	}
	fmt.Println(time.Now())
	tx, err := tokenme.Transfer(auth, common.HexToAddress(addrTo), big.NewInt(amount))
	if err != nil {
		fmt.Println("erc20合约授权失败：", err)
		return ""
	}
	fmt.Println(time.Now())




	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return ""
	}

	fmt.Println("交易数据：", hexutil.Encode(data))

	return hexutil.Encode(data)
}

//erc20转账测试
func TokenTransfer(addr string, amount int64) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(CONTRACT_ADDRESS_ERC20), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}

	tx, err := tokenme.Transfer(auth, common.HexToAddress(addr), big.NewInt(amount))
	if err != nil {
		fmt.Println("erc20合约转账失败：", err)
		return
	}
	fmt.Println("success :", tx.Hash().String())
}

//exchange错误数据测试
func ExchangeTest(addr string) {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewExchange(common.HexToAddress(CONTRACT_ADDRESS_EXCHANGE), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return
	}

	tx, err := tokenme.ChangeContractAddress(auth, common.HexToAddress(addr), "TBC", true)
	if err != nil {
		fmt.Println("exchang合测试失败：", err)
		return
	}

	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return
	}

	fmt.Println("交易数据：", hexutil.Encode(data))
	fmt.Println("success :", tx.Hash().String())
}

//提交订单错误数据测试
func ExchangeTestOrderError() {
	//连接以太坊
	conn, err := ethclient.Dial(ETH_CONN_ADDRESS)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewExchange(common.HexToAddress(CONTRACT_ADDRESS_EXCHANGE), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return
	}
	//转eth
	auth.Value = big.NewInt(10000000)
	tx, err := tokenme.DepositsSell(auth, 10001, common.HexToAddress(ADDRESS_ME), common.HexToAddress(ADDRESS_ME), true, big.NewInt(10000000), big.NewInt(100), common.HexToAddress(ADDRESS_NILL), big.NewInt(time.Now().Unix()))
	if err != nil {
		fmt.Println("exchang合测试失败：", err)
		return
	}
	fmt.Println("success :", tx.Hash().String())
}
