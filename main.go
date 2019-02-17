package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func (client *Client) SendMessages(httpMethod string, destinationNumbers []string, messageContent string, messageChannel chan Message) {
	fmt.Println("HEREE")
	for _, number := range destinationNumbers {
		go client.ExecuteRequest(httpMethod, number, messageContent, messageChannel)
	}

	for range destinationNumbers {
		message := <-messageChannel
		fmt.Printf("Successful Message -> ", message)
	}

	// What do they want outputted to them? Right now we are only printing out the successful messages
}

func main() {
	// So the first thing we need to is setup our client
	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

	destinationNumbers, messageChannel := []string{"7183009363", "6304077258"}, make(chan Message)

	accountSID, authToken := os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN")
	sourceNumber := os.Getenv("SOURCE_NUMBER")

	// Your choice of client to execute the request used ... default is the http.DefaultClient
	clientManager := NewClient(nil, sourceNumber, authToken, accountSID)

	// Now that we have the client manager we can need to construct our message

	clientManager.SendMessages("POST", destinationNumbers, "You suck", messageChannel)
}
