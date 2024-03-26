package ctlgmail

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/systemctl/go-google-service/mocks"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

func TestGetClient(t *testing.T) {
	// create mock token
	mockToken := &oauth2.Token{AccessToken: "mock-token"}

	// mock the tokenFromFile function to return the mock token
	_ = func(string) (*oauth2.Token, error) {
		return mockToken, nil
	}

	// mock config with a ClientFunc, just return a dummy http.Client
	mockConfig := &mocks.MockConfig{
		ClientFunc: func(context.Context, *oauth2.Token) *http.Client {
			return &http.Client{}
		},
	}

	client := GetClient(mockConfig.Config, "your_credentials.json") // <- replace your credentials.json here
	assert.NotNil(t, client)
}
