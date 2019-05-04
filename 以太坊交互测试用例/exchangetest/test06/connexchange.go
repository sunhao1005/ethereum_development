package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func  GetTokenTable() ([]EthTokenTable, error) {
	conn, err := ethclient.Dial(address_eth_conn)
	if err != nil {
		fmt.Println("连接以太坊失败:", err)

	}
	tokens := make([]EthTokenTable, 0)

	//连接交易所
	tokenme, err := NewExchange(common.HexToAddress(ethereum.CONTRACT_ADDRESS_EXCHANGE), conn)
	if err != nil {
		fmt.Println("连接交易所合约失败：", err)
		return nil, err
	}

	auth := bind.CallOpts{Pending: true, From: common.HexToAddress(ADDRESS_ME)}
	for {
		var a EthTokenTable
		var i uint64 = 1
		_, name, address, err := tokenme.GetTokenInfor(&auth, i)
		if err != nil {
			fmt.Println("tokenme.GetTokenInfor error:", err)
			return nil, err
		}

		if name != "" {
			a.Symbol = name
			a.TokenID = i
			a.TokenContract = address.String()
			//a.Decimal =
			//a.IsCancle=
			tokens = append(tokens, a)
		} else {
			break
		}
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

