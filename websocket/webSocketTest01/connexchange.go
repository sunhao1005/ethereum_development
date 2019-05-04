package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
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
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(address_erc20), conn)
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
func AddTokenAddress(addr string,tokenid uint64 ,name string,decimal uint64, trandtype bool) {
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//加载私钥
	privatekey,err:=crypto.HexToECDSA(privateKey)
	if err!=nil {
	fmt.Println("load privatekey error :",err)
		return
	}
    auth:=bind.NewKeyedTransactor(privatekey)

    //连接合约
    tokenme,err:=NewExchange(common.HexToAddress(address_exchange),conn)
	if err!=nil {
		fmt.Println("connect to contract error :",err)
		return
	}
    tx,err:=tokenme.ChangeContractAddress(auth,common.HexToAddress(addr),tokenid,name,decimal,trandtype)
	if err!=nil {
		fmt.Println(".ChangeContractAddress error :",err)
		return
	}
	fmt.Println("成功：", tx.Hash().String())
}

//授权代币
func ApproveTokenAmount(addr string, amount int64) (string,error){
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return"",err
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test_01)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return"",err
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return"",err
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(address_erc20), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return"",err
	}

	tx, err := tokenme.Approve(auth, common.HexToAddress(addr), big.NewInt(amount))
	if err != nil {
		fmt.Println("erc20合约授权失败：", err)
		return"",err
	}
	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return"",err
	}
	fmt.Println("交易数据：", hexutil.Encode(data))
	return hexutil.Encode(data),nil
}

//组装卖单
func BuildOrder( orderid uint64, useraddress string, receiverAddress string, tradetype bool, sellamount int64, price int64, buyaddress string, createtime int64) {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test_02)
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
	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接交易所合约合约失败：", err)
		return
	}
    auth.GasLimit=3000000
	auth.Value=computeAmount(sellamount)
	tx, err := tokenme.DepositsOrder(auth, orderid, common.HexToAddress(useraddress), receiverAddress, tradetype, computeAmount(sellamount), big.NewInt(price), common.HexToAddress(buyaddress), big.NewInt(createtime))
	if err != nil {
		fmt.Println("提交卖单失败：", err)
		return
	}

	//将签名后的交易转成[]byte
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return
	}
	//tx.MarshalJSON()
	fmt.Println("交易数据：", hexutil.Encode(data))
}

//
func computeAmount(amount int64) *big.Int {
	a := new(big.Int)
	a.Mul(big.NewInt(amount), math.BigPow(10,18))
	a.Div(a, math.BigPow(10,8))
	return a
}

//组装撤单
func BuildWithable(orderid int64) {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
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
	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
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
	primale := "132a55323963ea7f77ff812daad8ba1efcf4a3abc0fb75bfdc8fdf8b1fc4a58e"
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

	addrlocal := "0x7bf33E89E47AC84bD502dB307E70B9e9d6944bbf"
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
	conn, err := ethclient.Dial(address_eth_conn)
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
	tokenme, err := NewERC20(common.HexToAddress(address_erc20), conn)
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
func ExchangeTest(addr string,tokenid uint64,decimal uint64) {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
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
	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return
	}

	tx, err := tokenme.ChangeContractAddress(auth, common.HexToAddress(addr), tokenid,"TBC" ,decimal,true)
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
	conn, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA("d5f97b98ca45b9267dc7ea55fb868e39fdc86ef3ce3e83db7f7a90cceb12e4f3")
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
	tokenme, err := NewExchange(common.HexToAddress("0xac21e663cc41847b56a023a3ed32f7caf328004b"), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return
	}
	//转eth
	auth.GasLimit=3000000
	//auth.Value = big.NewInt(10000000)
	tx, err := tokenme.DepositsSell(auth, 1002, common.HexToAddress("0x9286673add2a52779771a1392b972101f69f0f85"), "0x9286673add2a52779771a1392b972101f69f0f85", false, big.NewInt(100), common.HexToAddress("0x424cb70dbb79736bf622e749fec8b981d2308da1"),big.NewInt(100), big.NewInt(time.Now().Unix()))
	if err != nil {
		fmt.Println("exchang合测试失败：", err)
		return
	}
	fmt.Println("success :", tx.Hash().String())
}

func GetAllonce(addr string)  {
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewERC20(common.HexToAddress(address_erc20), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}
	//
	balance, err := tokenme.Allowance(&bind.CallOpts{Pending: true}, common.HexToAddress(addr),common.HexToAddress(address_exchange))
	if err != nil {
		fmt.Println("获取余额失败：", err)
		return
	}
	fmt.Println("allownce :", balance.Uint64())
}

