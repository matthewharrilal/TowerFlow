package main

import (
	"net/url"
	"strings"
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, err = gorm.Open("sqlite3", "message.db")

// Message structure containing relevant information to our message object
type Message struct {

	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"`

	AccountIdentifier string `json:"account_sid"`

	MessageIdentifier string `json:"sid"`

	Body string `json:"body"`

	NumberOfSegments string `json:"num_segments"`
}

type DatabaseMessage struct {
	// Contains the full slice of destination numbers that the user sends in
	gorm.Model

	DateCreated string 

	MessageDirection string 

	AccountIdentifier string 

	MessageIdentifier string 

	Body string 

	NumberOfSegments string 

	DestinationNumbers pq.StringArray `gorm:"type:varchar(100)[]" `

}



// NewMessage formulates message object with source, destination and message contents
func (client *Client) NewMessage(messageContent string, destinationNumber string) *strings.Reader {
	// The goal of this function is to be able to construct a message object and return it

	messageData := url.Values{} // Map containing url query parameters

	// When creating a message object we need the source number the destination number and the message stub
	messageData.Set("From", client.SourceNumber)
	messageData.Set("To", destinationNumber)

	messageData.Set("Body", messageContent)

	messageDataBuffer := strings.NewReader(messageData.Encode())
	// fmt.Printf("Message data buffer ", messageDataBuffer)
	return messageDataBuffer // Return a buffer of data containing encapsulated configurations
}

// ConfigureDatabase in charge of migrating the schema
func ConfigureDatabase() {
	db.Debug().AutoMigrate(&DatabaseMessage{}) // Migrate the Message schema to our message database
}

// //PostMessage ... have to be able to successfully query for the message
// func PostMessage(message *Message, databaseChannel chan Message) Message {
// 	// defer db.Close
// 	db.Debug().Create(&message)
// 	databaseChannel <- *message
// 	return *message
// }

// PostMessage extended off the Message structure if user wants to save message
func PostMessage(message *Message, destinationNumbers []string) DatabaseMessage {
	// Use regualar message object and slice of destination numbers to formulate new database message object
	// Then return the database message object ... does not need to happen concurrently
	databaseMessage := &DatabaseMessage{DateCreated: message.DateCreated, MessageDirection: message.MessageDirection, AccountIdentifier: message.AccountIdentifier, MessageIdentifier: message.MessageIdentifier, Body: message.Body, NumberOfSegments: message.NumberOfSegments, DestinationNumbers: destinationNumbers}
	db.Debug().Create(&databaseMessage)
	return *databaseMessage
}

// OBJECTIVE: Be able to query for message

// func FindMessage() *gorm.DB {
// 	var messageObj []Message
// 	message := db.Debug().Where("body=?", "Sativa").First(&messageObj)
// 	fmt.Printf("PLUCKED VALUE ", message)
// 	return message
// }
