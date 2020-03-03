package ansible

import (
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type TaggedTasksCalculator struct {
}

func numberOfTaggedTasksInPlaybook(playbook []AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfTaggedTasks(playbook[k].PreTasks)
		count += numberOfTaggedTasks(playbook[k].Tasks)
		count += numberOfTaggedTasks(playbook[k].PostTasks)
	}

	return count
}

func numberOfTaggedTasks(task []AnsibleTask) int {
	count := 0
	for _, play := range task {
		if play.Tags != nil {
			count++
			if play.Block != nil {
				count += numberOfTaggedTasks(play.Block)
			}
			if play.RescueBlock != nil {
				count += numberOfTaggedTasks(play.RescueBlock)
			}
			if play.AlwaysBlock != nil {
				count += numberOfTaggedTasks(play.AlwaysBlock)
			}
		}
	}

	return count
}

func (calculator TaggedTasksCalculator) analyze(path, content string) metrics.Metric {
	return metrics.TaggedTasks{numberOfTaggedTasks(ReadTasksString(content))}
}
