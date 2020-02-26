package core

type FileData struct {
	RelativePath string
	Filetype     string
	Metrics      map[string]int
	Children     map[string]string
}
