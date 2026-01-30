package core

import (
	"context"

	"github.com/urfave/cli/v3"
)


type CliContext struct {
	context.Context
	command *cli.Command
}


func (c *CliContext) GetContextId() string {
	return ""
}


func (c *CliContext) GetClientLocale() string {
	return ""
}


func (c *CliContext) Bool(name string) bool {
	return c.command.Bool(name)
}


func (c *CliContext) Int(name string) int {
	return c.command.Int(name)
}


func (c *CliContext) Int64(name string) int64 {
	return c.command.Int64(name)
}


func (c *CliContext) String(name string) string {
	return c.command.String(name)
}


func WrapCilContext(ctx context.Context, cmd *cli.Command) *CliContext {
	return &CliContext{
		Context: ctx,
		command: cmd,
	}
}
