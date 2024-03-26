# go-google-service

This Go package provides a convenient way to interact with Google APIs, leveraging the official Google API Client Library for Go (https://github.com/googleapis/google-api-go-client) and incorporating a custom mail sending functionality.

### Installation
To use go-google-service in your Go project, you can install it via go get:
```bash
go get github.com/systemctl/go-google-service
```

## Custom Mail Sending Function

`go-google-service` includes a custom mail sending function to extend its functionality beyond the Google Docs API. You can utilize this function to send emails from your Go application with ease.


## Sending email templates
Example usage how to use this package and how to set up SMTPConfiguration
```go
func main(){
    // As a default, SMTP server start with google smtp server with default port.
    // you can change it whenever you want, just call function WithSMTPConfig() and pass parameters 
    config := NewSMTPConfig()

    // set up SendingCredentials
    sc := &SendingCredentials{
        From:           "your_email@gmail.com",
        To:             "your_email@gmail.com",
        SenderPassword: "your_password",
        Subject:        "subject",
        Template:       "your_template.tmpl",
        Cc:             "this filed is optional",
        Bcc:            "this field is optional"
    }
	
    // example data
    data := map[string]any{
        "userEmail": sc.To,
    }
	
    // Call SendEmail and pass parameters
    err := SendEmail(config, sc, data, true)
    if err != nil {
        panic(err)
    }
}
```