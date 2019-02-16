package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func ConstructMessage(destinationNumber string) strings.Reader {
	// Constructs message object with given source and destination
	messageData := url.Values{} // Used to store and encode following parameters to be sent over the network
	sourceNumber := "6468324582"
	messageStub := "Hey, this is Matthew and this is my project for school. It allows me to be able to send text messages to a group of people simultaneously"

	// Setting source number and destination number
	messageData.Set("From", sourceNumber)
	messageData.Set("To", destinationNumber)

	messageData.Set("Body", messageStub)

	// Message Data Reader acts as a buffer to transport data between processes
	messageDataReader := *strings.NewReader(messageData.Encode())
	return messageDataReader
}

func ConstructRequest(destinationNumber string) (http.Client, *http.Request) {
	accountSID, authToken, urlString := AccountConfiguration()
	messageDataReader := ConstructMessage(destinationNumber)

	client := http.Client{} // In charge of executing the request

	// Formulate POST request with the given url string, and the encoded representation of the message body
	request, _ := http.NewRequest("POST", urlString, &messageDataReader) // Passing the message data reader by reference

	// Adds header field with the key name 'Authorization' and the two credentials we send as values to the Twillio API
	request.SetBasicAuth(accountSID, authToken)

	// Additional header fields to accept json media types which can be used for the response
	request.Header.Add("Accept", "application/json")

	// To indicate the media type that is being sent through the request
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client, request
}

// What pairing does this function return
func ExecuteRequest(destinationNumber string, channel chan Message) (Message, error) {
	// Access to the request executor and the request itself with configurations already implemented
	client, request := ConstructRequest(destinationNumber)
	var message Message
	
	var dataCopy Message
	response, err := client.Do(request) // Execute the request and store the response

	// If there was an error executing the request
	if err != nil {
		fmt.Println("Error executing the request")
		log.Fatal(err)
	}

	// Checking if response came back successful
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		// Data consisting of string keys and dynamic value types depending on the JSON coming back
		
		// Decode the response body
		decoder := json.NewDecoder(response.Body)

		err := decoder.Decode(&message) // Read the decoded data into our data map

		if err != nil {
			log.Fatal(err)
			return Message{}, err
		}
		dataCopy = message
	} else {
		fmt.Printf("Status Code not successful ", response.StatusCode)
	}
	channel <- dataCopy
	// defer waitGroup.Done()
	return dataCopy, nil
}

func main() {
	destinationNumbers := []string{"7183009363"}
	channel := make(chan Message)
	for _, number := range destinationNumbers {
		go ExecuteRequest(number, channel)
	}

	for range destinationNumbers {
		fmt.Println("VALUE FROM CHANNEL >>> ", <-channel)
	}
	fmt.Println("Done")
}
