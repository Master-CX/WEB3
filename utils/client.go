package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"web3/config"
)

var Client *ethclient.Client
var BscClient *ethclient.Client
var EehClient *ethclient.Client
var MaticClient *ethclient.Client
var AvaxClient *ethclient.Client

// 初始化客户端
func InitClient(chain string) *ethclient.Client {
	for i := 0; i < 5; i++ {
		switch chain {
		case "Bsc":
			BSCClient, err := ethclient.Dial(config.Conf.BscConfig.Url)
			if err != nil {
				continue
			}
			Client = BSCClient
		case "Eth":
			ETHClient, err := ethclient.Dial(config.Conf.EthConfig.Url)
			if err != nil {
				continue
			}
			Client = ETHClient
		case "Avax":
			AvaxClient, err := ethclient.Dial(config.Conf.AvaxConfig.Url)
			if err != nil {
				continue
			}
			Client = AvaxClient
		case "Matic":
			MaticClient, err := ethclient.Dial(config.Conf.MaticConfig.Url)
			if err != nil {
				continue
			}
			Client = MaticClient
		}
		return Client
	}
	return Client
}
func CheckClient(chain string) bool {
	if Client != nil {
		return true
	}
	if chain == "" {
		client := InitClient("Bsc")
		return client != nil
	}
	client := InitClient(chain)
	return client != nil
}
