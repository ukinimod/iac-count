package metrics

type Rloc struct {
	Val int
}

func (metric Rloc) Description() string {
	return "Number of non-blank, non-comment lines in file"
}

func (metric Rloc) Name() string {
	return "rloc"
}

func (metric Rloc) Value() int {
	return metric.Val
}

func (metric Rloc) add(additional Metric) Metric {
	if additional == nil {
		return Rloc{
			Val: metric.Value(),
		}
	}

	return Rloc{
		Val: metric.Value() + additional.Value(),
	}
}
