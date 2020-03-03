package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func IsHidden(path string) bool {
	basename := filepath.Base(path)
	return basename != "." && strings.HasPrefix(basename, ".")
}

func ParentPath(path string) string {
	basename := filepath.Base(path)

	parentPath := strings.TrimSuffix(path, string(filepath.Separator)+basename)
	if parentPath == "" || parentPath == path {
		parentPath = "."
	}

	return parentPath
}

func RecursiveFileCount(path string) int {
	info, err := os.Stat(path)
	if err != nil {
		log.Warnf("%s", err)
		return 0
	} else if !info.IsDir() {
		return 1
	}

	numFiles := 0
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && !IsHidden(path) {
				numFiles++
			}
			return nil
		})
	if err != nil && !os.IsNotExist(err) {
		log.Warnf("%s", err)
	}
	return numFiles
}

func SubdirCount(path string) int {
	numDirs := 0

	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Warnf("%s", err)
	}
	for i := range fileinfo {
		if fileinfo[i].IsDir() && !IsHidden(path) {
			numDirs++
		}
	}

	return numDirs
}

func PathContainsDirName(path, dirName string) bool {
	matched, err := regexp.MatchString("[.*/]?"+dirName+"[/.*]?", path)
	if err != nil {
		log.Warnf("%s", err)
	}
	return matched
}
