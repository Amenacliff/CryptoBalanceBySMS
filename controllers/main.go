package controllers

import (
	"CryptoBalanceBySMS/utils/envUtils"
	"net/http"
)

type MainController struct {
	envUtil envUtils.IEnvUtil
}

type IMainController interface {
	GetBalance(w http.ResponseWriter, r *http.Request) error
}

func InitMainController(envUtil envUtils.IEnvUtil) *MainController {
	return &MainController{envUtil: envUtil}
}

func (mainCont *MainController) GetBalance(w http.ResponseWriter, r *http.Request) {

}
