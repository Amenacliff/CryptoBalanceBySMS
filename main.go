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
	log.Println("Server Running at", port)

	routes.SetUpRoutes(mux, envUtil)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Println(err.Error())
		return
	}
}
