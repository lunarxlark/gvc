package cmd

import (
	"fmt"

	ff "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lunarxlark/gvc/models"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func ActSwitch(ctx *cli.Context) error {
	vis, err := models.GetVersionsInfoList()
	if err != nil {
		return errors.Wrap(err, "failed to get version info's list")
	}

	idx, err := ff.Find(vis, func(i int) string {
		return vis[i].Version
	})
	if err != nil {
		return errors.Wrap(err, "failed to fuzzy find version info")
	}

	if err := vis[idx].GoSwitch(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get switch version to '%s'", vis[idx].Version))
	}

	return nil
}
