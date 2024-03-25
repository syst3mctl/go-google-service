package ctlgmail

import (
	"bytes"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"io"
	"log"
	"os"
)

// SendingCredentials defines the information needed for sending an email
type SendingCredentials struct {
	From           string              // The email address of the sender
	To             string              // The email address of the recipient
	SenderPassword string              // The email sender password
	Subject        string              // The subject line of the email
	Template       string              // The template content for the email body
	Cc             string              // The carbon copy
	Bcc            map[string][]string // The blind carbon copy
}

// SendEmail sends an email using the provided credentials and data
//
// Parameters:
//
//	-Pointer of SendingCredentials
//	-Data(any): Second parameter should be any type, map, string, struct etc...
//	-skipSecurity(bool): By default, gmail validates the server's TLS certificate, to avoid potential security issues due to invalid certificates, set 'false' during production.
//
// Returns:
//
//	-An error if there was a problem sending email
func SendEmail(config *SMTPConfig, sc *SendingCredentials, data any, skipSecurity bool) error {
	// Create a new gmail message object
	m := gomail.NewMessage()

	// Open the email template file
	tmplFile, err := os.Open(sc.Template)
	if err != nil {
		log.Printf("Error opening template file: %v", err)
		return err // Return error if opening template fails
	}
	defer tmplFile.Close() // Close the file deferentially

	// Read the template content into a byte slice
	tmplBytes, err := io.ReadAll(tmplFile)
	if err != nil {
		log.Printf("Error reading template content: %v", err)
		return err // Return error if reading template fails
	}

	// Parse the email template
	tmpl, err := template.New("email").Parse(string(tmplBytes))
	if err != nil {
		log.Printf("Error parsing email template: %v", err)
		return err // Return error if parsing template fails
	}

	// Create a buffer to hold the HTML email body
	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		log.Printf("Error executing email template: %v", err)
		return err // Return error if executing template fails
	}

	// Set the email sender address
	m.SetHeader("From", sc.From)

	// Set the email recipient(s) address(es)
	m.SetHeader("To", sc.To)

	// Set the email subject line
	m.SetHeader("Subject", sc.Subject)

	// Set carbon copy
	m.SetHeader("Cc", sc.Cc)

	// Set blind carbon copy
	m.SetHeaders(sc.Bcc)

	// Set the email body as HTML content
	m.SetBody("text/html", htmlBody.String())

	// Configure the SMTP dialer for Gmail server
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, sc.From, sc.SenderPassword) // Replace "your_password" with actual password

	// By default, gmail validates the server's TLS certificate.
	// This line is for development purposes only and should be set to 'false' in production
	// to avoid potential security issues due to invalid certificates.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: skipSecurity}

	// Send the email
	err = d.DialAndSend(m)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err // Return error if sending email fails
	}

	return nil
}
