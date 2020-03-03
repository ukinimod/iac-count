package ansible

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/MaibornWolff/iac-count/internal/util"
	log "github.com/sirupsen/logrus"
)

func filetype(path string) string {
	ext := filepath.Ext(path)

	if ext != ".yml" {
		return "file"
	}

	dir := filepath.Dir(path)

	if util.PathContainsDirName(dir, "tasks") {
		return "task"
	}
	if util.PathContainsDirName(dir, "handlers") {
		return "handler"
	}
	if util.PathContainsDirName(dir, "group_vars") {
		return "group_vars"
	}
	if util.PathContainsDirName(dir, "host_vars") {
		return "host_vars"
	}
	if util.PathContainsDirName(dir, "vars") {
		return "vars"
	}
	if util.PathContainsDirName(dir, "defaults") {
		return "defaults"
	}
	if util.PathContainsDirName(dir, "meta") {
		return "meta"
	}
	if util.PathContainsDirName(dir, "playbooks") || dir == "" || dir == "." {
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

func hasRoleStructure(root string) bool {
	fileinfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Warnf("%s", err)
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
		matched, err := regexp.MatchString("[.*/]?roles/.*/"+necessaryRolePaths[i], path)

		if err != nil {
			log.Warnf("%s", err)
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
		log.Warnf("%s", err)
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

	if hasAnsibleProjectStructure(path) {
		return "ansible_project"
	}

	if !(strings.HasSuffix(basename, "playbooks") ||
		strings.HasSuffix(path, "playbooks/"+basename) ||
		strings.HasSuffix(basename, "plugins") ||
		basename == "roles" ||
		basename == "group_vars" ||
		basename == "host_vars") &&
		!isRoleDir(path) {
		log.Infof("Unknown Directory %s\n", path)
	}

	return "dir"
}
