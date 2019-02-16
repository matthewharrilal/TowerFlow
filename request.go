package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// func ConstructRequest(destinationNumber string) (http.Client, *http.Request) {
// 	accountSID, authToken, urlString := AccountConfiguration()
// 	messageDataReader := ConstructMessage(destinationNumber)

// 	client := http.Client{} // In charge of executing the request

// 	// Formulate POST request with the given url string, and the encoded representation of the message body
// 	request, _ := http.NewRequest("POST", urlString, &messageDataReader) // Passing the message data reader by reference

// 	// Adds header field with the key name 'Authorization' and the two credentials we send as values to the Twillio API
// 	request.SetBasicAuth(accountSID, authToken)

// 	// Additional header fields to accept json media types which can be used for the response
// 	request.Header.Add("Accept", "application/json")

// 	// To indicate the media type that is being sent through the request
// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	return client, request
// }

func (client *Client) NewRequest(httpMethod string, messageDataBuffer *strings.Reader) *http.Request {
	// have to verify that the http method that the user passes in is valid
	request, err := http.NewRequest(httpMethod, client.BaseURL, messageDataBuffer)

	if err != nil {
		fmt.Println("FATAL ERROR CONSTRUCTING HTTP REQUEST")
		log.Fatal(err)
	}

	request.SetBasicAuth(client.AccountSID, client.AuthToken) // Authenticating user credentials

	// Should the header fields be static ... depending on what client is using this service it is going to have to be dynamic
	// We will make that dynamic in the next step


	// Additional header fields to accept json media types which can be used for the response
	request.Header.Add("Accept", "application/json")

	// To indicate the media type that is being sent through the request
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request
}
