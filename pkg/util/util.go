package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetOutPath(path string) string {
	ofileName := filepath.Base(path)
	opath := os.Getenv("HTMGENIE_OPATH")
	// TODO: File permissions
	if opath == "" {
		opath = "."
	}

	ofileName = replaceLast(ofileName, ".md", ".html")
	opath = fmt.Sprintf("%s/%s", opath, ofileName)

	return opath
}

func replaceLast(x, y, z string) (x2 string) {
	i := strings.LastIndex(x, y)
	if i == -1 {
		return x
	}
	return x[:i] + z + x[i+len(y):]
}
