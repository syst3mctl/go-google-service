package ctlgmail

// SMTPConfig defines the configuration for an SMTP server
type SMTPConfig struct {
	SMTPHost string // The hostname of the SMTP server
	SMTPPort int    // The port number of the SMTP server
}

// WithSMTPConfig creates a new SMTPConfig with custom settings
func (c *SMTPConfig) WithSMTPConfig(host string, port int) *SMTPConfig {
	// Create a new SMTPConfig instance with the provided host and port
	return &SMTPConfig{
		SMTPHost: host,
		SMTPPort: port,
	}
}

// NewSMTPConfig creates a new SMTPConfig with default settings
func NewSMTPConfig() *SMTPConfig {
	return &SMTPConfig{
		SMTPHost: "smtp.gmail.com", // Default Gmail SMTP server
		SMTPPort: 587,              // Default port for Gmail SMTP
	}
}
