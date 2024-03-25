package ctlgmail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendingTemplate(t *testing.T) {
	config := NewSMTPConfig()

	sc := &SendingCredentials{
		From:           "ggdnicolas@gmail.com",
		To:             "ggdnicolas@gmail.com",
		SenderPassword: "qvpmuhxabbzpviwt",
		Subject:        "Gomail test subject",
		Template:       "test.tmpl",
	}

	data := map[string]any{
		"userEmail": sc.To,
	}

	err := SendEmail(config, sc, data, true)
	assert.NoError(t, err)
}
