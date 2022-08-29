package bsc

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var RPC *rpc.Client

func GetBalance(address string) {
	//"0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	account := common.HexToAddress(address)
	balance, err := Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(balance)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	//pendingBalance, _ := Client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}