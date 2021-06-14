package cmd

import (
	"fmt"
	"log"

	"github.com/lunarxlark/gvc/models"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func ActToolInstall(ctx *cli.Context) error {
	tis, err := models.GetToolInfoList()
	if err != nil {
		return errors.Wrap(err, "failed to get tool config")
	}

	for _, ti := range tis {
		if ti.Enable {
			log.Println(fmt.Sprintf("installing %s ... ", ti.Name))
			if err := ti.GoGet(); err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to install tool '%s'", ti.Name))
			}
		}
	}
	return nil
}