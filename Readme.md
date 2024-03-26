# go-google-service

This Go package provides a convenient way to interact with Google APIs, leveraging the official Google API Client Library for Go (https://github.com/googleapis/google-api-go-client) and incorporating a custom mail sending functionality.

### Installation
To use go-google-service in your Go project, you can install it via go get:
```bash
go get github.com/systemctl/go-google-service
```

## Custom Mail Sending Function

`go-google-service` includes a custom mail sending function to extend its functionality beyond the Google Docs API. You can utilize this function to send emails from your Go application with ease.

## Sending credentials
To send emails, you'll need to set up SendingCredentials which includes information like sender email, recipient email, password, subject, template path, and optional CC/BCC recipients.

## Sending an Email
Here's an example of how to send an email using the SendEmail function:

```go
    // Import required packages
    import (
    "log"
    
    "github.com/gomail/gomail"
    "your-username/email-sender" // Replace with your package import path
    )
    
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
    err := mail.SendEmail(config, credentials, data)
    
    if err != nil {
        log.Printf("Error sending email: %v", err)
    } else {
        log.Println("Email sent successfully!")
    }
```


## Important Notes:
- Replace placeholders in the SendingCredentials with your actual values.
- Security: It's highly recommended to use a secure SMTP server and avoid using skipSecurity=true in production environments. Always manage your email password securely.
- Error Handling: Consider implementing proper error handling practices like logging errors and returning appropriate error codes.
- Testing: Write unit tests to ensure the functionality of your email sending code.
