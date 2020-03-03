package metrics

type Comments struct {
	Val int
}

func (metric Comments) Description() string {
	return "Number of comment-only lines in file"
}

func (metric Comments) Name() string {
	return "comment_lines"
}

func (metric Comments) Value() int {
	return metric.Val
}

func (metric Comments) add(additional Metric) Metric {
	if additional == nil {
		return Comments{
			Val: metric.Value(),
		}
	}

	return Comments{
		Val: metric.Value() + additional.Value(),
	}
}
