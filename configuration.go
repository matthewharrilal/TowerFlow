package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

type Client struct {
	// What is the neccesary information needed to create a client?
	RequestExecutor http.Client // Each individual has the ability to execute their request with the added configurations

	AuthToken string 

	AccountSID string 

	Url string 
}

// Constructs a new client with the given credentials

func NewClient(requestExecutor http.Client, authToken string, accountSID string) Client {
	baseURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

	client := Client{requestExecutor, authToken, accountSID, baseURL} // Creating a client with the configurations added 

	// Now that we have the client with the configurations added 

	return client
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
