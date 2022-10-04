package evmService

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type EvmService struct {
	client *ethclient.Client
}

func InitEvmService(rpcProvider string) *EvmService {
	client, err := ethclient.Dial(rpcProvider)
	if err != nil {
		log.Println(err.Error(), "Unable to dial Client")
	}
	return &EvmService{client: client}
}

type IEvmService interface {
}
