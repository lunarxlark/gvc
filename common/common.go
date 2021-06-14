package common

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var configFilePath = ""

func SetConfigFilePath(ctx *cli.Context) error {
	str, err := xdg.ConfigFile(filepath.Join(ctx.App.Name, "setting.json"))
	if err != nil {
		return errors.Wrap(err, "failed to set config file path")
	}
	configFilePath = str
	return nil
}

func GetConfigFilePath() string {
	return configFilePath
}
