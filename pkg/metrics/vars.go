package metrics

type Vars struct {
	Val int
}

func (metric Vars) Description() string {
	return "Number of variables (defaults, host_vars, group_vars)"
}

func (metric Vars) Name() string {
	return "vars"
}

func (metric Vars) Value() int {
	return metric.Val
}

func (metric Vars) add(additional Metric) Metric {
	if additional == nil {
		return Vars{
			Val: metric.Value(),
		}
	}

	return Vars{
		Val: metric.Value() + additional.Value(),
	}
}
