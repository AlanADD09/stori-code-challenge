package utils

import (
	"fmt"
	"os"
)

type FileConfig struct {
	// Directorio de archivos CSV
	Directory string

	// Configuración SMTP
	SmtpHost       string
	SmtpPort       string
	SmtpUsername   string
	SmtpPass       string
	SenderEmail    string
	RecipientEmail string

	// Configuración MSSQL
	MssqlHost     string
	MssqlPort     string
	MssqlUser     string
	MssqlPassword string
	MssqlDatabase string

	// Configuración del server HTTP
	Port string
}

func LoadConfigFromEnv() (FileConfig, error) {
	config := FileConfig{
		Directory:      os.Getenv("FILE_DIRECTORY"),
		SmtpHost:       os.Getenv("SMTP_HOST"),
		SmtpPort:       os.Getenv("SMTP_PORT"),
		SmtpUsername:   os.Getenv("SMTP_USERNAME"),
		SmtpPass:       os.Getenv("SMTP_PASS"),
		SenderEmail:    os.Getenv("SENDER_EMAIL"),
		RecipientEmail: os.Getenv("RECIPIENT_EMAIL"),
		MssqlHost:      os.Getenv("MSSQL_HOST"),
		MssqlPort:      os.Getenv("MSSQL_PORT"),
		MssqlUser:      os.Getenv("MSSQL_USER"),
		MssqlPassword:  os.Getenv("MSSQL_PASSWORD"),
		MssqlDatabase:  os.Getenv("MSSQL_NAME"),
		Port:           os.Getenv("PORT"),
	}

	// Validar configuración mínima (lo importante para correr la app)
	if config.Directory == "" || config.SmtpHost == "" || config.SmtpPort == "" ||
		config.SmtpPass == "" || config.SenderEmail == "" || config.RecipientEmail == "" ||
		config.SmtpUsername == "" || config.MssqlHost == "" || config.MssqlPort == "" ||
		config.MssqlUser == "" || config.MssqlPassword == "" || config.MssqlDatabase == "" {
		return FileConfig{}, fmt.Errorf(
			"configuración incompleta en variables de entorno:" + "\n" +
				"Directory: " + config.Directory + "\n" +
				"SMTP Host: " + config.SmtpHost + "\n" +
				"SMTP Port: " + config.SmtpPort + "\n" +
				"SMTP Username: " + config.SmtpUsername + "\n" +
				"SMTP Pass: " + config.SmtpPass + "\n" +
				"Sender Email: " + config.SenderEmail + "\n" +
				"Recipient Email: " + config.RecipientEmail + "\n" +
				"MSSQL Host: " + config.MssqlHost + "\n" +
				"MSSQL Port: " + config.MssqlPort + "\n" +
				"MSSQL User: " + config.MssqlUser + "\n" +
				"MSSQL Password: (hidden)" + "\n" +
				"MSSQL Database: " + config.MssqlDatabase + "\n",
		)
	}

	return config, nil
}
