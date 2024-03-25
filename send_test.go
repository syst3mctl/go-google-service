package ctlgmail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendingTemplate(t *testing.T) {
	config := NewSMTPConfig()

	Cc := "example@gmail.com"
	Bcc := map[string][]string{
		"Bcc": {"example@gmail.com"},
	}

	sc := &SendingCredentials{
		From:           "yourgmail@gmail.com",
		To:             "yourgmail@gmail.com",
		SenderPassword: "your_password",
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
