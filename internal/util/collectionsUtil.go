package util

import (
	"sort"
)

func Contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func SortedKeys(aMap map[string]interface{}) []string {
	var keys []string
	for aKey := range aMap {
		keys = append(keys, aKey)
	}
	sort.Strings(keys)

	return keys
}
