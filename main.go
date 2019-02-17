package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func (client *Client) SendMessages(destinationNumbers []string, messageContent string, messageChannel chan Message) {
	fmt.Println("HEREE")
	for _, number := range destinationNumbers {
		go client.ExecuteRequest("POST", number, messageContent, messageChannel)
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

	// Pass in credentials
	accountSID, authToken := os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN")
	sourceNumber := os.Getenv("SOURCE_NUMBER")

	// Your choice of client to execute the request used ... default is the http.DefaultClient
	clientManager := NewClient(nil, sourceNumber, authToken, accountSID)

	// Now that we have the client manager we can need to construct our message

	// The process of creating the channel they should not have to see that process

	clientManager.SendMessages(destinationNumbers, "You suck", messageChannel)
}
