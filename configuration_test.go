package main

import (
	"testing"
	"net/http"
	"strings"
)

// Not robust enough
func TestNewClient(t *testing.T) {
	testClient := Client{http.Client{}, "", "1", "1", "1"}

	if testClient.AuthToken == "" {
		t.Fatal("Auth token can not be empty nor nil!")
	}

	if testClient.AccountSID == ""{
		t.Fatal("Account Identifier can not be empty or nil")
	}

	if testClient.BaseURL == ""{
		t.Fatal("Need to provide url to receive messages")
	} else if !strings.Contains(testClient.BaseURL, testClient.AccountSID) {
		t.Fatal("Base URL must contain account identifier")
	}
}