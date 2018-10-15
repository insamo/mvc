package sliceHelper

import (
	"sort"
)

func DeDuplicatesUnordered(elements []string) []string {
	if elements == nil {
		return nil
	}
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}
