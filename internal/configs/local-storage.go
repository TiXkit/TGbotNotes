package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var PathFileLocalStorage string

func StartLocalStorage() error {
	err := godotenv.Load()
	if err != nil {
		return err // ошибка - не удалось загрузить файл конфигурации
	}

	notFoundVarsErr := errors.New("параметр(LOCAL_STORAGE) для подключения к локальному хранилищу не найден в конфигурационном файле \".env\"")
	localStorage, okLcS := os.LookupEnv("LOCAL_STORAGE")
	if !okLcS {
		return notFoundVarsErr
	}

	PathFileLocalStorage = localStorage
	return nil
}
