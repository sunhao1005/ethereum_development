package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
)

func TestDecodeTx(tx string) {

	//构造交易对象
	var trasaction types.Transaction

	//解析交易
	input, err := hexutil.Decode(tx)
	if err != nil {
		fmt.Println("hexutil.Decode", err)
		return
	}

	//解析交易到对象
	err = rlp.DecodeBytes(input, &trasaction)
	if err != nil {
		fmt.Println("rlp.DecodeBytes", err)
		return
	}
	fmt.Println("交易对象：",trasaction)
   fmt.Println(trasaction.Data())



	//解析到txdata
	var txdata txdata
	err = rlp.DecodeBytes(input, &txdata)
	if err != nil {
		fmt.Println("txdata rlp.DecodeBytes", err)
		return
	}
	fmt.Println("txdata对象：",txdata)
	fmt.Println(txdata.Payload)
	//构造abi
	abi, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		fmt.Println("abi.JSON error:", err)
		return
	}

	meth, ok := abi.Methods["changeContractAddress"]
	fmt.Println("abi.Methods", meth, ok)

   a,_:= meth.Inputs.UnpackValues(trasaction.Data()[4:])
   fmt.Println("Inputs.UnpackValues",a)

   fmt.Println("解析后input结构体",ChangeType(a).addr.String())
	var trabs tokentrabsfer
	//解析txdata.Payload
	fmt.Println(len(trasaction.Data()))
	err = abi.Unpack(&trabs, "changeContractAddress",trasaction.Data())
	//fmt.Println(txdata.Payload[4:])
	//err=meth.Inputs.Unpack(&trabs,trasaction.Data())
	if err != nil {
		fmt.Println("abi.Unpack error:", err)
		return
	}

	fmt.Println(trabs)

}

type tokentrabsfer struct {
	addr   common.Address
	types string
	changetype bool
}

type txdata struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

func ChangeType(v []interface{})tokentrabsfer  {

	return tokentrabsfer{
		addr:v[0].(common.Address),
		types:v[1].(string),
		changetype:v[2].(bool),
	}

}