package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// NewRequest in charge of creating request with given credentials and headers
func (client *Client) NewRequest(messageDataBuffer *strings.Reader) (*http.Request, error) {
	request, err := http.NewRequest("POST", client.BaseURL, messageDataBuffer)

	if err != nil {
		errStr := fmt.Sprintf("Error constructing the HTTP network request ... here is the error %v", err)
		return &http.Request{}, &errorString{errStr}
	}

	// fmt.Printf("ACCOUNT SID ", client.AccountSID , "Auth Token ", client.AuthToken)
	request.SetBasicAuth(client.AccountSID, client.AuthToken) // Authenticating user credentials

	// Additional header fields to accept json media types which can be used for the response
	request.Header.Add("Accept", "application/json")

	// To indicate the media type that is being sent through the request
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

// ExecuteRequest in charge of executing request and marshalling data inside our message object
func (client *Client) ExecuteRequest(destinationNumber string, messageContent string, messageChannel chan Message) (Message, error) {
	// Returns you a message Object back

	var message Message

	messageDataBuffer := client.NewMessage(messageContent, destinationNumber)

	request, err := client.NewRequest(messageDataBuffer)
	if err != nil {
		errStr := fmt.Sprintf("Error concerning HTTP credentials ... here is the error %v", err)
		return Message{}, &errorString{errStr}
	}

	response, err := client.RequestExecutor.Do(request)

	if err != nil {
		errStr := fmt.Sprintf("Error executing the HTTP request ... here is the error %v", err)
		return Message{}, &errorString{errStr}
	}

	if response.StatusCode >= 300 {
		errStr := fmt.Sprintf("Status Code : %v", response.StatusCode)
		return Message{}, &errorString{errStr}
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&message)

	if err != nil {
		errStr := fmt.Sprintf("Error decoding data into Message Object ... here is the data %v", err)
		return Message{}, &errorString{errStr}
	}

	messageChannel <- message
	return message, nil
}
