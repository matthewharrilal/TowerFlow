package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// db, err := gorm.Open("sqlite3", "message.db")

type Message struct {
	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"`

	AccountIdentifier string `json:"account_sid"`

	MessageIdentifier string `json:"sid"`

	MessageBody string `json:"body"`

	NumberOfSegments string `json:"num_segments"`
}

// func PostMessage(message Message) (Message, error) {

// }
