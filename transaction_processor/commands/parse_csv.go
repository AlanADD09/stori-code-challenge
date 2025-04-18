package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"transaction_processor/core"
)

type ParseCSVCommand struct{}

func (p *ParseCSVCommand) Execute(ctx *core.ProcessorContext) error {
	file, err := os.Open(ctx.FilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var transactions []core.Transaction

	currentYear := time.Now().Year()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Ejemplo línea: 0,7/25,+60.5,
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			fmt.Printf("Skipping malformed line: %s\n", line)
			continue
		}

		dateStr := strings.TrimSpace(fields[1])
		amountStr := strings.TrimSpace(fields[2])

		// Parsear fecha: asumir año actual
		date, err := time.Parse("1/2/2006", fmt.Sprintf("%s/%d", dateStr, currentYear))
		if err != nil {
			fmt.Printf("Skipping invalid date: %s (%v)\n", dateStr, err)
			continue
		}

		// Parsear monto: puede tener + o -
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Printf("Skipping invalid amount: %s (%v)\n", amountStr, err)
			continue
		}

		transactions = append(transactions, core.Transaction{
			Date:   date,
			Amount: amount,
		})
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	ctx.Transactions = transactions
	fmt.Printf("Parsed %d transactions.\n", len(transactions))
	return nil
}
