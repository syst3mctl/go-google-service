package mocks

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

// Mock oauth2.Config struct for testing
type MockConfig struct {
	*oauth2.Config
	ClientFunc func(ctx context.Context, t *oauth2.Token) *http.Client
}

// Client is a simple mock function
//
// Parameters:
//
//	-First parameters is a context.Context
//	-Second parameter is a pointer of oauth2.Token
//
// Returns:
//
//	-Pointer of http.Client
func (m *MockConfig) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return m.ClientFunc(ctx, t)
}
