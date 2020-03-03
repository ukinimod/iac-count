package ansible

import (
	"io/ioutil"
	"strings"

	"github.com/MaibornWolff/iac-count/internal/util"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

type PluginsCalculator struct {
}

func (calculator PluginsCalculator) Analyze(path, content string) metrics.Metric {
	n := 0

	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Warnf("%s", err)
	}
	for i := range fileinfo {
		if fileinfo[i].IsDir() && strings.HasSuffix(fileinfo[i].Name(), "plugins") && !util.IsHidden(fileinfo[i].Name()) {
			n += util.SubdirCount(path + "/" + fileinfo[i].Name())
		}
	}

	return metrics.Plugins{n}
}
