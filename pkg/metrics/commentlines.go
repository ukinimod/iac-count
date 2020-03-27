package metrics

type Commentlines struct {
	Val int
}

func (metric CommentLines) Description() string {
	return "Number of comment-only lines in file"
}

func (metric CommentLines) Name() string {
	return "comment_lines"
}

func (metric CommentLines) Value() int {
	return metric.Val
}

func (metric CommentLines) add(additional Metric) Metric {
	if additional == nil {
		return CommentLines{
			Val: metric.Value(),
		}
	}

	return CommentLines{
		Val: metric.Value() + additional.Value(),
	}
}
