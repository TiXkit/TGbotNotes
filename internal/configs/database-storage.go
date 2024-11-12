package configs

import (
	"database/sql"
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var DBClient *sql.DB

func StartDB() error {
	err := godotenv.Load()
	if err != nil {
		return err // ошибка - не удалось загрузить файл конфигурации
	}

	notFoundVarsErr := errors.New("параметры(DATABASE,DATABASE_PARAM) для подключения к базе данных не найдены в конфигурационном файле \".env\"")
	dataBase, okDB := os.LookupEnv("DATABASE")
	dataBaseParam, okDBP := os.LookupEnv("DATABASE_PARAM")
	if !okDB || !okDBP {
		return notFoundVarsErr
	}

	db, err := sql.Open(dataBase, dataBaseParam)
	if err != nil {
		return err
	}

	DBClient = db
	return nil
}
