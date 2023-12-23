package dir

import (
	"os"
	"path/filepath"
)

func Cwd() string {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return cwd
}

func GetPathname(pathname string) string {
	cwd := Cwd()

	return filepath.Join(cwd, pathname)
}
