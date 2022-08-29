package utils

import (
	"fmt"
	"testing"
	"web3/config"
)

func init() {

}
func TestGetBlockInfoByNumber(t *testing.T) {
	res := GetBlockInfoByNumber(0)
	fmt.Println(res)
}
func TestGetTransactionsByNumber(t *testing.T) {
	config.InitConfig()
	BscClient =InitClient("Bsc")
	GetTransactionsByNumber(0)
}