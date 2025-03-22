package email

import (
	"gopkg.in/gomail.v2"
	"html/template"
)

// Config holds email service configuration
type Config struct {
	SMTPHost     string
	SMTPPort     int
	Username     string
	Password     string
	FromAddress  string
	FromName     string
	TemplatesDir string
}

// Service represents the email service
type Service struct {
	config      Config
	dialer      *gomail.Dialer
	templates   map[string]*template.Template
	initialized bool
}

// Attachment represents an email attachment
type Attachment struct {
	Filename    string
	ContentType string
	Data        []byte
}

// Address represents an email address with optional name
type Address struct {
	Email string
	Name  string
}

// Message represents an email message to be sent
type Message struct {
	To           []Address
	CC           []Address
	BCC          []Address
	Subject      string
	TemplateName string
	TemplateData interface{}
	Attachments  []Attachment
}

// TemplateType defines the available email templates
type TemplateType string

const (
	RentReminder     TemplateType = "rent_reminder"
	InvoiceOwner     TemplateType = "invoice_owner"
	InvoiceCustomer  TemplateType = "invoice_customer"
	MaintenanceAlert TemplateType = "maintenance_alert"
	LeaseRenewal     TemplateType = "lease_renewal"
	PropertyReport   TemplateType = "property_report"
	WelcomeEmail     TemplateType = "welcome_email"
	PaymentReceipt   TemplateType = "payment_receipt"
)
