package api

import (
	"flag"
	"log"
)

func MustToken() string {
	token := flag.String(
		"token-DataSave-bot",
		"",
		"Please input the token to launch DataSaveBot")

	flag.Parse()

	if *token == "" {
		log.Fatal("the token is empty. Please enter a valid token")
	}
	return *token
}
