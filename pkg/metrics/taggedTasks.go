package metrics

type TaggedTasks struct {
	Val int
}

func (metric TaggedTasks) Description() string {
	return "Number of tasks that have tags"
}

func (metric TaggedTasks) Name() string {
	return "tagged_tasks"
}

func (metric TaggedTasks) Value() int {
	return metric.Val
}

func (metric TaggedTasks) add(additional Metric) Metric {
	if additional == nil {
		return TaggedTasks{
			Val: metric.Value(),
		}
	}

	return TaggedTasks{
		Val: metric.Value() + additional.Value(),
	}
}
