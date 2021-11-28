package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	flashLoan "github.com/zant/flashloan/abigen/FlashLoan"
	auth "github.com/zant/flashloan/utils"
)

const privateKey = "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"

func main() {
  client, err := ethclient.Dial("http://localhost:8545")
  if err != nil {
    fmt.Println("Error: ", err)
  }

  auth, err := auth.GenAuth(privateKey, client, 0)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  address, tx, _, err := flashLoan.DeployFlashLoan(auth, client, common.HexToAddress("0xb53c1a33016b2dc2ff3653530bff1848a515c8c5"))
  // address, tx, _, err := flashLoan.DeployFlashLoan(auth, client, "hi")
  if err != nil {
    fmt.Println("Error: ", err)
  }

  fmt.Println(address.Hex())
  fmt.Println(tx.Hash().Hex())
}
