package main

import (
	"fmt"

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

func ConfigureDatabase() {
	db.Debug().AutoMigrate(&Message{}) // Migrate the Message schema to our message database
}

func PostMessage(message *Message, messageChannel chan Message) Message {
	db.Debug().Create(message)
	foundMessage := db.Debug().First(message)
	fmt.Printf("Find after creating -> ", foundMessage.Value)
	messageChannel <- *message
	return *message
}

func FindMessage() *gorm.DB {
	message := db.Debug().Where("body = ?", "Make School")
	return message
}
