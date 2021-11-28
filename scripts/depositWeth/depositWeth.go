package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/zant/flashloan/abigen/ERC20"
	auth "github.com/zant/flashloan/utils"
)

const privateKey = "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"
const wethAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
const amount = int64(9000000000000000000)

func main() {
  client, err := ethclient.Dial("http://localhost:8545")
  if err != nil {
    fmt.Println("Error: ", err)
  }

  auth.Transfer(privateKey, client, wethAddress, amount)

  weth, err := erc20.NewErc20(common.HexToAddress(wethAddress), client)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  balance, err := weth.BalanceOf(nil, common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53"))
  if err != nil {
    fmt.Println("Error: ", err)
  }

  fmt.Println(balance)
}
