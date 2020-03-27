package metrics

import (
	"github.com/MaibornWolff/iac-count/internal/util"
)

type FilesCalculator struct {
}

func (calculator FilesCalculator) Analyze(path, content string) Metric {
	return Files{
		Val: util.RecursiveFileCount(path),
	}
}
