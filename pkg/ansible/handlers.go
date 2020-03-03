package ansible

import (
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type HandlersCalculator struct {
}

func numberOfHandlersInPlaybook(playbook []AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += len(playbook[k].Handlers)
	}

	return count
}

func (calculator HandlersCalculator) Analyze(path, content string) metrics.Metric {
	return metrics.Handlers{len(ReadHandlersString(content))}
}
