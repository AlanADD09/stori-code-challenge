package commands

import "transaction_processor/core"

type Command interface {
	Execute(ctx *core.ProcessorContext) error
}
