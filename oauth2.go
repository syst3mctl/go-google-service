package ctlgmail

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

// GetClient func retrieves a token, saves the token, then returns the generated client.
//
// Parameters:
//
//	-Pointer of oauth2.Config
//	-File name of tokenFile should be string
//
// Returns:
//
//	-Pointer of http.Client
func GetClient(config *oauth2.Config, tokenFile string) *http.Client {
	// retrieve token from given file
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		// get token from web
		tok = getTokenFromWeb(config)
		// save token in file
		saveToken(tokenFile, tok)
	}

	// return generated client
	return config.Client(context.Background(), tok)
}

// Requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	// while oauth2.NoContext is depreciated, replaced with context.TODO()
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	//  read and write the file or directory
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	defer f.Close()

	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to encode token, got error: %v", err)
	}
}
