package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lunarxlark/gvc/common"
	"github.com/pkg/errors"
)

type ToolInfo struct {
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	CVS    CVSKind `json:"cvs"`
	Enable bool    `json:"enable"`
}

type CVSKind string

const (
	KindGit CVSKind = "git"
	KindSvn CVSKind = "svn"
)

var DefaultToolConfig = []*ToolInfo{
	//{
	//	Name:   "x/tools",
	//	URL:    "golang.org/x/tools/...",
	//	CVS:    kindGit,
	//	Enable: true,
	//},
	{
		Name:   "gopls",
		URL:    "golang.org/x/tools/gopls",
		CVS:    KindGit,
		Enable: true,
	},
}

func GetToolInfoList() (tis []*ToolInfo, err error) {
	f, err := os.Open(common.GetConfigFilePath())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config.json")
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read all config.json")
	}

	t := []*ToolInfo{}
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config.json")
	}

	for _, v := range t {
		if v.Enable {
			tis = append(tis, v)
		}
	}

	return tis, nil
}

var ignoreList = []string{
	"golang.org",
	"honnef.co",
}

func (ti *ToolInfo) GoGet() error {
	// TODO : ダウンロード元のpackageをどうやって特定するか。
	hasDomain := false
	for _, i := range ignoreList {
		if strings.HasPrefix(ti.URL, i) {
			hasDomain = true
		}
	}

	packages := ""
	if !hasDomain {
		packages = filepath.Join("github.com", ti.URL)
	}
	cmd := exec.Command("go", "install", packages)
	stdouterr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(stdouterr))
		return errors.Wrap(err, fmt.Sprintf("failed to execute command '%s'", cmd.String()))
	}
	return nil
}
