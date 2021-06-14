package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	versionListURL = "https://golang.org/dl/?mode=json&include=all"
)

type VersionInfo struct {
	Version string            `json:"version"`
	Stable  bool              `json:"stable"`
	File    []*BinaryFileInfo `json:"files"`
}

func (vi *VersionInfo) GoGet() error {
	log.Print(fmt.Sprintf("getting %s ...", vi.Version))
	cmd := exec.Command("go", "get", fmt.Sprintf("golang.org/dl/%s", vi.Version))
	stdouterr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdouterr))
		return errors.Wrap(err, fmt.Sprintf("failed to get %s with command '%s'", vi.Version, cmd.String()))
	}
	return nil
}

func (vi *VersionInfo) GoDownload() error {
	log.Print(fmt.Sprintf("downloading %s ...", vi.Version))
	cmd := exec.Command(vi.Version, "download")
	stdouterr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdouterr))
		return errors.Wrap(err, fmt.Sprintf("failed to download %s with command '%s'", vi.Version, cmd.String()))
	}
	return nil
}

func (vi *VersionInfo) GoSwitch() error {
	log.Println(fmt.Sprintf("remake symbolic link for %s...", vi.Version))
	binDir := os.Getenv("GOBIN")

	oldname := filepath.Join(binDir, "go")
	if _, err := os.Lstat(oldname); err == nil {
		if err := os.Remove(oldname); err != nil {
			return errors.Wrap(err, "failed to unlink")
		}
	}

	newname := filepath.Join(binDir, vi.Version)
	if _, err := os.Stat(newname); err != nil {
		if os.IsNotExist(err) {
			return errors.Wrap(err, fmt.Sprintf("%s is not installed", vi.Version))
		} else {
			return errors.Wrap(err, "failed to stat")
		}
	}

	if err := os.Symlink(newname, oldname); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to make symblic link for %s", vi.Version))
	}
	return nil
}

func GetVersionsInfoList() (vis []*VersionInfo, err error) {
	resp, err := http.Get(versionListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &vis); err != nil {
		return nil, err
	}
	return vis, nil

}
