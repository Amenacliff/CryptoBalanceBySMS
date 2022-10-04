package main

import (
	"CryptoBalanceBySMS/constants"
	"CryptoBalanceBySMS/utils/envUtils"
)

func main() {
	envUtil := envUtils.InitEnvUtil()
	rpcLink, _ := envUtil.GetEnv(constants.INFURA_ETH_LINK)

}
