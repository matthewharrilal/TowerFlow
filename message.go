package main

import (
	"fmt"
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

func (client *Client) NewMessage(messageStub string, destinationNumber string) *strings.Reader {
	// The goal of this function is to be able to construct a message object and return it

	messageData := url.Values{} // Map containing url query parameters

	// When creating a message object we need the source number the destination number and the message stub
	messageData.Set("From", client.SourceNumber)
	messageData.Set("To", destinationNumber)

	messageData.Set("Body", messageStub)

	messageDataBuffer := strings.NewReader(messageData.Encode())

	return messageDataBuffer // Return a buffer of data containing encapsulated configurations
}

func ConfigureDatabase() {
	db.Debug().AutoMigrate(&Message{}) // Migrate the Message schema to our message database
}

func PostMessage(message *Message) Message {
	db.Debug().Create(&message)
	test := db.Debug().Where("Body = ?", "Sativa")
	fmt.Printf("!!!!!!!!! ", test)
	// messageChannel <- *message
	return *message
}

func FindMessage() *gorm.DB {
	var messageObj []Message
	message := db.Debug().Where("body=?", "Sativa").First(&messageObj)
	fmt.Printf("PLUCKED VALUE ", message)
	return message
}
