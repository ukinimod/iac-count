package analyzer

import (
	"github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func numberOfPlays(playbook []input.AnsiblePlay) int {
	return len(playbook)
}

func numberOfHandlersInPlaybook(playbook []input.AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += len(playbook[k].Handlers)
	}

	return count
}

func numberOfTasksInPlaybook(playbook []input.AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfTasks(playbook[k].PreTasks)
		count += numberOfTasks(playbook[k].Tasks)
		count += numberOfTasks(playbook[k].PostTasks)
	}

	return count
}

func numberOfTaggedTasksInPlaybook(playbook []input.AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfTaggedTasks(playbook[k].PreTasks)
		count += numberOfTaggedTasks(playbook[k].Tasks)
		count += numberOfTaggedTasks(playbook[k].PostTasks)
	}

	return count
}

func numberOfCustomFactsInPlaybook(playbook []input.AnsiblePlay) int {
	count := 0

	for k := range playbook {
		count += numberOfCustomFacts(playbook[k].PreTasks)
		count += numberOfCustomFacts(playbook[k].Tasks)
		count += numberOfCustomFacts(playbook[k].PostTasks)
	}

	return count
}

func mccInTasksInPlaybook(playbook []input.AnsiblePlay) int {
	count := 0

	for k := range playbook {
		if playbook[k].PreTasks != nil {
			count += mcc(playbook[k].PreTasks)
		}
		if playbook[k].Tasks != nil {
			count += mcc(playbook[k].Tasks)
		}
		if playbook[k].PostTasks != nil {
			count += mcc(playbook[k].PostTasks)
		}
	}

	return count
}

func analyzePlaybook(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	metrics := analyzeYamlString(yamlString)

	playbook := input.ReadPlaybookString(yamlString)

	metrics[core.Plays] = numberOfPlays(playbook)
	metrics[core.Handlers] = numberOfHandlersInPlaybook(playbook)
	metrics[core.Tasks] = numberOfTasksInPlaybook(playbook)
	metrics[core.TaggedTasks] = numberOfTaggedTasksInPlaybook(playbook)
	metrics[core.CustomFacts] = numberOfCustomFactsInPlaybook(playbook)
	metrics[core.Complexity] = metrics[core.Plays] + metrics[core.Handlers] + mccInTasksInPlaybook(playbook)

	return metrics
}
