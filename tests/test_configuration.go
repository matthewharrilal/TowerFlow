package towerflow

import (
	"net/http"
	"strings"
	"testing"
	""
)
// Not robust enough
func TestNewClient(t *testing.T) {
	testClient := 

	if testClient.AuthToken == "" {
		t.Fatal("Auth token can not be empty nor nil!")
	}

	if testClient.AccountSID == "" {
		t.Fatal("Account Identifier can not be empty or nil")
	}

	if testClient.BaseURL == "" {
		t.Fatal("Need to provide url to receive messages")
	} else if !strings.Contains(testClient.BaseURL, testClient.AccountSID) {
		t.Fatal("Base URL must contain account identifier")
	}
}
