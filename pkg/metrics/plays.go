package metrics

type Plays struct {
	Val int
}

func (metric Plays) Description() string {
	return "Number of plays"
}

func (metric Plays) Name() string {
	return "plays"
}

func (metric Plays) Value() int {
	return metric.Val
}

func (metric Plays) add(additional Metric) Metric {
	if additional == nil {
		return Plays{
			Val: metric.Value(),
		}
	}

	return Plays{
		Val: metric.Value() + additional.Value(),
	}
}
