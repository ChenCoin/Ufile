package main

import (
	"path/filepath"
)

func check(filePath string) bool {
	if filePath=="/" { return true }
	path1, err := filepath.Abs("." + filePath)
	path2, err := filepath.Abs(".")
	path3 := filepath.Clean(filePath)
	if err != nil {	return false }
	return path1 == path2+path3
}
