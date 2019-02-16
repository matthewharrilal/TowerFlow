package main

type Message struct {
	DateCreated string `json:"date_created"`

	MessageDirection string `json:"direction"`

	AccountIdentifier string `json:"account_sid"`

	MessageIdentifier string `json:"sid"`

	MessageBody string `json:"body"`

	NumberOfSegments int `json:"num_segments"`
}
