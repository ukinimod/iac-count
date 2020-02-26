package reader

import (
	"io/ioutil"
	"log"
)

// ReadFileToString returns content of a file as string if it exists
// or empty string if it doesn't exist
func ReadFileToString(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return string(content)
}
