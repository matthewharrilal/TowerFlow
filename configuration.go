package TowerFlow

import (
	"fmt"
	"net/http"
)

// Client containing relevant information to our client
type Client struct {
	RequestExecutor http.Client // Each individual has the ability to execute their request with the added configurations

	SourceNumber string

	AuthToken string

	AccountSID string

	BaseURL string
}

// NewClient in charge of returning a new client with users dynamic authentication credentials
func NewClient(requestExecutor *http.Client, sourceNumber string, authToken string, accountSID string) Client {
	// In charge of creating a client capable of executing requests with dynamic configurations already attached

	if requestExecutor == nil {
		requestExecutor = http.DefaultClient
	}

	baseURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%v/Messages.json", accountSID)
	client := Client{*requestExecutor, sourceNumber, authToken, accountSID, baseURL} // Creating a client with the configurations added

	return client
}
