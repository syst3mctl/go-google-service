# go-google-service

This Go package provides a convenient way to interact with Google APIs, leveraging the official Google API Client Library for Go (https://github.com/googleapis/google-api-go-client) and incorporating a custom mail sending functionality.

### Installation
To use go-google-service in your Go project, you can install it via go get:
```bash
go get github.com/syst3mctl/go-google-service
```

## Custom Mail Sending Function

This package provides a convenient way to send emails using customizable templates and SMTP credentials. It leverages the gomail library for email functionality.

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
    // By default, gmail validates the server's TLS certificate, to avoid potential security issues due to invalid certificates, set 'false' during production.
    err := ctlgmail.SendEmail(config, credentials, data, true)
    
    if err != nil {
        log.Printf("Error sending email: %v", err)
    } else {
        log.Println("Email sent successfully!")
    }
```

## Use populated data in email template
Define `htmlBody` in start of `html` or `tmpl` as follow
```html
{{define "htmlBody"}}
<!DOCTYPE html>
<html>
// rest of your html code
</html>
// close defined htmlBody
{{end}}
```

## Call  a Google Workspace API
Google Workspace quickstarts use the API client libraries to handle some details of the authentication and authorization flow. more information: (https://developers.google.com/docs/api/quickstart/go)


```go
// Import required packages

func main() {
        ctx := context.Background()
        // your credentials.json file
        b, err := os.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/documents.readonly")
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := ctlgmail.GetClient(config)

        srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                log.Fatalf("Unable to retrieve Docs client: %v", err)
        }

        // Prints the title of the requested doc:
        // https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
        docId := "195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE"
        doc, err := srv.Documents.Get(docId).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve data from document: %v", err)
        }
        fmt.Printf("The title of the doc is: %s\n", doc.Title)
}
```

## Important Notes:
- Replace placeholders in the SendingCredentials with your actual values.
- Security: It's highly recommended to use a secure SMTP server and avoid using skipSecurity=true in production environments. Always manage your email password securely.
- Error Handling: Consider implementing proper error handling practices like logging errors and returning appropriate error codes.
- Testing: Write unit tests to ensure the functionality of your email sending code.
