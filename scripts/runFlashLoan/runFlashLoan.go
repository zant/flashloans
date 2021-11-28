package main

import (
  "fmt"
  "math/big"

  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  erc20 "github.com/zant/flash-loan/abigen/ERC20"
  flashLoan "github.com/zant/flash-loan/abigen/FlashLoan"
  auth "github.com/zant/flash-loan/utils"
)

const privateKey = "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"
const contractAddress = "0x2e144af3bde9b518c7c65fbe170c07c888f1ff1a"
const wethAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"

func main() {
  client, err := ethclient.Dial("http://localhost:8545")
  if err != nil {
    fmt.Println("Error: ", err)
  }

  auth, err := auth.GenAuth(privateKey, client, int64(0))
  if err != nil {
    fmt.Println("Error: ", err)
  }

  weth, err := erc20.NewErc20(common.HexToAddress(wethAddress), client)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  weth.Transfer(auth, common.HexToAddress(wethAddress), big.NewInt(int64(1000000000000000000)))

  balance, err := weth.BalanceOf(nil, common.HexToAddress(wethAddress))
  if err != nil {
    fmt.Println("Error: ", err)
  }

  fmt.Println(balance)

  instance, err := flashLoan.NewFlashLoan(common.HexToAddress(contractAddress), client)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  tx, err := instance.FlashLoan(auth)
  fmt.Println(tx)
  if err != nil {
    fmt.Println("Error: ", err)
  }
}
