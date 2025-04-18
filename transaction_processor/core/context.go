package core

import "time"

type Transaction struct {
	Date   time.Time
	Amount float64
}

type Summary struct {
	TotalBalance          float64
	TransactionsPerMonth  map[string]int
	AverageCreditPerMonth map[string]float64
	AverageDebitPerMonth  map[string]float64
	AverageCreditTotal    float64
	AverageDebitTotal     float64
}

type SMTPConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Sender    string
	Recipient string
}

type ProcessorContext struct {
	FilePath     string
	Transactions []Transaction
	Summary      Summary
	EmailBody    string

	SMTPConfig SMTPConfig
}
