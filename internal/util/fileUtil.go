package util

import (
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	mapping := make(map[string]string)
	mapping[".yaml"] = "text/yaml"
	mapping[".yml"] = "text/yaml"

	for key, value := range mapping {
		err := mime.AddExtensionType(key, value)

		if err != nil {
			log.Errorf("Error %v", err)
		}
	}
}

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
		log.Errorf("%s", err)
	}
	return numFiles
}

func SubdirCount(path string) int {
	numDirs := 0

	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Errorf("%s", err)
		return 0
	}
	for i := range fileinfo {
		if fileinfo[i].IsDir() && !IsHidden(fileinfo[i].Name()) {
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

func IsTextFile(path string) bool {
	mimeType := mime.TypeByExtension(filepath.Ext(path))
	return strings.HasPrefix(mimeType, "text/") ||
		strings.HasPrefix(mimeType, "application/json")
}

func IsYamlFile(path string) bool {
	mimeType := mime.TypeByExtension(filepath.Ext(path))
	return strings.HasPrefix(mimeType, "text/yaml")
}
