package main

import (
	"CryptoBalanceBySMS/constants"
	"CryptoBalanceBySMS/routes"
	"CryptoBalanceBySMS/utils/envUtils"
	"log"
	"net/http"
)

func main() {
	envUtil := envUtils.InitEnvUtil()

	port, errGetPort := envUtil.GetEnv(constants.PORT)

	if errGetPort != nil {
		log.Println("An Error Occurred when getting Port", errGetPort.Error())
	}
	mux := http.NewServeMux()

	routes.SetUpRoutes(mux, envUtil)

	http.ListenAndServe(port, mux)
}
