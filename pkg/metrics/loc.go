package metrics

type Loc struct {
	Val int
}

func (metric Loc) Description() string {
	return "Number of code lines in file"
}

func (metric Loc) Name() string {
	return "loc"
}

func (metric Loc) Value() int {
	return metric.Val
}

func (metric Loc) add(additional Metric) Metric {
	if additional == nil {
		return Loc{
			Val: metric.Value(),
		}
	}

	return Loc{
		Val: metric.Value() + additional.Value(),
	}
}
