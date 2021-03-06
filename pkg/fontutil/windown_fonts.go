// +build windows

// Package fontutil .
package fontutil

import (
	"os"
	"path"

	"github.com/yeqown/infrastructure/pkg/fs"
)

var (
	winFontPath string
)

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	winFontPath = path.Join(homedir, "AppData", "Local", "Microsoft", "Windows", "Fonts")
}

// GetSysDefaultFont . return current system default font
func GetSysDefaultFont() string {
	return ""
}

// GetSysFontList get font list from curretn system
func GetSysFontList() (fonts []string) {
	files := fs.ListFiles(winFontPath, fs.IgnoreDirFilter())
	if len(files) != 0 {
		fonts = make([]string, len(files))
		// true: handle files
		for idx, p := range files {
			_, fonts[idx] = path.Split(p)
		}
	}

	return
}

// AssemFontPath .
func AssemFontPath(fontfile string) string {
	return path.Join(winFontPath, fontfile)
}
