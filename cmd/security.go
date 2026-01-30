package cmd

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


var SecurityUtils = &cli.Command{
	Name:  "security",
	Usage: "Split Chill AI security utilities",
	Commands: []*cli.Command{
		{
			Name:   "gen-secret-key",
			Usage:  "Generate a random secret key",
			Action: bindAction(genSecretKey),
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:        "length",
					Aliases:     []string{"l"},
					Required:    false,
					DefaultText: "32",
					Usage:       "The length of secret key",
				},
			},
		},
	},
}

func genSecretKey(c *core.CliContext) error {
	length := c.Int("length")

	if length <= 0 {
		length = 32
	}

	secretKey, err := utils.GetRandomNumberOrLetter(length)

	if err != nil {
		return err
	}

	fmt.Printf("[Secret Key] %s\n", secretKey)

	return nil
}
