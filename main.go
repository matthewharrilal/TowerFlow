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

func ConstructRequest() {
	accountSID, authToken, urlString := AccountConfiguration()
	messageDataReader := ConstructMessage()

	client := http.Client{} // In charge of executing the request
	
	// Formulate POST request with the given url string, and the encoded representation of the message body
	req, _ := http.NewRequest("POST", urlString, &messageDataReader) // Passing the message data reader by reference

	// Adds header field with the key name 'Authorization' and the two credentials we send as values to the Twillio API
	req.SetBasicAuth(accountSID, authToken)

	// Additional header fields to accept json media types which can be used for the response
	req.Header.Add("Accept", "application/json")

	// To indicate the media type that is being sent through the request
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}