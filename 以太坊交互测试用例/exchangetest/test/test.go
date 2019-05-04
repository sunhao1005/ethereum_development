
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"strings"
)

func TestDecodeJsonTx(tx string) {

	//构造交易对象
	var tras types.Transaction

	data,err:=hexutil.Decode(tx)
	if err != nil {
		fmt.Println("hexutil.Decode", err)
		return
	}
	//解析交易
	err = tras.UnmarshalJSON(data)
	if err != nil {
		fmt.Println("hexutil.Decode", err)
		return
	}

	fmt.Println(tras)
	//解析data
	abi, err := abi.JSON(strings.NewReader(erc_abi))


	meth, ok := abi.Methods["transfer"]
	fmt.Println("abi.Methods", meth, ok)


	var txdata txdata
	rlp.DecodeBytes(tras.Data(),&txdata)
	if err != nil {
		fmt.Println("rlp.DecodeBytes error:", err)
		return
	}
	fmt.Println(txdata)
	//err= meth.Inputs.Unpack(&trabs,tras.Payload[4:])
	trabs := tokentrabsfer{}
	//var trabs map[string]interface{}


	err = abi.Unpack(&trabs, "transfer",tras.Data()[36:])
	if err != nil {
		fmt.Println("abi.Unpack error:", err)
		return
	}

	fmt.Println(trabs)
	/*fmt.Println(trabs.to.Hash().String())
	fmt.Println(trabs.value.Int64())*/

}



