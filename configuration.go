package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
	// What is the neccesary information needed to create a client?
	AuthToken string 

	AccountSID string 

	Url string 
}

func AccountConfiguration() (string, string, string) {
	// Loads our environement variables and configures url that we are going to be pinging

	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

	return accountSID, authToken, url
}
