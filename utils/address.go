package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	ERC20 "web3/utils/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 获取token余额
func GetBalance(address string, cli *ethclient.Client, token string) string {
	if cli == nil {
		cli = Client
	}
	// 获取原生token
	if token == "" {
		//"0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
		account := common.HexToAddress(address)
		balance, err := cli.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(balance)
		return balance.String()
		//fbalance := new(big.Float)
		//fbalance.SetString(balance.String())
		//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
		//fmt.Println(ethValue) // 25.729324269165216041
		//return ethValue.String()
		//pendingBalance, _ := Client.PendingBalanceAt(context.Background(), account)
		//fmt.Println(pendingBalance) // 25729324269165216042
	} else {
		fmt.Println("进行余额查询")
		tokenAddress := common.HexToAddress(token)
		instance, err := ERC20.NewUtils(tokenAddress, Client)
		if err != nil {
			log.Fatal(err)
		}
		addr := common.HexToAddress(address)

		balance, err := instance.BalanceOf(&bind.CallOpts{}, addr)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(balance.String())
		//fbalance := new(big.Float)
		//fbalance.SetString(balance.String())
		//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(6)))
		//fmt.Println(ethValue.String())
		return balance.String()
	}
}

// 生成一个新钱包 返回密钥和地址
func GetNewWallet() (string, string) {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
	priKey := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E
	// 以下是使用 golang.org/x/crypto/sha3 的 Keccak256函数手动完成的方法。
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyBytes[1:])
	//fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e

	return priKey, address
}

// 判断是否是合约地址 合约地址返回ture
func IsContract(addr string) bool {
	// 判断是不是有效地址
	//re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	//fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	//fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// a random user account address
	address := common.HexToAddress(addr)
	bytecode, err := Client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	return isContract
}
