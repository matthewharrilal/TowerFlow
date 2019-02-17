package main

import (
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, err = gorm.Open("sqlite3", "message.db")

type Message struct {
	gorm.Model
	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"`

	AccountIdentifier string `json:"account_sid"`

	MessageIdentifier string `json:"sid"`

	Body string `json:"body"`

	NumberOfSegments string `json:"num_segments"`
}

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

func ConfigureDatabase() {
	db.Debug().AutoMigrate(&Message{}) // Migrate the Message schema to our message database
}

func PostMessage(message *Message, databaseChannel chan Message) Message {
	db.Debug().Create(&message)
	databaseChannel <- *message
	return *message
}

// OBJECTIVE: Be able to query for message

// func FindMessage() *gorm.DB {
// 	var messageObj []Message
// 	message := db.Debug().Where("body=?", "Sativa").First(&messageObj)
// 	fmt.Printf("PLUCKED VALUE ", message)
// 	return message
// }
