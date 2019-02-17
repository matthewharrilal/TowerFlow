package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// // What pairing does this function return
// func ExecuteRequest(destinationNumber string, channel chan Message) (Message, error) {
// 	// Access to the request executor and the request itself with configurations already implemented
// 	client, request := ConstructRequest(destinationNumber)
// 	var message Message

// 	dataCopy := &Message{}
// 	response, err := client.Do(request) // Execute the request and store the response

// 	// If there was an error executing the request
// 	if err != nil {
// 		fmt.Println("Error executing the request")
// 		log.Fatal(err)
// 	}

// 	// Checking if response came back successful
// 	if response.StatusCode >= 200 && response.StatusCode < 300 {
// 		// Data consisting of string keys and dynamic value types depending on the JSON coming back

// 		// Decode the response body
// 		decoder := json.NewDecoder(response.Body)

// 		err := decoder.Decode(&message) // Read the decoded data into our data map

// 		if err != nil {
// 			log.Fatal(err)
// 			return Message{}, err
// 		}
// 		// Pass in a channel and pass the message that was created through the channel and recieve it
// 		dataCopy = &message
// 	} else {
// 		fmt.Printf("Status Code not successful ", response.StatusCode)
// 	}

// 	channel <- *dataCopy
// 	return *dataCopy, nil
// }

func (client *Client) NewRequest(httpMethod string, messageDataBuffer *strings.Reader) *http.Request {
	// have to verify that the http method that the user passes in is valid
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

func (client *Client) ExecuteRequest(httpMethod string, destinationNumber string, messageContent string,  messageChannel chan Message) Message {
	// Returns you a message Object back

	fmt.Println("ENTERING")
	var message Message

	messageDataBuffer := client.NewMessage(messageContent, destinationNumber)

	// fmt.Printf("Message Data Bufffer ", messageDataBuffer)
	request := client.NewRequest(httpMethod, messageDataBuffer)
	// fmt.Printf("REQUEST ", request)

	response, err := client.RequestExecutor.Do(request)

	if err != nil {
		fmt.Println("ERROR EXECUTING THE REQUEST")
		log.Fatal(err)
	}

	// fmt.Printf("RESPONSE ", response)
	if response.StatusCode >= 300 {
		err := fmt.Sprintf("Status Code :", response.StatusCode)
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message -> ", message)
	messageChannel <- message
	return message
}
