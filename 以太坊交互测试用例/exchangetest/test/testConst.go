package main

const (
	ETH_CONN_ADDRESS          = "https://ropsten.infura.io/v3/4a74edfffce34c07bc80a78e3eb0874c" //以太坊网络连接地址（需注册才能允许通行）
	CONTRACT_ADDRESS_EXCHANGE = "0x910230c35043D8E0dF3792d4765043b94c5efAc1"                    //交易所合约地址
	CONTRACT_ADDRESS_ERC20    = "0x9ee54be4b6b8163226b357609b03b3de396b89d2"                    //token合约地址 //10 0000 0000
	ADDRESS_ME                = "0x94066127301abf11e1594cCaD1582A4444e51c4d"                    //交易所合约部署账户
	ADDRESS_ERC20             = "0x94066127301abf11e1594cCaD1582A4444e51c4d"                    //token合约部署账户
	ADDRESS_NILL              = "0x0000000000000000000000000000000000000000"
)

//账户keystore
const keystore = `{"address":"94066127301abf11e1594ccad1582a4444e51c4d","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"259e8835c732247ff24cf2c46700817d"},"ciphertext":"351d82ecea362123085b28be7a1de3fb2ef4fc1d93eb997573604dd2aef40767","kdf":"scrypt","kdfparams":{"dklen":32,"n":65536,"p":1,"r":8,"salt":"a107ebec40763392e747179e137d36755609ed289fa3e0e34d610c916d078104"},"mac":"53f8dd3c8b169b0641ac2cdfd5fbe8e01891fb4a864fa0fd7f691d8c2acac778"},"id":"e79c3a89-a263-4fac-8156-68ddc1f57179","version":3}
`
const password = "17719317036" //钱包密码
const privateKey = "71ACF79F0760F62EDB6CBE995D342A357A392860881A94DF782EEC6C27028FF2"
const address_test = "0x2daF64B86d5b4aa224852EC0B7C1142557ad15fb"
const privateKey_test = "A0D2ECC4C6F1255A77A56CEB6B9B00313FC327CD3A7DB088074D089A7BD83517"
const address_test02 = "0x8d91134CC5800A4472402adC2Ef17583e0858487"

const erc_abi = `[{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`
