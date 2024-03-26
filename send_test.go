package ctlgmail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendingTemplate(t *testing.T) {
	config := NewSMTPConfig()

	Cc := "ggdnicolas@gmail.com"
	Bcc := map[string][]string{
		"Bcc": {"ggdnicolas@gmail.com"},
	}

	sc := &SendingCredentials{
		From:           "ggdnicolas@gmail.com",
		To:             "ggdnicolas@gmail.com",
		SenderPassword: "qvpmuhxabbzpviwt",
		Subject:        "Gomail test subject",
		Template:       "test.tmpl",
		Cc:             &Cc,
		Bcc:            &Bcc,
	}

	data := map[string]any{
		"userEmail": sc.To,
	}

	err := SendEmail(config, sc, data, true)
	assert.NoError(t, err)
}
