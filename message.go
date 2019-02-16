package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, err = gorm.Open("sqlite3", "message.db")

type Message struct {
	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"`

	AccountIdentifier string `json:"account_sid"`

	MessageIdentifier string `json:"sid"`

	MessageBody string `json:"body"`

	NumberOfSegments string `json:"num_segments"`
}

func ConfigureDatabase() {
	db.Debug().AutoMigrate(&Message{}) // Migrate the Message schema to our message database
}

func PostMessage(message *Message) (Message) {
	db.Debug().Create(message)
	fmt.Printf("Created Message -> ", message)

	return *message
}
