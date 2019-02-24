package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// SendMessages houses logic for passing, saving and sending messages
func (client *Client) SendMessages(destinationNumbers []string, messageContent string, messageChannel chan Message) {
	fmt.Println("HEREE")
	databaseChannel := make(chan Message)

	

	for _, number := range destinationNumbers {
		go client.ExecuteRequest("POST", number, messageContent, messageChannel)
	}

	// How do you know the message will come in time ? We don't need to wait for all messages to come back
	for range destinationNumbers {
		message := <-messageChannel // Blocking operation waiting for a sender to send message through the channel

		// Lock the resource so that there isnt concurrent writes ... only viable so you dont have to block execution to save to the database
		go PostMessage(&message, databaseChannel)
	}

	// Is there a way to make this faster for now lets make the option of persistence optional
	func(destinationNumbers []string) {
		for range destinationNumbers {
			<-databaseChannel
		}
	}(destinationNumbers)
	
	// What do they want outputted to them? Right now we are only printing out the successful messages
}

// Acts as our main driver executes functionality with added logic
func main() {
	// So the first thing we need to is setup our client
	err := godotenv.Load() // First load environment variables file
	if err != nil {
		log.Fatal(err)
	}

	ConfigureDatabase()
	destinationNumbers, messageChannel := []string{"7183009363"}, make(chan Message)

	// Pass in credentials
	accountSID, authToken := os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN")
	sourceNumber := os.Getenv("SOURCE_NUMBER")

	// Your choice of client to execute the request used ... default is the http.DefaultClient
	clientManager := NewClient(nil, sourceNumber, authToken, accountSID)

	clientManager.SendMessages(destinationNumbers, "Beerus?", messageChannel)
}
