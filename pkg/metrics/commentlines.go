package metrics

type Commentlines struct {
	Val int
}

func (metric Commentlines) Description() string {
	return "Number of comment-only lines in file"
}

func (metric Commentlines) Name() string {
	return "comment_lines"
}

func (metric Commentlines) Value() int {
	return metric.Val
}

func (metric Commentlines) add(additional Metric) Metric {
	if additional == nil {
		return Commentlines{
			Val: metric.Value(),
		}
	}

	return Commentlines{
		Val: metric.Value() + additional.Value(),
	}
}
