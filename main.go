package main

import (
	"web3/config"
	"web3/handler"
	"web3/utils"
	"web3/utils/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	utils.BscClient = utils.InitClient("Bsc")
	//utils.EehClient = utils.InitClient("Eth")
	//utils.MaticClient = utils.InitClient("Matic")
	//utils.AvaxClient = utils.InitClient("Avax")
}

func main() {
	//utils.GetBalance("0x61dd481a114a2e761c554b641742c973867899d3",nil, "0x1C749D5f5630Cf365673BF6C0B6B0570C48Da112")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger(), gin.Recovery())
	web3 := r.Group("/web3")
	{
		web3.GET("/getNewWallet/:sum", handler.GetNewWallet)
		web3.GET("/GetBalance/:address", handler.GetBalance)
	} // 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
