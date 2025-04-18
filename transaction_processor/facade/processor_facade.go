package facade

import (
	"transaction_processor/commands"
	"transaction_processor/core"
	"transaction_processor/utils"
)

type ProcessorFacade struct {
	Context *core.ProcessorContext
	Queue   []commands.Command
}

func NewProcessorFacade(config utils.FileConfig) *ProcessorFacade {
	ctx := &core.ProcessorContext{FilePath: config.Directory + "/transactions.csv"}
	ctx.SMTPConfig = core.SMTPConfig{
		Host:      config.SmtpHost,
		Port:      config.SmtpPort,
		Username:  config.SmtpUsername,
		Password:  config.SmtpPass,
		Sender:    config.SenderEmail,
		Recipient: config.RecipientEmail,
	}
	return &ProcessorFacade{
		Context: ctx,
		Queue: []commands.Command{
			&commands.ParseCSVCommand{},
			&commands.ComputeStatsCommand{},
			&commands.SendEmailCommand{},
			&commands.SaveToDBCommand{},
		},
	}
}

func (p *ProcessorFacade) Run() error {
	for _, cmd := range p.Queue {
		if err := cmd.Execute(p.Context); err != nil {
			return err
		}
	}
	return nil
}
