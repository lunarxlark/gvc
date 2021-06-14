package models

type BinaryFileInfo struct {
	OS      string   `json:"os"`
	Arch    string   `json:"arch"`
	Version string   `json:"version"`
	Sha256  string   `json:"sha256"`
	Size    int      `json:"size"`
	Kind    KindType `json:"kind"`
}

type KindType string

const (
	src       KindType = "source"
	archive   KindType = "archive"
	installer KindType = "installer"
)

func (k *KindType) String() string {
	return k.String()
}
