package ansible

import (
	"github.com/MaibornWolff/iac-count/internal/util"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type FilesCalculator struct {
}

func (calculator FilesCalculator) Analyze(path, content string) metrics.Metric {
	return metrics.Files{
		Val: util.RecursiveFileCount(path),
	}
}
