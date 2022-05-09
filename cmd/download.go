package cmd

import (
	ff "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lunarxlark/gvc/models"
	"github.com/urfave/cli/v2"
)

func ActDownload(ctx *cli.Context) error {
	vis, err := models.GetVersionsInfoList()
	if err != nil {
		return err
	}

	idx, err := ff.Find(vis, func(i int) string {
		return vis[i].Version
	})
	if err != nil {
		return err
	}

	if err := vis[idx].GoInstall(); err != nil {
		return err
	}
	if err := vis[idx].GoDownload(); err != nil {
		return err
	}
	return nil
}
