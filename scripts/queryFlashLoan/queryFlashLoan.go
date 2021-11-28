package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	flashLoan "github.com/zant/flashloan/abigen/FlashLoan"
)

const contractAddress = "0xceb05c38a87c5c4df3dbb5e1934d144dac98e67f"

func main() {
  client, err := ethclient.Dial("http://localhost:8545")
  if err != nil {
    fmt.Println("Error: ", err)
  }

  instance, err := flashLoan.NewFlashLoan(common.HexToAddress(contractAddress), client)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  greet, err := instance.Check(nil)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  fmt.Println(greet)
}
