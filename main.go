package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"net/url"
	"net/http"
	"github.com/joho/godotenv"
)

func AccountConfiguration() (string, string, string) {
	// Loads our environement variables and configures url that we are going to be pinging

	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	url := fmt.Printf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)

	return accountSID, authToken, url
}

func ConstructMessage() strings.Reader {
	// Constructs message object with given source and destination
	// _, _, url := AccountConfiguration()

	messageData := url.Values{} // Used to store and encode following parameters to be sent over the network
	destinationNumber, sourceNumber := "7183009363", "6468324582"
	messageStub := "You are receiving a test message"
	
	// Setting source number and destination number
	messageData.Set("From", sourceNumber)
	messageData.Set("To", destinationNumber)

	
	messageData.Set("Body", messageStub)

	messageDataReader := *strings.NewReader(messageData.Encode())
	fmt.Printf("Message Data Reader ", messageDataReader)
	return messageDataReader
}



