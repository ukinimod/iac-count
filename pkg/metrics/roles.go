package metrics

type Roles struct {
	val int
}

func (metric Roles) description() string {
	return "Number of roles"
}

func (metric Roles) name() string {
	return "roles"
}

func (metric Roles) value() int {
	return metric.val
}
