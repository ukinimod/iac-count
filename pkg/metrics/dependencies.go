package metrics

type Dependencies struct {
	Val int
}

func (metric Dependencies) description() string {
	return "Number of dependencies"
}

func (metric Dependencies) name() string {
	return "dependencies"
}

func (metric Dependencies) value() int {
	return metric.Val
}
