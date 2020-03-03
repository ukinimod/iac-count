package ansible

import (
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type TasksCalculator struct {
}

func numberOfTasksInPlaybook(playbook []AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfTasks(playbook[k].PreTasks)
		count += numberOfTasks(playbook[k].Tasks)
		count += numberOfTasks(playbook[k].PostTasks)
	}

	return count
}

func numberOfTasks(task []AnsibleTask) int {
	count := 0

	for _, play := range task {
		count++
		if play.Block != nil {
			count += numberOfTasks(play.Block)
		}
		if play.RescueBlock != nil {
			count += numberOfTasks(play.RescueBlock)
		}
		if play.AlwaysBlock != nil {
			count += numberOfTasks(play.AlwaysBlock)
		}
	}

	return count
}

func (calculator TasksCalculator) analyze(path, content string) metrics.Metric {
	return metrics.Tasks{numberOfTasks(ReadTasksString(content))}
}
