package metrics

type Complexity struct {
	Val int
}

func (metric Complexity) Description() string {
	return "Complexity of file"
}

func (metric Complexity) Name() string {
	return "mcc"
}

func (metric Complexity) Value() int {
	return metric.Val
}

func (metric Complexity) add(additional Metric) Metric {
	if additional == nil {
		return Complexity{
			Val: metric.Value(),
		}
	}

	return Complexity{
		Val: metric.Value() + additional.Value(),
	}
}
