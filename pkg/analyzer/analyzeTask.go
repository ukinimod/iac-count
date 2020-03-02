package analyzer

import (
	"github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func numberOfTaggedTasks(task []input.AnsibleTask) int {
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

func numberOfTasks(task []input.AnsibleTask) int {
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

func numberOfCustomFacts(task []input.AnsibleTask) int {
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

func isLoop(task *input.AnsibleTask) bool {
	return task.LoopClause != nil ||
		task.WithDictClause != nil ||
		task.WithFlattenedClause != nil ||
		task.WithIndexedItemsClause != nil ||
		task.WithItemsClause != nil ||
		task.WithListClause != nil ||
		task.WithSequenceClause != nil ||
		task.WithSubelementsClause != nil ||
		task.WithTogetherClause != nil ||
		task.WithCartesianClause != nil ||
		task.WithNestedClause != nil
}

func mcc(tasks []input.AnsibleTask) int {
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
			m *= mcc(task.Block)
		}
		if task.RescueBlock != nil {
			m *= 2 * mcc(task.RescueBlock)
		}
		if task.AlwaysBlock != nil {
			m *= mcc(task.AlwaysBlock)
		}
		if isLoop(&task) {
			m *= 2
		}
		count += m
	}

	return count
}

func analyzeTask(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	metrics := analyzeYamlString(yamlString)

	task := input.ReadTasksString(yamlString)

	metrics[core.Tasks] = numberOfTasks(task)
	metrics[core.TaggedTasks] = numberOfTaggedTasks(task)
	metrics[core.CustomFacts] = numberOfCustomFacts(task)
	metrics[core.Complexity] = mcc(task)

	return metrics
}
