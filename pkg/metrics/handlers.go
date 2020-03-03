package metrics

type Handlers struct {
	Val int
}

func (metric Handlers) Description() string {
	return "Number of handlers"
}

func (metric Handlers) Name() string {
	return "handlers"
}

func (metric Handlers) Value() int {
	return metric.Val
}

func (metric Handlers) add(additional Metric) Metric {
	if additional == nil {
		return Handlers{
			Val: metric.Value(),
		}
	}

	return Handlers{
		Val: metric.Value() + additional.Value(),
	}
}
