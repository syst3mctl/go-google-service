# go-google-service

This go package is focused to sending emails, have a possibilities to working with `templ` or `html` templates with defined tags

### Installation
To use go-google-service in your Go project, you can install it via go get:
```bash
go get github.com/systemctl/go-google-service
```

## Custom Mail Sending Function

This package provides a convenient way to send emails using customizable templates and SMTP credentials. It leverages the gomail library for email functionality.

## Sending credentials
To send emails, you'll need to set up SendingCredentials which includes information like sender email, recipient email, password, subject, template path, and optional CC/BCC recipients.

## Sending an Email
Here's an example of how to send an email using the SendEmail function:

```go
    package main

    // Import required packages
    import (
        "log"
        
        "github.com/gomail/gomail"
        "your-username/email-sender" // Replace with your package import path
    )
    
    func main() {
            // Configure SMTP settings (optional, defaults to Google SMTP)
            // To change SMTP configuration within your config, just call function WithSMTPConfig() and pass parameters
            config := ctlgmail.NewSMTPConfig()
            
            // Set up sending credentials (replace placeholders with actual values)
            credentials := &mail.SendingCredentials{
                From:           "your_email@example.com",
                To:             "recipient@example.com",
                SenderPassword: "your_password", // **Important:**  Keep this secure!
                Subject:        "Email Subject",
                Template:       "path/to/your/template.html",
                Cc:             nil, // Optional CC recipient (string)
                Bcc:            nil, // Optional BCC recipients (map[string][]string)
            }
            
            // Data to populate the email template (can be of any type)
            data := map[string]string{
                "Name": "John Doe",
            }
            
            // Send the email
            // By default, gmail validates the server's TLS certificate, to avoid potential security issues due to invalid certificates, set 'false' during production.
            err := ctlgmail.SendEmail(config, credentials, data, true)
            
            if err != nil {
                log.Printf("Error sending email: %v", err)
            } else {
                log.Println("Email sent successfully!")
            }
    }
```
