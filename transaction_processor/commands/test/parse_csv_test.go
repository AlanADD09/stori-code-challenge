package commands

import (
	"fmt"
	"os"
	"testing"

	"transaction_processor/commands"
	"transaction_processor/core"
)

func TestParseCSVCommand_Execute(t *testing.T) {
	content := `0,7/25,+60.5,
				0,7/25,-20.0,
				0,8/01,+35.0,
				0,8/15,-10.5,
				`
	tmpFile, err := os.CreateTemp("", "test_transactions_*.csv")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}
	tmpFile.Close()

	ctx := &core.ProcessorContext{FilePath: tmpFile.Name()}

	cmd := &commands.ParseCSVCommand{}
	if err := cmd.Execute(ctx); err != nil {
		t.Errorf("Execute() returned error: %v", err)
	}

	fmt.Println("Parsed Transactions:")
	for i, tx := range ctx.Transactions {
		fmt.Printf("Transaction %d: %+v\n", i, tx)
	}

	if len(ctx.Transactions) != 4 {
		t.Errorf("expected 4 transactions, got %d", len(ctx.Transactions))
	}

	expectedAmounts := []float64{60.5, -20.0, 35.0, -10.5}
	for i, tx := range ctx.Transactions {
		if tx.Amount != expectedAmounts[i] {
			t.Errorf("expected amount %v at index %d, got %v", expectedAmounts[i], i, tx.Amount)
		}
	}
}
