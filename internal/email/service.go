package email

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/gomail.v2"
)

// NewService creates a new email service with the given configuration
func NewService(config Config) (*Service, error) {
	service := &Service{
		config:    config,
		dialer:    gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.Username, config.Password),
		templates: make(map[string]*template.Template),
	}

	// Load templates
	err := service.loadTemplates()
	if err != nil {
		return nil, fmt.Errorf("failed to load email templates: %w", err)
	}

	service.initialized = true
	return service, nil
}

// loadTemplates loads all email templates from the templates directory
func (s *Service) loadTemplates() error {
	// Make sure templates directory exists
	if _, err := os.Stat(s.config.TemplatesDir); os.IsNotExist(err) {
		return fmt.Errorf("templates directory does not exist: %s", s.config.TemplatesDir)
	}

	// Create a shared template with common functions
	sharedFunctions := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("January 2, 2006")
		},
		"formatCurrency": func(amount float64) string {
			return fmt.Sprintf("$%.2f", amount)
		},
	}

	// Walk through templates directory
	err := filepath.Walk(s.config.TemplatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-HTML files
		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		// Get template name from filename (without extension)
		templateName := filepath.Base(path)
		templateName = templateName[:len(templateName)-len(filepath.Ext(templateName))]

		// Parse the template with shared functions
		tmpl, err := template.New(templateName).Funcs(sharedFunctions).ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", templateName, err)
		}

		// Store parsed template
		s.templates[templateName] = tmpl
		log.Printf("Loaded email template: %s", templateName)
		return nil
	})

	if err != nil {
		return err
	}

	if len(s.templates) == 0 {
		return fmt.Errorf("no email templates found in directory: %s", s.config.TemplatesDir)
	}

	return nil
}

// SendEmail sends an email using the specified template and data
func (s *Service) SendEmail(msg Message) error {
	if !s.initialized {
		return fmt.Errorf("email service not properly initialized")
	}

	// Get template
	tmpl, exists := s.templates[string(msg.TemplateName)]
	if !exists {
		return fmt.Errorf("email template not found: %s", msg.TemplateName)
	}

	// Create a new email message
	m := gomail.NewMessage()

	// Set sender
	m.SetAddressHeader("From", s.config.FromAddress, s.config.FromName)

	// Set recipients
	if len(msg.To) == 0 {
		return fmt.Errorf("no recipients specified")
	}

	for _, recipient := range msg.To {
		if recipient.Name != "" {
			m.SetAddressHeader("To", recipient.Email, recipient.Name)
		} else {
			m.SetHeader("To", recipient.Email)
		}
	}

	// Set CC recipients if any
	if len(msg.CC) > 0 {
		ccAddresses := make([]string, len(msg.CC))
		for i, cc := range msg.CC {
			if cc.Name != "" {
				ccAddresses[i] = m.FormatAddress(cc.Email, cc.Name)
			} else {
				ccAddresses[i] = cc.Email
			}
		}
		m.SetHeader("Cc", ccAddresses...)
	}

	// Set BCC recipients if any
	if len(msg.BCC) > 0 {
		bccAddresses := make([]string, len(msg.BCC))
		for i, bcc := range msg.BCC {
			if bcc.Name != "" {
				bccAddresses[i] = m.FormatAddress(bcc.Email, bcc.Name)
			} else {
				bccAddresses[i] = bcc.Email
			}
		}
		m.SetHeader("Bcc", bccAddresses...)
	}

	// Set subject
	m.SetHeader("Subject", msg.Subject)

	// Execute template with data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, msg.TemplateData); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	// Set HTML body
	m.SetBody("text/html", body.String())

	// Add attachments if any
	for _, attachment := range msg.Attachments {
		m.Attach(attachment.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(attachment.Data)
			return err
		}), gomail.SetHeader(map[string][]string{
			"Content-Type": {attachment.ContentType},
		}))
	}

	// Send email
	if err := s.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// SendRentReminder sends a rent reminder email
func (s *Service) SendRentReminder(tenant Address, propertyDetails interface{}, dueDate time.Time) error {
	return s.SendEmail(Message{
		To:           []Address{tenant},
		Subject:      "Rent Payment Reminder",
		TemplateName: string(RentReminder),
		TemplateData: map[string]interface{}{
			"TenantName":      tenant.Name,
			"PropertyDetails": propertyDetails,
			"DueDate":         dueDate,
		},
	})
}

// SendInvoiceToOwner sends an invoice email to a property owner
func (s *Service) SendInvoiceToOwner(owner Address, invoiceData interface{}, invoicePDF []byte) error {
	return s.SendEmail(Message{
		To:           []Address{owner},
		Subject:      "Property Management Invoice",
		TemplateName: string(InvoiceOwner),
		TemplateData: map[string]interface{}{
			"OwnerName":   owner.Name,
			"InvoiceData": invoiceData,
		},
		Attachments: []Attachment{
			{
				Filename:    "invoice.pdf",
				ContentType: "application/pdf",
				Data:        invoicePDF,
			},
		},
	})
}

// SendPropertyReport sends a property report email
func (s *Service) SendPropertyReport(recipients []Address, cc []Address, reportData interface{}, reportPDF []byte) error {
	return s.SendEmail(Message{
		To:           recipients,
		CC:           cc,
		Subject:      "Property Performance Report",
		TemplateName: string(PropertyReport),
		TemplateData: map[string]interface{}{
			"ReportData": reportData,
		},
		Attachments: []Attachment{
			{
				Filename:    "property_report.pdf",
				ContentType: "application/pdf",
				Data:        reportPDF,
			},
		},
	})
}

// CreateAttachmentFromFile creates an Attachment from a file
func CreateAttachmentFromFile(filePath string, contentType string) (Attachment, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Attachment{}, fmt.Errorf("failed to read file: %w", err)
	}

	return Attachment{
		Filename:    filepath.Base(filePath),
		ContentType: contentType,
		Data:        data,
	}, nil
}
