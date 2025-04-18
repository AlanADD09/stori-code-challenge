package commands

import (
	"fmt"
	"transaction_processor/core"
	"transaction_processor/utils"
)

type SaveToDBCommand struct{}

func (s *SaveToDBCommand) Execute(ctx *core.ProcessorContext) error {
	for _, tx := range ctx.Transactions {
		query := `
            INSERT INTO transactions (date, amount)
            VALUES (@date, @amount)
        `
		variables := []utils.SqlArgs{
			{Name: "date", Value: tx.Date},
			{Name: "amount", Value: tx.Amount},
		}

		_, err := utils.DoMutation(query, variables)
		if err != nil {
			return fmt.Errorf("failed to insert transaction: %w", err)
		}
	}

	fmt.Println("All transactions saved successfully to database.")
	return nil
}
