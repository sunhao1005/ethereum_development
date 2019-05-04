pragma solidity >=0.4.0 <0.7.0;
//*************************
//abigent 0.4.25
//SafeMath库
library SafeMath {

    /**
    * @dev Multiplies two numbers, throws on overflow.
    */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
        uint256 c = a * b;
        assert(c / a == b);
        return c;
    }

    /**
    * @dev Integer division of two numbers, truncating the quotient.
    */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    /**
    * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    /**
    * @dev Adds two numbers, throws on overflow.
    */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}

/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
interface ERC20 {
  function totalSupply() external view returns (uint256);

  function balanceOf(address who) external view returns (uint256);

  function allowance(address owner, address spender)
    external view returns (uint256);

  function transfer(address to, uint256 value) external returns (bool);

  function approve(address spender, uint256 value)
    external returns (bool);

  function transferFrom(address from, address to, uint256 value)
    external returns (bool);

  event Transfer(
    address indexed from,
    address indexed to,
    uint256 value
  );

  event Approval(
    address indexed owner,
    address indexed spender,
    uint256 value
  );
}

contract exchange{
   using SafeMath for uint256;

   //contract owner
   address   public contractOwner;

   //start contract
   bool public startContract;


   //the contract mutex
   bool contractMutex =false;

   //orders struct
   struct orders{
       uint64  order_id; //transcation txid
       address user_address; //user address
       address user_receiver;
       bool    trade_type; // trade type
       uint256 sell_amount; //amount of seel token
       address sell_address;
       uint256 price; //price=seel/buy
       uint256 buy_amount;  //amount of want to buy
       address buy_address; // contract address
       uint256 fees;   //
       bool withdrawalable; //Determine if the user can widthdrawal
       bool ex_success;
       uint256 create_at;
       uint256 update_at;
   }

   struct tokens{
        uint64   token_id;
        string    token_name;
        address   token_address;
   }

   //mapping of exchange
   mapping(uint64 => orders)public  orderidToUsers;
  //
   mapping(address => bool)public contractTypeAddress;
   mapping(uint64 =>tokens)internal tokenTable;
   //init data
   constructor()public payable{
      contractOwner=msg.sender;
      startContract=true;
   }

   //event listing

   event DepositsSell(uint64 orderid, address receiver,bool tradetype, uint256 sellamount,address selladdress,uint256 price,uint256 buyamount,uint256 createTime);
   event Deposits(uint64 orderid,bool tradetype, uint256 sellamount,uint256 price,uint256 buyamount,address buyaddress,uint256 createTime);
   event Withdrawl(uint64 txid);
   event Submit(uint64 txid,address to,uint256 sell,uint256 buy,uint256 fee);
   event GetMoney(address owner,uint256 money,uint256 time);



    //deposits Buy (sell eth buy erc20)
    function depositsOrder(uint64 orderid,
                      address useraddress,
                      address receiver,
                      bool tradetype,
                      uint256 sellamount,
                      uint256 price,
                      address buyaddress,
                      uint256 createTime
                      ) payable public returns(bool) {
       require(useraddress==msg.sender);
       require(startContract==true);
       require(sellamount==msg.value);
       if(tradetype==false){
           require(buyaddress!=address(0));
           require(contractTypeAddress[buyaddress]==true) ;
       }


       orders memory user = orders(
       orderid,
       useraddress,
       receiver,
       tradetype,
       sellamount,
       address(0),
       price,
       0,
       buyaddress,
       0,
       true,
       false,
       createTime,
       createTime);

       //addressToOrderid[msg.sender].push(orderid);
       orderidToUsers[orderid]=user;
       //

       emit Deposits(orderid,tradetype,sellamount,price,0,buyaddress,createTime);
       return true;
    }

     //deposits sell
    function depositsSell(uint64 orderid,
                      address useraddress,
                      address receiver,
                      bool tradetype,
                      uint256 sellamount,
                      address selladdress,
                      uint256 price,
                      uint256 createTime
                      )checkowner(msg.sender) payable public returns(bool) {

         require(startContract==true);
         require(tradetype==true);
         require(selladdress!=address(0));
         require(contractTypeAddress[selladdress]==true);
         require(ERC20(selladdress).allowance(useraddress,address(this))>=sellamount);

         ERC20(selladdress).transferFrom(useraddress,address(this),sellamount);

       orders memory user = orders(
       orderid,
       useraddress,
       receiver,
       tradetype,
       sellamount,
       selladdress,
       price,
       0,
       address(0),
       0,
       true,
       false,
       createTime,
       createTime);

       //addressToOrderid[msg.sender].push(orderid);
       orderidToUsers[orderid]=user;
       //

       emit DepositsSell(orderid,receiver,tradetype,sellamount,selladdress,price,0,createTime);
       return true;
    }



    //withdrawal
    function withdrawal(uint64 txid)payable public returns(bool){
       require(startContract==true);
       require(orderidToUsers[txid].user_address==msg.sender);
       uint256 money=orderidToUsers[txid].sell_amount;
       require(orderidToUsers[txid].withdrawalable==true && orderidToUsers[txid].ex_success!=true &&money>0);
       require(!contractMutex);
        if(orderidToUsers[txid].sell_address!=address(0)){
            contractMutex=true;
            ERC20(orderidToUsers[txid].sell_address).transfer(msg.sender,money);
            contractMutex=false;

            if(orderidToUsers[txid].buy_amount==0){
                delete(orderidToUsers[txid]);
            }else{
                orderidToUsers[txid].sell_amount=0;
                orderidToUsers[txid].withdrawalable=false;
                orderidToUsers[txid].ex_success=true;
                orderidToUsers[txid].update_at=now;
            }

            emit Withdrawl(txid);
            return true;
        }else{
           contractMutex=true;
           (msg.sender).transfer(money);
           contractMutex=false;

            if(orderidToUsers[txid].buy_amount==0){
                delete(orderidToUsers[txid]);
            }else{
                orderidToUsers[txid].sell_amount=0;
                orderidToUsers[txid].withdrawalable=false;
                orderidToUsers[txid].ex_success=true;
                orderidToUsers[txid].update_at=now;
            }

            emit Withdrawl(txid);
            return true;
        }

    }

    //update status
    function upstate(uint64 txid)public checkowner(msg.sender) {
        require(startContract==true);
        bool status =orderidToUsers[txid].withdrawalable;
        require(orderidToUsers[txid].ex_success!=true);
        if(status ==true){
            orderidToUsers[txid].withdrawalable=false;
        }else{
            orderidToUsers[txid].withdrawalable=true;
        }
    }

    //submit chain
    function submit(uint64 txid,address to,uint256 sell,uint256 buy,uint256 fee)
    checkowner(msg.sender) payable public returns(bool) {
         require(startContract==true);
         require(orderidToUsers[txid].sell_amount>=sell.add(fee)&&sell>0&&fee>0&&buy>0);


         uint256 sellAmount=orderidToUsers[txid].sell_amount;
              if(orderidToUsers[txid].sell_address!=address(0)){
              ERC20(orderidToUsers[txid].sell_address).transfer(to,sell);
              ERC20(orderidToUsers[txid].sell_address).transfer(contractOwner,fee);
              }else{
               (to).transfer(sell);
              }

              orderidToUsers[txid].sell_amount= sellAmount.sub(sell.add(fee));
              orderidToUsers[txid].fees=orderidToUsers[txid].fees.add(fee);
              orderidToUsers[txid].buy_amount=orderidToUsers[txid].buy_amount.add(buy);
              orderidToUsers[txid].update_at=now;

              if(sellAmount==sell.add(fee)){
                orderidToUsers[txid].withdrawalable=false;
                orderidToUsers[txid].ex_success=true;
              }else{
                orderidToUsers[txid].withdrawalable=true;
              }

              emit Submit(txid,to,sell,buy,fee);
              return true;

    }



    function getContractBalance()checkowner(msg.sender) public view returns(uint256) {
       // contrace_balance=address(this).balance;
        return address(this).balance;
    }

    //func of get money
    function getMoney(uint256 money) checkowner(msg.sender) public{
      require(money>0&&address(this).balance>=money);
      get_money_eth(money);
      emit GetMoney(msg.sender,money,now);
    }

    //func pause contract

    function pauseContract()checkowner(msg.sender) public{
       require(startContract==true);
       startContract=false;
    }
    //func restart contract
    function restartContract()checkowner(msg.sender) public{
       require(startContract==false);
       startContract=true;
    }

    //func of destory contract

    function destoryContract()checkowner(msg.sender) public{
       selfdestruct(contractOwner);
    }

    function get_money_eth(uint256 money)  internal  {
       contractOwner.transfer(money);
    }

    function changeContractAddress(address addr,uint64 tokenid,string types,bool chnageType)checkowner(msg.sender)public{
       if(chnageType==true){
           contractTypeAddress[addr]=true;
           tokens memory token=tokens(tokenid,types,addr);
           tokenTable[tokenid]=token;
       }else{
           delete(contractTypeAddress[addr]);
           delete(tokenTable[tokenid]);
       }
    }

   function getTokenInfor(uint64 tokenid)checkowner(msg.sender) view public returns(uint64,string,address) {
       uint64 a =tokenTable[tokenid].token_id;
       string storage b =tokenTable[tokenid].token_name;
       address c =tokenTable[tokenid].token_address;
       return (a,b,c);
  }


   //submit safe check
    modifier checkowner(address addr){
       require(contractOwner==addr);
       _;
    }



}