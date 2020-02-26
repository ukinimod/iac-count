package analyzer

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/ukinimod/iac-count/pkg/core"
)

func analyzeRole(path string) map[string]int {
	metrics := analyzeDir(path)

	metrics[core.Templates] = recursiveFileCount(path + "/templates")
	metrics[core.StaticFiles] = recursiveFileCount(path + "/files")
	metrics[core.Plugins] = 0

	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}
	for i := range fileinfo {
		if fileinfo[i].IsDir() && strings.HasSuffix(fileinfo[i].Name(), "plugins") {
			metrics[core.Plugins] += subdirCount(path + "/" + fileinfo[i].Name())
		}
	}

	return metrics
}
