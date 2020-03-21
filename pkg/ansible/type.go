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

type NodeType string

const (
	File      = "file"
	Tasks     = "tasks"
	Handlers  = "handlers"
	GroupVars = "group_vars"
	HostVars  = "host_vars"
	Vars      = "vars"
	Defaults  = "defaults"
	Meta      = "meta"
	Playbook  = "playbook"
	Yml       = "yml"
	Role      = "role"
	Project   = "project"
	Directory = "directory"
)

func filetype(path string) NodeType {
	ext := filepath.Ext(path)

	if ext != ".yml" {
		return File
	}

	dir := filepath.Dir(path)

	if util.PathContainsDirName(dir, "tasks") {
		return Tasks
	}
	if util.PathContainsDirName(dir, "handlers") {
		return Handlers
	}
	if util.PathContainsDirName(dir, "group_vars") {
		return GroupVars
	}
	if util.PathContainsDirName(dir, "host_vars") {
		return HostVars
	}
	if util.PathContainsDirName(dir, "vars") {
		return Vars
	}
	if util.PathContainsDirName(dir, "defaults") {
		return Defaults
	}
	if util.PathContainsDirName(dir, "meta") {
		return Meta
	}
	if util.PathContainsDirName(dir, "playbooks") || dir == "" || dir == "." {
		return Playbook
	}

	return Yml
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
		index := sort.Search(
			len(necessaryRolePaths),
			func(i int) bool { return necessaryRolePaths[i] >= fileinfoContent.Name() }, // nolint:scopelint
		)
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
		if fileinfoContent.Name() == "roles" ||
			fileinfoContent.Name() == "site.yml" ||
			fileinfoContent.Name() == "playbooks" {
			return true
		}
	}

	return false
}

func dirtype(path string) NodeType {
	basename := filepath.Base(path)

	if strings.HasSuffix(path, "roles/"+basename) {
		hasRoleStructure(path)
		return Role
	}

	if hasAnsibleProjectStructure(path) {
		return Project
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

	return Directory
}
