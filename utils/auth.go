package auth

import (
  "context"
  "crypto/ecdsa"
  "fmt"
  "log"
  "math/big"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/ethclient"
)

func handleError(err error) {
  if err != nil {
    fmt.Println("Error has occurred", err)
  }
}

func GenAuth(hexKey string, client *ethclient.Client, value int64) (*bind.TransactOpts, error) {
  privateKey, err := crypto.HexToECDSA(hexKey)
  handleError(err)

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
    log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
  }
  fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
  nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
  handleError(err)

  gasPrice, err := client.SuggestGasPrice(context.Background())
  handleError(err)

  chainID, err := client.NetworkID(context.Background())
  fmt.Println(chainID)
  if err != nil {
    log.Fatal(err)
  }

  auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
  auth.Nonce = big.NewInt(int64(nonce))
  auth.Value = big.NewInt(value)
  auth.GasLimit = uint64(3000000)
  auth.GasPrice = gasPrice

  return auth, err
}

func Transfer(hexKey string, client *ethclient.Client, toAddress string, value int64) error {
  privateKey, err := crypto.HexToECDSA(hexKey)
  handleError(err)

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
    log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
  }
  fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
  nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
  handleError(err)

  // gasPrice, err := client.SuggestGasPrice(context.Background())
  // handleError(err)

  gasLimit := uint64(3000000)
  var data []byte

  tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), big.NewInt(value), gasLimit, big.NewInt(int64(2000000000)), data)

  chainID, err := client.NetworkID(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
  if err != nil {
    log.Fatal(err)
  }

  err = client.SendTransaction(context.Background(), signedTx)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("tx sent: ", signedTx.Hash().Hex())
  return err
}
