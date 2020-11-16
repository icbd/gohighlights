package utils

import (
	"os"
	"strings"
)

const RootName = "gohighlights"

func RootPath() (rootPath string) {
	dir, _ := os.Getwd()
	for _, s := range strings.Split(dir, string(os.PathSeparator)) {
		rootPath += string(os.PathSeparator)
		rootPath += s
		if s == RootName {
			break
		}
	}
	return rootPath
}
