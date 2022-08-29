package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

type BlockStruct struct {
	Number uint64
	Time uint64
	Difficulty uint64
	Hash string
	Transactions types.Transactions
	Len int
}
// 传入block号 查询区块信息
func GetBlockInfoByNumber(block int64) BlockStruct {
	CheckClient("")
	var blockNumber *big.Int
	if block == 0 {
		blockNumber = nil
	}else {
		blockNumber = big.NewInt(block)
	}
	blockInfo, err := Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	number := blockInfo.Number().Uint64()     // 5671744
	// 区块时间戳
	Time := blockInfo.Time()       // 1527211625
	// 区块摘要
	Difficulty := blockInfo.Difficulty().Uint64() // 3217000136609065
	// 区块难度
	Hash := blockInfo.Hash().Hex()          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// 交易列表
	len := len(blockInfo.Transactions())   // 144
	return BlockStruct{number,Time,Difficulty,Hash,blockInfo.Transactions(),len}
}

func GetTransactionsByNumber(block int64)  {
	if ok := CheckClient(""); ok{
		fmt.Println("已经初始化客户端")
	}
	BlockStructs := GetBlockInfoByNumber(block)
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range BlockStructs.Transactions {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID),nil); err == nil {
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
		receipt, err := Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status) // 1
	}
}