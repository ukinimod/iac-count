package model

import "github.com/MaibornWolff/iac-count/pkg/metrics"

type NodeData struct {
	Path     string
	NodeType string
	Metrics  map[string]metrics.Metric
	Children map[string]string
}
