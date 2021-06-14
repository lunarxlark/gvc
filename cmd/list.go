package cmd

import (
	"fmt"

	"github.com/lunarxlark/gvc/models"
	"github.com/urfave/cli/v2"
)

func ActList(ctx *cli.Context) error {
	vis, err := models.GetVersionsInfoList()
	if err != nil {
		return nil
	}

	fmt.Println("versions\tstable\tos\tarch\tsha256\tsize\tkind")
	for _, vi := range vis {
		for _, bi := range vi.File {
			fmt.Printf("%s\t%t\t%s\t%s\t%s\t%d\t%s\n", vi.Version, vi.Stable, bi.OS, bi.Arch, bi.Sha256, bi.Size, bi.Kind)
		}
	}
	return nil
}
