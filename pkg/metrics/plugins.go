package metrics

type Plugins struct {
	Val int
}

func (metric Plugins) Description() string {
	return "Number of plugins"
}

func (metric Plugins) Name() string {
	return "plugins"
}

func (metric Plugins) Value() int {
	return metric.Val
}

func (metric Plugins) add(additional Metric) Metric {
	if additional == nil {
		return Plugins{
			Val: metric.Value(),
		}
	}

	return Plugins{
		Val: metric.Value() + additional.Value(),
	}
}
