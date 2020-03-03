package metrics

type CustomFacts struct {
	Val int
}

func (metric CustomFacts) Description() string {
	return "Number of custom facts"
}

func (metric CustomFacts) Name() string {
	return "custom_facts"
}

func (metric CustomFacts) Value() int {
	return metric.Val
}

func (metric CustomFacts) add(additional Metric) Metric {
	if additional == nil {
		return CustomFacts{
			Val: metric.Value(),
		}
	}

	return CustomFacts{
		Val: metric.Value() + additional.Value(),
	}
}
