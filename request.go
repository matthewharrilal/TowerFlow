package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (client *Client) NewRequest(httpMethod string, messageDataBuffer *strings.Reader) *http.Request {
	request, err := http.NewRequest(httpMethod, client.BaseURL, messageDataBuffer)

	if err != nil {
		fmt.Println("FATAL ERROR CONSTRUCTING HTTP REQUEST")
		log.Fatal(err)
	}

	// fmt.Printf("ACCOUNT SID ", client.AccountSID , "Auth Token ", client.AuthToken)
	request.SetBasicAuth(client.AccountSID, client.AuthToken) // Authenticating user credentials

	// Should the header fields be static ... depending on what client is using this service it is going to have to be dynamic
	// We will make that dynamic in the next step

	// Additional header fields to accept json media types which can be used for the response
	request.Header.Add("Accept", "application/json")

	// To indicate the media type that is being sent through the request
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request
}

func (client *Client) ExecuteRequest(httpMethod string, destinationNumber string, messageContent string, messageChannel chan Message) Message {
	// Returns you a message Object back

	var message Message

	messageDataBuffer := client.NewMessage(messageContent, destinationNumber)

	request := client.NewRequest(httpMethod, messageDataBuffer)

	response, err := client.RequestExecutor.Do(request)

	if err != nil {
		fmt.Println("ERROR EXECUTING THE REQUEST")
		log.Fatal(err)
	}

	if response.StatusCode >= 300 {
		err := fmt.Sprintf("Status Code :%v", response.StatusCode)
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Succesful Message %v", message)
	messageChannel <- message
	return message
}
