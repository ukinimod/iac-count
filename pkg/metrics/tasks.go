package metrics

type Tasks struct {
	Val int
}

func (metric Tasks) Description() string {
	return "Number of tasks"
}

func (metric Tasks) Name() string {
	return "tasks"
}

func (metric Tasks) Value() int {
	return metric.Val
}

func (metric Tasks) add(additional Metric) Metric {
	if additional == nil {
		return Tasks{
			Val: metric.Value(),
		}
	}

	return Tasks{
		Val: metric.Value() + additional.Value(),
	}
}
