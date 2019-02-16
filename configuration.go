package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	// What is the neccesary information needed to create a client?
	RequestExecutor http.Client // Each individual has the ability to execute their request with the added configurations

	SourceNumber string

	AuthToken string

	AccountSID string

	BaseURL string
}

// Constructs a new client with the given credentials

// func AccountConfiguration() (string, string, string) {
// 	// Loads our environement variables and configures url that we are going to be pinging

// 	err := godotenv.Load() // First load environment variables file
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	accountSID := os.Getenv("ACCOUNT_SID")
// 	authToken := os.Getenv("AUTH_TOKEN")

// 	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

// 	return accountSID, authToken, url
// }

func NewClient(requestExecutor *http.Client, sourceNumber string, authToken string, accountSID string) Client {
	// In charge of creating a client capable of executing requests with dynamic configurations already attached

	if requestExecutor == nil {
		requestExecutor = http.DefaultClient
	}

	baseURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

	client := Client{*requestExecutor, sourceNumber, authToken, accountSID, baseURL} // Creating a client with the configurations added

	// Now that we have the client with the configurations added

	return client
}
