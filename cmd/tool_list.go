package cmd

import (
	"fmt"

	"github.com/lunarxlark/gvc/models"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func ActToolList(ctx *cli.Context) error {
	tcs, err := models.GetToolInfoList()
	if err != nil {
		return errors.Wrap(err, "failed to get tool list")
	}

	fmt.Printf("%s\t%s\n", "Name", "URL")
	for _, tc := range tcs {
		fmt.Printf("%s\t%s\n", tc.Name, tc.URL)
	}

	return nil
}
