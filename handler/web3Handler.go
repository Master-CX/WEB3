package handler

import (
	"strconv"

	"web3/utils"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	address := c.Param("address")
	balance := utils.GetBalance(address, nil, "")
	c.JSON(200, gin.H{
		"地址": address,
		"余额": balance,
	})
}
func GetNewWallet(c *gin.Context) {
	type AddressAndPriKey struct {
		Address string
		PriKey  string
	}
	sum := c.Param("sum")
	atoi, _ := strconv.Atoi(sum)
	res := make([]AddressAndPriKey, atoi)

	for i := 0; i < atoi; i++ {

		wallet, priKey := utils.GetNewWallet()
		//fmt.Println("密钥：", priKey, " 地址：", wallet)
		res[i] = AddressAndPriKey{
			Address: wallet,
			PriKey:  priKey,
		}
	}

	c.JSON(200, res)
}
