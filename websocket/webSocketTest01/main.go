package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

/*type ReqBody struct {
	Transaction *eos.Transaction `json:"transaction"`
	Signatures  []ecc.Signature  `json:"signatures"`
}*/

// 接收数据
/*type content struct {
	Request_id string
	Code       int
	Msg        string
	Data       string
}*/
type BizMessage struct {
	//MsgType int			`json:"msg_type"`
	Namespace string      `json:"namespace"` // type消息类型
	Action    string      `json:"action"`
	Data      interface{} `json:"data"` // data数据字段
}

/*const (
	BUY  = true
	SELL = false
)*/

func main() {
  sendEThSellOrder()
 //sendEthBuyOrder()
 //sendEthBUYEOsorder()
// sendEthWidthorder()
}


//撤单
func sendEthWidthorder(){
	dataMap,err:=BuildWidthorder(1)
	if err!=nil {
		return
	}
	fmt.Println(dataMap)
	sendWS("user", "withdraweth", dataMap)

}

//跨链
func sendEthBUYEOsorder()  {
	dataMap,err:=BuildBuyEosOrder(11111108,1,100)
	if err!=nil {
		return
	}
	fmt.Println(dataMap)
	sendWS("user", "sendeth", dataMap)
}



func sendEthBuyOrder()  {
	dataMap,err:=BuildBuyOrder(1009,1,0.002)
	if err!=nil {
		return
	}
	fmt.Println(dataMap)
	sendWS("user", "sendeth", dataMap)
}

//卖单
func sendEThSellOrder(){
	dataMap,err:=BuildSellOrder(10003,1000,0.005)
	if err!=nil {
		return
	}
	fmt.Println(dataMap)
	sendWS("user", "sendeth", dataMap)


}





func sendWS(namespace, action string, data map[string]interface{}) {
	u := url.URL{Scheme: "ws", Host: "47.102.110.128:8082", Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	//var r BizMessage
	// var rs string

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()
	me := BizMessage{Namespace: namespace, Action: action, Data: data}

	//go func() {
		err = c.WriteJSON(me)
		if err != nil {
			log.Println("write-----:", err)
			return
		}
	//}()
	//
	//go func() {
	//	err := c.ReadJSON(rs)
	//	if err != nil {
	//		log.Println("read------", err)
	//		return
	//	}
	//}()
	time.Sleep(100 * time.Second)
}


/*func NewBuyTransfer(selltokenMoney, id, price string) (*eos.Action, error) {
	from := eos.AN("initb")
	to := eos.AN("tbcexchange1")

	contract := eos.AN("eosio.token")
	sendAsset, err := eos.NewAsset(selltokenMoney)
	if err != nil {
		fmt.Println("NewAsset err----:", err)
		return nil, err
	}
	//transfer := contract2.NewTransfer(from, to, contract, sendAsset, "{\"i\":\"223456789\",\"a\":\"0\",\"p\":\"10000000000\",\"bs\":\"ETH,18\",\"bc\":\"0x0000000000000000000000000000000000000000\",\"ta\":\"0x60C278c444f4a0a65441f1bE51a1B63DE608FBC0\"}")
	transfer := contract2.NewTransfer(from, to, contract, sendAsset, "{\"i\":\""+id+"\",\"a\":\"1\",\"p\":"+price+",\"bs\":\"CCC,4\",\"bc\":\"initb\",\"ta\":\"initb\"}")
	return transfer, nil
}*/

/*func NewSellTransfer(selltokenMoney, id, price string) (*eos.Action, error) {
	from := eos.AN("initb")
	to := eos.AN("tbcexchange1")

	contract := eos.AN("initb")
	sendAsset, err := eos.NewAsset(selltokenMoney)
	if err != nil {
		fmt.Println("NewAsset err----:", err)
		return nil, err
	}
	//transfer := contract2.NewTransfer(from, to, contract, sendAsset, "{\"i\":\"223456789\",\"a\":\"0\",\"p\":\"10000000000\",\"bs\":\"ETH,18\",\"bc\":\"0x0000000000000000000000000000000000000000\",\"ta\":\"0x60C278c444f4a0a65441f1bE51a1B63DE608FBC0\"}")
	transfer := contract2.NewTransfer(from, to, contract, sendAsset, "{\"i\":\""+id+"\",\"a\":\"1\",\"p\":"+price+",\"bs\":\"SYS,4\",\"bc\":\"eosio.token\",\"ta\":\"initb\"}")
	return transfer, nil
	//a := con.NewTransfer(eos.AN(from), eos.AN(to), q, memo)
	//fmt.Println(string(s))
	//txstr := make([]byte, 0)
	//for k, v := range s { //   ":34   {:123  }:125    \:92
	//	if k == len(s)-18 {
	//		txstr = append(txstr, v, 34)
	//		continue
	//	}
	//	if k > len(s)-20 {
	//		txstr = append(txstr, v)
	//		continue
	//	}
	//	if v == 34 && k > 8 && k != len(s)-2 {
	//		txstr = append(txstr, 92)
	//		//s[k] = 92+34
	//	}
	//	if k == 9 {
	//		txstr = append(txstr, 34, 123)
	//		continue
	//	}
	//
	//	txstr = append(txstr, v)
	//}
}*/


/*func NewTxOptions() *eos.TxOptions {
	opt := &eos.TxOptions{}
	eosurl := "http://47.102.110.128:8888"

	api := client.NewClient(eosurl)
	err := opt.FillFromChain(api.API)
	if err != nil {
		return nil
	}
	return opt
}

func BuildTx(a *eos.Action) (*table.EosPutOrderData, error) {
	action := make([]*eos.Action, 1)
	action[0] = a

	opts := NewTxOptions()
	tx := eos.NewTransaction(action, opts)
	sigtx := eos.NewSignedTransaction(tx)
	key := eos.NewKeyBag()
	err := key.ImportPrivateKey("5KipT9AhAdx5vxCb7T9gGmSjVkFL962WacCAsdonkdk4L11X7v3")
	if err != nil {
		return nil, err
	}
	pub, _ := ecc.NewPublicKey("EOS61Jat8Y83RRS4Pp1cLUe48eDpdxRX8W4pKc5iLUzcQNogcmzwD")
	signTx, err := key.Sign(sigtx, opts.ChainID, pub)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(ReqBody{
		Transaction: tx,
		Signatures:  signTx.Signatures,
	})

	if err != nil {
		fmt.Println("json.Marshal error-------:", err)
		return nil, err
	}

	signedTx := &table.EosPutOrderData{
		Order:      string(bytes),
		MarketCode: "1",
		ChainCode:  "1",
		TargetCode: "1",
	}

	//ret, _ := json.Marshal(signedTx)

	return signedTx, nil
}*/


