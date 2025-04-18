package commands

import (
	"fmt"
	"strconv"
	"transaction_processor/core"

	"gopkg.in/gomail.v2"
)

type SendEmailCommand struct{}

func (s *SendEmailCommand) Execute(ctx *core.ProcessorContext) error {
	var emailBody string

	emailBody += "<html><body>"
	emailBody += "<h2>Resumen de Transacciones</h2>"
	emailBody += fmt.Sprintf("<p><strong>Saldo Total:</strong> %.2f</p>", ctx.Summary.TotalBalance)

	emailBody += "<h3>Transacciones por mes:</h3><ul>"
	for month, count := range ctx.Summary.TransactionsPerMonth {
		emailBody += fmt.Sprintf("<li><strong>%s:</strong> %d transacciones</li>", month, count)
	}
	emailBody += "</ul>"

	emailBody += "<h3>Promedios por mes:</h3><ul>"
	for month := range ctx.Summary.TransactionsPerMonth {
		avgCredit := ctx.Summary.AverageCreditPerMonth[month]
		avgDebit := ctx.Summary.AverageDebitPerMonth[month]
		emailBody += fmt.Sprintf("<li><strong>%s:</strong> Promedio Crédito: %.2f | Promedio Débito: %.2f</li>", month, avgCredit, avgDebit)
	}
	emailBody += "</ul>"

	// Aquí agregamos los nuevos campos globales
	emailBody += "<h3>Promedios Globales:</h3><ul>"
	emailBody += fmt.Sprintf("<li><strong>Promedio Crédito Total:</strong> %.2f</li>", ctx.Summary.AverageCreditTotal)
	emailBody += fmt.Sprintf("<li><strong>Promedio Débito Total:</strong> %.2f</li>", ctx.Summary.AverageDebitTotal)
	emailBody += "</ul>"

	emailBody += "<img src='https://play-lh.googleusercontent.com/oXTAgpljdbV5LuAOt1NP9_JafUZe9BNl7pwQ01ndl4blYL4N4IQh4-n456P5l_hc1A' alt='Stori Logo' style='width:150px;margin-top:20px;'>"
	emailBody += "</body></html>"

	ctx.EmailBody = emailBody

	// --- Enviar el correo usando gomail ---
	m := gomail.NewMessage()
	m.SetHeader("From", ctx.SMTPConfig.Sender)
	m.SetHeader("To", ctx.SMTPConfig.Recipient)
	m.SetHeader("Subject", "Resumen de Transacciones")
	m.SetBody("text/html", ctx.EmailBody)

	port, err := strconv.Atoi(ctx.SMTPConfig.Port)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	d := gomail.NewDialer(
		ctx.SMTPConfig.Host,
		port,
		ctx.SMTPConfig.Username,
		ctx.SMTPConfig.Password,
	)

	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully to", ctx.SMTPConfig.Recipient)
	return nil
}
