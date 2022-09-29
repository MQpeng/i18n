package core

import (
	"path/filepath"
)

func ByGlob(glob_str string) (matches []string, err error){
	return filepath.Glob(glob_str)
}