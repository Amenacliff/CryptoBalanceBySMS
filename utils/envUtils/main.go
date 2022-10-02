package envUtils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type IEnvUtil interface {
	GetEnv(key string) (string, error)
}

type EnvUtil struct {
}

func InitEnvUtil() *EnvUtil {
	newEvnUtil := EnvUtil{}
	return &newEvnUtil
}

func (envUtil *EnvUtil) GetEnv(key string) (string, error) {

	value, isValueIn := os.LookupEnv(key)

	if isValueIn {
		return value, nil
	}

	errLoadEvn := godotenv.Load()

	if errLoadEvn != nil {
		log.Println("Unable to Load ENV : ", errLoadEvn.Error())
		return "", errLoadEvn
	}

	envData, errReadEvn := godotenv.Read()

	if errReadEvn != nil {
		log.Println("Unable to Read ENV : ", errLoadEvn.Error())
		return "", errLoadEvn
	}

	return envData[key], nil
}
