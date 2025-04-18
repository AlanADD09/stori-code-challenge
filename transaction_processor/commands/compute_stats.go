package commands

import (
	"fmt"
	"transaction_processor/core"
)

type ComputeStatsCommand struct{}

func (c *ComputeStatsCommand) Execute(ctx *core.ProcessorContext) error {
	summary := core.Summary{
		TransactionsPerMonth:  make(map[string]int),
		AverageCreditPerMonth: make(map[string]float64),
		AverageDebitPerMonth:  make(map[string]float64),
	}

	// Estructuras temporales para promedios por mes
	creditSums := make(map[string]float64)
	creditCounts := make(map[string]int)
	debitSums := make(map[string]float64)
	debitCounts := make(map[string]int)

	// Estructuras para promedio global
	var totalCreditSum float64
	var totalCreditCount int
	var totalDebitSum float64
	var totalDebitCount int

	for _, tx := range ctx.Transactions {
		monthKey := tx.Date.Format("2006-01") // "YYYY-MM"

		// Balance total
		summary.TotalBalance += tx.Amount

		// Conteo de transacciones por mes
		summary.TransactionsPerMonth[monthKey]++

		// Acumuladores por tipo
		if tx.Amount > 0 {
			creditSums[monthKey] += tx.Amount
			creditCounts[monthKey]++
			// Acumuladores globales
			totalCreditSum += tx.Amount
			totalCreditCount++
		} else {
			debitSums[monthKey] += tx.Amount
			debitCounts[monthKey]++
			// Acumuladores globales
			totalDebitSum += tx.Amount
			totalDebitCount++
		}
	}

	// Calcular promedios por mes
	for month, sum := range creditSums {
		count := creditCounts[month]
		if count > 0 {
			summary.AverageCreditPerMonth[month] = sum / float64(count)
		}
	}
	for month, sum := range debitSums {
		count := debitCounts[month]
		if count > 0 {
			summary.AverageDebitPerMonth[month] = sum / float64(count)
		}
	}

	// Calcular promedios globales
	if totalCreditCount > 0 {
		summary.AverageCreditTotal = totalCreditSum / float64(totalCreditCount)
	}
	if totalDebitCount > 0 {
		summary.AverageDebitTotal = totalDebitSum / float64(totalDebitCount)
	}

	ctx.Summary = summary

	fmt.Println("Computed statistics:")
	fmt.Printf("Total balance: %.2f\n", summary.TotalBalance)
	for month, count := range summary.TransactionsPerMonth {
		fmt.Printf("- %s: %d transactions\n", month, count)
		fmt.Printf("  Avg credit: %.2f\n", summary.AverageCreditPerMonth[month])
		fmt.Printf("  Avg debit: %.2f\n", summary.AverageDebitPerMonth[month])
	}
	fmt.Printf("Global average credit: %.2f\n", summary.AverageCreditTotal)
	fmt.Printf("Global average debit: %.2f\n", summary.AverageDebitTotal)

	return nil
}
