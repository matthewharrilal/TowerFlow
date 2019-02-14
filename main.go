package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func AccountConfiguration() (string, string, string) {
	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

	return accountSID, authToken, url
}
