package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/lunarxlark/gvc/cmd"
	"github.com/lunarxlark/gvc/common"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	version = "0.0.2"
)

func main() {
	app := cli.App{
		Name:    "gvc",
		Usage:   "go version control",
		Version: version,
		Before:  before,
		Commands: []*cli.Command{
			{Name: "list", Action: cmd.ActList},
			{Name: "download", Action: cmd.ActDownload},
			{Name: "switch", Action: cmd.ActSwitch},
			// subcommand for 'tools' in config.json
			{
				Name: "tool", Subcommands: []*cli.Command{
					{Name: "init", Action: cmd.ActToolInit},
					{Name: "list", Action: cmd.ActToolList},
					{Name: "install", Action: cmd.ActToolInstall},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

func before(ctx *cli.Context) error {
	if err := common.SetConfigFilePath(ctx); err != nil {
		return errors.Wrap(err, "failed to set config file path")
	}

	if err := os.MkdirAll(filepath.Dir(common.GetConfigFilePath()), 0755); err != nil {
		return errors.Wrap(err, "failed to mkdir for config")
	}

	return nil
}
