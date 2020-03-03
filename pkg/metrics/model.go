package metrics

type NodeData struct {
	Path     string
	NodeType string
	Metrics  map[string]Metric
	Children map[string]string
}
