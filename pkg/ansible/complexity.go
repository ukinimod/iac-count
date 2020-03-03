package ansible

import (
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type ComplexityCalculator struct {
}

func playbookComplexity(playbook []AnsiblePlay) int {
	count := 0

	for k := range playbook {
		if playbook[k].PreTasks != nil {
			count += taskComplexity(playbook[k].PreTasks)
		}
		if playbook[k].Tasks != nil {
			count += taskComplexity(playbook[k].Tasks)
		}
		if playbook[k].PostTasks != nil {
			count += taskComplexity(playbook[k].PostTasks)
		}
	}

	return count
}

func taskComplexity(tasks []AnsibleTask) int {
	count := 0

	for _, task := range tasks {
		m := 1
		if task.Assert != nil {
			return 2
		}
		if task.WhenClause != nil {
			m *= 2
		}
		if task.Block != nil {
			m *= taskComplexity(task.Block)
		}
		if task.RescueBlock != nil {
			m *= 2 * taskComplexity(task.RescueBlock)
		}
		if task.AlwaysBlock != nil {
			m *= taskComplexity(task.AlwaysBlock)
		}
		/* 		if IsLoop(&task) {
		   			m *= 2
		   		}
		*/count += m
	}

	return count
}

func (calculator ComplexityCalculator) Analyze(path, content string) metrics.Metric {
	return metrics.Complexity{len(ReadHandlersString(content))}
}
