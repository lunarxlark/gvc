package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/lunarxlark/gvc/common"
	"github.com/lunarxlark/gvc/internal"
	"github.com/lunarxlark/gvc/models"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func ActToolInit(ctx *cli.Context) error {
	log.Println("initialize gvc/config.json...")
	if internal.FileExists(common.GetConfigFilePath()) {
		log.Print("there is already setting.json")
		return nil
	}

	tcs := models.DefaultToolConfig
	b, err := json.MarshalIndent(tcs, "", "\t")
	if err != nil {
		return errors.Wrap(err, "failed to marshal tool config")
	}

	if err := ioutil.WriteFile(common.GetConfigFilePath(), b, 0644); err != nil {
		return errors.Wrap(err, "failed to write init config file")
	}
	return nil
}
