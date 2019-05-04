package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

)

const (
	keystore1 =`{"address":"7e16a073c4668fef9788e8e41a413d5d34155eed","crypto":{"cipher":"aes-128-ctr","ciphertext":"91d63948f969999f8087f670629aa6ee5adfca69a2891b0aca5d13635e252770","cipherparams":{"iv":"253afba7a8ffc4d2b01a0f5a4faeb187"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d3a36b9f98859b1b8c29e3fadc568a14e5560e6768f5eae020e3377aa642a66b"},"mac":"4e458a50daa5cf1f5427ced7994f52b15edad96db3c86e4d9362f057634d26a4"},"id":"2b7a0d62-f495-49e3-8d59-defad494c9a8","version":3}`
	keystore2 =`{"address":"9e97aa9e9ae32b2235ca1680e93783fa5d6e6c7b","crypto":{"cipher":"aes-128-ctr","ciphertext":"385501ef6dda493da762cf4867750e8a7c78a4872a77d26e55be45c01e2a7741","cipherparams":{"iv":"aca3c79a153de2818e123eb3046c3833"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bf62134017f0956cafe4705046a31ad1bf1d573221453ad769964f6c3476a0cd"},"mac":"5803047d402e2e33cbfb47c3628acada5e8b7dff4415da0c8f192db9147c4b51"},"id":"27185be9-31cd-4fd2-be8d-84ea95e228b4","version":3}`
	keystore3 =`{"address":"fe3dd72c9c5d62c35647306fb71ea0b56202897e","crypto":{"cipher":"aes-128-ctr","ciphertext":"0a87fe700b021a8fe6401f8f823b4543cdbdc4683e11fe3d84c12401dd576a74","cipherparams":{"iv":"5c39b99fac7593e5cdcab3a64a59fd6c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"aa9dbbb8be4ebc7920c8c36fc9197fd354234d8d0f629d3b0a5867a54c61d5b9"},"mac":"93bca4183bb31e256dfc646681ee6b71036664bc0972cfad0368e598f97fcf8d"},"id":"19fba4c8-6cda-465f-9c28-fbcf5bdb396d","version":3}`

	exchange_address =`0xa86e302708d43ee01e4e95daf184f5338988c392`
	erc20_address =`0xaf32f53c11b2c5ad1f95f332cd58d5d885e793c4`
    publickey1 =`0x7E16A073C4668Fef9788E8E41A413D5d34155eEd`
    primarykey1 =`E8019DB21AB0A4DF09869F040B4EE75A2DF3E81D19B51C1864E381163824E84F`


    publickey2 =`0x9e97aa9e9aE32b2235cA1680E93783Fa5d6E6c7b`
	primarykey2 =`F4EA2BA745C41C9738E7FE38346C742FE353CC8FDD7C044C09484806C0ECDAD6`

	publickey3 =`0xd433d3cD2f3cb19B35EA6F8A4d56e9D4ACbA05dE`
	primarykey3 =`CFD7CFA4F38474ED0575BFE31E2DC6AF71A23212CF659199798E7A4F35862E04`

    )

func main()  {
//deployContract()
//Check()bind.NewTransactor()


GetBalanceErc20(address_test_01)
GetBalanceErc20(address_test_02)
//AddTokenAddress(address_erc20,1,"TBC",0,true)
 t,_:=GetTokenTable()
 fmt.Println(t)

}


















func deployContract(){
	conn, err := ethclient.Dial("http://47.102.110.128:8545")
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	piivatekey, err := crypto.HexToECDSA(primarykey1)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}

	auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}

	a,b,c,d:= DeployExchange(auth,conn)
	if d!=nil {
		fmt.Println(d)
		return
	}
	fmt.Println("address:",a.String())
	fmt.Println("transaction:",b.Hash().String())
	addr,_:=c.ContractOwner(&bind.CallOpts{Pending:true})
	//fmt.Println("exchange",c.ContractOwner(bind.CallOpts{Pending:true}))
    fmt.Println(addr.String())
}
//geth/mychain
// geth --networkid 3 --datadir mychain --identity "mychain" --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi "admin,debug,eth,miner,net,personal,shh,txpool,web3" --port 30303 --rpccorsdomain "*" --nodiscover console 2>>mychain.log

func Check()  {
	conn, err := ethclient.Dial("http://47.102.110.128:8545")
	if err != nil {
		fmt.Println("连接以太坊失败:", err)
		return
	}

	/*piivatekey, err := crypto.HexToECDSA(ppkey1)
	if err != nil {
		fmt.Println("加载私钥失败：", err)
		return
	}*/

	/*auth := bind.NewKeyedTransactor(piivatekey)
	if err != nil {
		fmt.Println("构建交易失败", err)
		return
	}*/

	  tokenme,err:=NewExchange(common.HexToAddress(exchange_address),conn)
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}
	 address,err :=tokenme.ContractOwner(&bind.CallOpts{Pending:true,From:common.HexToAddress(publickey1)})
	if err != nil {
		fmt.Println("连接erc20合约失败：", err)
		return
	}
	 fmt.Println("成功：",address.String())
}

