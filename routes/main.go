package routes

import (
	"CryptoBalanceBySMS/controllers"
	"CryptoBalanceBySMS/utils/envUtils"
	"net/http"
)

func SetUpRoutes(mux *http.ServeMux, envUtil envUtils.IEnvUtil) {
	mainController := controllers.InitMainController(envUtil)
	mux.HandleFunc("/getBalance", mainController.GetBalance)
}
