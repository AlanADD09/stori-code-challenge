package commands

import (
	"testing"
	"time"

	"transaction_processor/commands"
	"transaction_processor/core"
)

func TestComputeStatsCommand_Execute(t *testing.T) {
	// Crear transacciones simuladas
	transactions := []core.Transaction{
		{Date: time.Date(2025, 7, 25, 0, 0, 0, 0, time.UTC), Amount: 60.5},
		{Date: time.Date(2025, 7, 25, 0, 0, 0, 0, time.UTC), Amount: -20.0},
		{Date: time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC), Amount: 35.0},
		{Date: time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC), Amount: -10.5},
	}

	ctx := &core.ProcessorContext{
		Transactions: transactions,
	}

	cmd := &commands.ComputeStatsCommand{}
	if err := cmd.Execute(ctx); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	summary := ctx.Summary

	// Validar balance total
	expectedBalance := 60.5 - 20.0 + 35.0 - 10.5
	if summary.TotalBalance != expectedBalance {
		t.Errorf("expected balance %.2f, got %.2f", expectedBalance, summary.TotalBalance)
	}

	// Validar transacciones por mes
	if summary.TransactionsPerMonth["2025-07"] != 2 {
		t.Errorf("expected 2 transactions for July, got %d", summary.TransactionsPerMonth["2025-07"])
	}
	if summary.TransactionsPerMonth["2025-08"] != 2 {
		t.Errorf("expected 2 transactions for August, got %d", summary.TransactionsPerMonth["2025-08"])
	}

	// Validar promedios
	if summary.AverageCreditPerMonth["2025-07"] != 60.5 {
		t.Errorf("expected July avg credit 60.5, got %.2f", summary.AverageCreditPerMonth["2025-07"])
	}
	if summary.AverageDebitPerMonth["2025-07"] != -20.0 {
		t.Errorf("expected July avg debit -20.0, got %.2f", summary.AverageDebitPerMonth["2025-07"])
	}
	if summary.AverageCreditPerMonth["2025-08"] != 35.0 {
		t.Errorf("expected August avg credit 35.0, got %.2f", summary.AverageCreditPerMonth["2025-08"])
	}
	if summary.AverageDebitPerMonth["2025-08"] != -10.5 {
		t.Errorf("expected August avg debit -10.5, got %.2f", summary.AverageDebitPerMonth["2025-08"])
	}
}
