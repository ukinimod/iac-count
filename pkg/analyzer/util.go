package analyzer

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func recursiveFileCount(path string) int {
	numFiles := 0
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				numFiles++
			}
			return nil
		})
	if err != nil && !os.IsNotExist(err) {
		log.Printf("[WARN] %s", err)
	}
	return numFiles
}

func subdirCount(path string) int {
	numDirs := 0

	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}
	for i := range fileinfo {
		if fileinfo[i].IsDir() {
			numDirs++
		}
	}

	return numDirs
}

func numberOfLines(s string) int {
	n := 0
	for _, r := range s {
		if r == '\n' {
			n++
		}
	}
	if len(s) > 0 && !strings.HasSuffix(s, "\n") {
		n++
	}
	return n
}

func pathContainsDirName(path, dirName string) bool {
	matched, err := regexp.MatchString("[.*/]?"+dirName+"[/.*]?", path)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}
	return matched
}

func filetype(path string) string {
	ext := filepath.Ext(path)

	if ext != ".yml" {
		return "file"
	}

	dir := filepath.Dir(path)

	if pathContainsDirName(dir, "tasks") {
		return "task"
	}
	if pathContainsDirName(dir, "handlers") {
		return "handler"
	}
	if pathContainsDirName(dir, "group_vars") {
		return "group_vars"
	}
	if pathContainsDirName(dir, "host_vars") {
		return "host_vars"
	}
	if pathContainsDirName(dir, "vars") {
		return "vars"
	}
	if pathContainsDirName(dir, "defaults") {
		return "defaults"
	}
	if pathContainsDirName(dir, "meta") {
		return "meta"
	}
	if pathContainsDirName(dir, "playbooks") || dir == "" || dir == "." {
		return "playbook"
	}

	return "yml"
}

// role must contain at least one of the following paths
var necessaryRolePaths []string = []string{
	"defaults",
	"files",
	"handlers",
	"meta",
	"tasks",
	"templates",
	"vars",
}

var optionalRolePaths []string = []string{
	"docs",
	"library",
	"tests",
}

func hasRoleStructure(root string) bool {
	fileinfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}
	for _, fileinfoContent := range fileinfo {
		index := sort.Search(len(necessaryRolePaths), func(i int) bool { return necessaryRolePaths[i] >= fileinfoContent.Name() })
		if index < len(necessaryRolePaths) && necessaryRolePaths[index] == fileinfoContent.Name() {
			return true
		}
	}

	log.Printf("[INFO] No role structure in %s.", root)

	return false
}

func isRoleDir(path string) bool {
	for i := range necessaryRolePaths {
		matched, err := regexp.MatchString(".*/roles/.*/"+necessaryRolePaths[i], path)

		if err != nil {
			log.Printf("[WARN] %s", err)
		}

		if matched {
			return true
		}
	}

	for i := range optionalRolePaths {
		matched, err := regexp.MatchString(".*/roles/.*/"+optionalRolePaths[i], path)

		if err != nil {
			log.Printf("[WARN] %s", err)
		}

		if matched {
			return true
		}
	}

	return false
}

func hasAnsibleProjectStructure(root string) bool {
	fileinfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}
	for _, fileinfoContent := range fileinfo {
		if fileinfoContent.Name() == "roles" || fileinfoContent.Name() == "site.yml" || fileinfoContent.Name() == "playbooks" {
			return true
		}
	}

	return false
}

func dirtype(path string) string {
	basename := filepath.Base(path)

	if strings.HasSuffix(path, "roles/"+basename) {
		hasRoleStructure(path)

		return "role"
	}

	if strings.HasSuffix(basename, "playbooks") || strings.HasSuffix(path, "playbooks/"+basename) {
		return "dir"
	}

	if basename == "roles" || strings.HasSuffix(basename, "plugins") || basename == "group_vars" || basename == "host_vars" {
		return "dir"
	}

	if hasAnsibleProjectStructure(path) {
		return "ansible_project"
	}

	if !isRoleDir(path) {
		log.Printf("[INFO] Unknown Directory %s\n", path)
	}

	return "dir"
}
