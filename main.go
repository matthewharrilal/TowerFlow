package main

import (
	"fmt"
)

func (client *Client) SendMessages(httpMethod string, destinationNumbers [80]string, messageChannel chan Message) {
	for _, number := range destinationNumbers {
		go client.ExecuteRequest(httpMethod, number, messageChannel)
	}

	for range destinationNumbers {
		message := <-messageChannel
		fmt.Sprintf("Successful Message -> ", message)
	}

	// What do they want outputted to them? Right now we are only printing out the successful messages
}

func main() {
	// So the first thing we need to is setup our client 

	

	authToken, accountSid := 

	// Your choice of client to execute the request used ... default is the http.DefaultClient
	clientManager := NewClient(nil, sourceNumber, authToken, accountSid)
}
