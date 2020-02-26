package reader

import (
	"log"

	"gopkg.in/yaml.v2"
)

// Returns list if a yaml string is provided or an empty string otherwise
func ReadYamlAsList(data string) []interface{} {
	var m []interface{}

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}

// Returns a map if a yaml string is provided or an empty map otherwise
func ReadYamlAsMap(data string) map[string]interface{} {
	m := make(map[string]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return m
}
