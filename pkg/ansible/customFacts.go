package ansible

import (
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type CustomFactsCalculator struct {
}

func numberOfCustomFactsInPlaybook(playbook []AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfCustomFacts(playbook[k].PreTasks)
		count += numberOfCustomFacts(playbook[k].Tasks)
		count += numberOfCustomFacts(playbook[k].PostTasks)
	}

	return count
}

func numberOfCustomFacts(task []AnsibleTask) int {
	count := 0

	for _, play := range task {
		if play.SetFact != nil {
			for key := range play.SetFact {
				if key != "cacheable" {
					count++
				}
			}
		}
		if play.Block != nil {
			count += numberOfCustomFacts(play.Block)
		}
		if play.RescueBlock != nil {
			count += numberOfCustomFacts(play.RescueBlock)
		}
		if play.AlwaysBlock != nil {
			count += numberOfCustomFacts(play.AlwaysBlock)
		}
	}

	return count
}

func (calculator CustomFactsCalculator) analyze(path, content string) metrics.Metric {
	return metrics.CustomFacts{numberOfCustomFacts(ReadTasksString(content))}
}
