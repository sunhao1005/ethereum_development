package main

import (
	"fmt"
	"math"
	"strconv"
)

func BuildSellOrder(id uint64, amount int64, price float64) (map[string]interface{}, error) {

	//组装approve
	tx, err := ApproveTokenAmount(address_exchange, amount)
	if err != nil {
		return nil, err
	}

	datamap := make(map[string]interface{})
	datamap["orderId"] = strconv.FormatUint(id, 10)
	datamap["userAddress"] = address_test_01
	datamap["tradeType"] = true
	datamap["marketType"] = 2
	datamap["sellTokenSymbol"] = "DDD"
	datamap["sellTokenAmount"] = amount * 100000000
	datamap["sellTokenAddress"] = address_erc20
	datamap["price"] = uint64(price * 100000000)
	datamap["targetAddress"] = address_test_01
	datamap["approveTx"] = tx
	return datamap,nil
}

func BuildBuyOrder(id uint64, amount int64, price float64 )(map[string]interface{},error){
    fmt.Println("price:",int64(price*math.Pow(10,8)))
	tx,err:=BuildDepositOrder(id,amount *100000000,int64(price*math.Pow(10,8)))
	if err!=nil {
		fmt.Println("BuildDepositOrder error ：",err)
		return nil,err
	}
	datamap := make(map[string]interface{})
	/*datamap["orderId"] = ""
	datamap["userAddress"] = ""
	datamap["tradeType"] = ""*/
	datamap["marketType"] =2
	/*datamap["sellTokenSymbol"] = ""
	datamap["sellTokenAmount"] = ""
	datamap["sellTokenAddress"] = ""
	datamap["price"] = ""
	datamap["targetAddress"] = ""*/
	datamap["approveTx"] = tx

	return datamap,nil
}

func BuildBuyEosOrder(id uint64, amount int64, price float64 )(map[string]interface{},error){
	fmt.Println("price:",int64(price*math.Pow(10,8)))
	tx,err:=BuildDepositBuyEosOrder(id,amount *100000000,int64(price*math.Pow(10,8)))
	if err!=nil {
		fmt.Println("BuildDepositOrder error ：",err)
		return nil,err
	}
	datamap := make(map[string]interface{})
	/*datamap["orderId"] = ""
	datamap["userAddress"] = ""
	datamap["tradeType"] = ""*/
	datamap["marketType"] =4
	/*datamap["sellTokenSymbol"] = ""
	datamap["sellTokenAmount"] = ""
	datamap["sellTokenAddress"] = ""
	datamap["price"] = ""
	datamap["targetAddress"] = ""*/
	datamap["approveTx"] = tx

	return datamap,nil
}

func BuildWidthorder(id uint64)(map[string]interface{},error)  {

	tx,err:=BuildEthWidthOrder(id)
	if err!=nil {
		fmt.Println("BuildWidthorder error ：",err)
		return nil,err
	}
	datamap := make(map[string]interface{})
	datamap["orderId"] = strconv.FormatUint(id,10)
	datamap["approveTx"] = tx
	return datamap,nil
}