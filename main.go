package TowerFlow

import (
	"fmt"
)

// SendMessages houses logic for passing, saving and sending messages
func (client *Client) SendMessages(destinationNumbers []string, messageContent string, messageChannel chan Message) Message {
	var message Message

	// Spin a goroutine for each desired number
	for _, number := range destinationNumbers {
		go client.ExecuteRequest(number, messageContent, messageChannel)
	}

	// Receive message through channel
	for range destinationNumbers {
		message = <-messageChannel
		fmt.Printf("MESSAGE RECEIVED -> %v", message)

	}

	// Return last sent message
	return message
}
