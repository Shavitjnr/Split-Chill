package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/Shavitjnr/split-chill-ai/cmd"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var (
	
	Version string

	
	CommitHash string

	
	BuildUnixTime string
)

func main() {
	settings.Version = Version
	settings.CommitHash = CommitHash
	settings.BuildTime = BuildUnixTime

	cmd := &cli.Command{
		Name:    "Split Chill AI",
		Usage:   "A lightweight, self-hosted personal finance app with a user-friendly interface and powerful bookkeeping features.",
		Version: GetFullVersion(),
		Commands: []*cli.Command{
			cmd.WebServer,
			cmd.Database,
			cmd.UserData,
			cmd.CronJobs,
			cmd.SecurityUtils,
			cmd.Utilities,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "conf-path",
				Usage: "Custom config `FILE` path",
			},
			&cli.BoolFlag{
				Name:  "no-boot-log",
				Usage: "Disable boot log",
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)

	if err != nil {
		log.Fatalf("Failed to run Split Chill AI with %s: %v", os.Args, err)
	}
}


func GetFullVersion() string {
	fullVersion := "Local Build"

	if Version != "" {
		fullVersion = Version
	}

	additionalInfos := make([]string, 0, 2)

	if CommitHash != "" {
		additionalInfos = append(additionalInfos, "commit "+CommitHash)
	}

	if BuildUnixTime != "" {
		unixTime, err := utils.StringToInt64(BuildUnixTime)

		if unixTime > 0 && err == nil {
			additionalInfos = append(additionalInfos, "build time "+utils.FormatUnixTimeToLongDateTimeInServerTimezone(unixTime))
		}
	}

	if len(additionalInfos) > 0 {
		fullVersion = fmt.Sprintf("%s (%s)", fullVersion, strings.Join(additionalInfos, ", "))
	}

	return fullVersion
}
