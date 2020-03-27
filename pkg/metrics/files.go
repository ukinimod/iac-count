package metrics

type Files struct {
	Val int
}

func (metric Files) Description() string {
	return "Number of files"
}

func (metric Files) Name() string {
	return "files"
}

func (metric Files) Value() int {
	return metric.Val
}

func (metric Files) add(additional Metric) Metric {
	if additional == nil {
		return Files{
			Val: metric.Value(),
		}
	}

	return Files{
		Val: metric.Value(),
	}
}
