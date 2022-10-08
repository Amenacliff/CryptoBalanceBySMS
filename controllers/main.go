package controllers

import (
	"CryptoBalanceBySMS/dtos"
	"CryptoBalanceBySMS/utils/envUtils"
	"encoding/json"
	"io"
	"log"
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
	log.Println("Got Request")
	requestBody := r.Body
	readRequestBody, errReadRequestBody := io.ReadAll(requestBody)

	if errReadRequestBody != nil {
		log.Println(errReadRequestBody.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errReadRequestBody.Error()))
		return
	}

	getBalanceRequest := new(dtos.GetBalanceReqDTO)

	errMarshalReq := json.Unmarshal(readRequestBody, getBalanceRequest)

	if errMarshalReq != nil {
		log.Println(errReadRequestBody.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMarshalReq.Error()))
		return
	}

	log.Println(getBalanceRequest)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Response Gotten"))

	return

}
