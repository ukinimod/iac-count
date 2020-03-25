package metrics

type Node struct {
	Path     string
	NodeType string
	Metrics  map[string]Metric
}