func BuildDepositOrder(orderid uint64 ,ammount int64,price int64)(string,error)  {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return"",err
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test_02)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return"",err
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return"",err
	}


	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return"",err
	}
	//转eth
	auth.GasLimit=3000000


	a :=new(big.Int)
	a.Mul(big.NewInt(ammount),math.BigPow(10,18))
	a.Div(a,math.BigPow(10,8))

	auth.Value = a
	tx, err := tokenme.DepositsOrder(auth, orderid, common.HexToAddress(address_test_02), address_test_02, false, a,big.NewInt(price), common.HexToAddress(address_erc20), big.NewInt(time.Now().Unix()))
	if err != nil {
		fmt.Println("exchang合测试失败：", err)
		return"",err
	}

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "",err
	}

	fmt.Println("交易数据：", hexutil.Encode(data))

	return hexutil.Encode(data),nil



	return tx.Hash().String(),nil
}

func BuildDepositBuyEosOrder(orderid uint64 ,ammount int64,price int64)(string,error)  {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return"",err
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test_02)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return"",err
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return"",err
	}


	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return"",err
	}
	//转eth
	auth.GasLimit=3000000


	a :=new(big.Int)
	a.Mul(big.NewInt(ammount),math.BigPow(10,18))
	a.Div(a,math.BigPow(10,8))

	auth.Value = a
	tx, err := tokenme.DepositsOrder(auth, orderid, common.HexToAddress(address_test_02), "inita", true, a,big.NewInt(price), common.HexToAddress(address_nill), big.NewInt(time.Now().Unix()))
	if err != nil {
		fmt.Println("exchang合测试失败：", err)
		return"",err
	}

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "",err
	}

	fmt.Println("交易数据：", hexutil.Encode(data))

	return hexutil.Encode(data),nil



	return tx.Hash().String(),nil
}




//查询map
func GetExchangeInfor(orderid uint64) {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}
	//
	order, err := tokenme.OrderidToUsers(&bind.CallOpts{Pending: true}, orderid)
	if err != nil {
		fmt.Println("OrderidToUsers：", err)
		return
	}
	fmt.Println("OrderidToUsers :", order)
}
func upstate(orderid uint64){
	conn, err := ethclient.Dial(address_eth_conn)
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


	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return
	}
	//转eth
	tx, err := tokenme.Upstate(auth,orderid)
	if err != nil {
		fmt.Println("更新合约失败：", err)
		return
	}
	fmt.Println("success :", tx.Hash().String())
}


func GetTokenInfo(tokenid uint64) {
	//连接以太坊
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	//创建erc20合约对象
	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}

	var auth =bind.CallOpts{Pending:true,From:common.HexToAddress(address_exchange_owner)}
	//
	a,b,c,d,e, err := tokenme.GetTokenInfor(&auth, tokenid)
	if err != nil {
		fmt.Println("获取余额失败：", err)
		return
	}
	fmt.Println("id :", a," name: ",b," address: ",c,d,e)
}

func  GetTokenTable() ([]EthTokenTable, error) {
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)

	}
	tokens := make([]EthTokenTable, 0)

	//连接交易所
	tokenme, err := ethereum.NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接交易所合约失败：", err)
		return nil, err
	}

	auth := bind.CallOpts{Pending: true, From: common.HexToAddress(address_exchange_owner)}
	var i uint64 = 1

	for {
		var a EthTokenTable
		_,b,c, d, e, err :=tokenme.GetTokenInfor(&auth, i)
		if err != nil {
			fmt.Println("tokenme.GetTokenInfor error:", err)
			return nil, err
		}

		if b != "" {
			a.Symbol = b
			a.TokenID = i
			a.TokenContract = c.String()
			a.Decimal =uint(d)
			a.IsCancle=e
			tokens = append(tokens, a)
		} else {
			break
		}
		i++
	}
	return tokens, nil
}
type EthTokenTable struct {
	TokenID       uint64          `json:"id"`
	Symbol        string          `json:"symbol"`
	TokenContract string `json:"tcontract"`
	Decimal       uint            `json:"decilmal"`
	IsCancle      bool            `json:"is_cancle"`
}

func BuildEthWidthOrder(id uint64)(string,error)  {
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return"",err
	}
	//加载私钥，构造TransactOpts对象
	piivatekey, err := crypto.HexToECDSA(privateKey_test_01)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return"",err
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return"",err
	}

	tokenme, err := NewExchange(common.HexToAddress(address_exchange), conn)
	if err != nil {
		fmt.Println("连接exchange合约失败：", err)
		return"",err
	}

	tx,err:=tokenme.Withdrawal(auth,id)
	if err!=nil {
		return "",err
	}

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "",err
	}

	fmt.Println("交易数据：", hexutil.Encode(data))

	return hexutil.Encode(data),nil

}