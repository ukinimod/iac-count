package metrics

type Metric interface {
	Name() string
	Value() int
	Description() string
	add(metric Metric) Metric
}

type MetricCalculator interface {
	Analyze(path, content string) Metric
}

func AggregateMetrics(origin, additional *(map[string]Metric)) { // nolint:gocritic
	for key, metric := range *origin {
		if addMetric, exists := (*additional)[key]; exists {
			(*origin)[key] = metric.add(addMetric)
		}
	}

	for key, addMetric := range *additional {
		if _, exists := (*origin)[key]; !exists {
			(*origin)[key] = addMetric.add(nil)
		}
	}
}
