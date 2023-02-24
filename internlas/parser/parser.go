package parser

import (
	"sort"
)

// MapValuesToSlice converts the given map as parameter to a slice with all his values.
// It is ordered with the order given in the slice passed as second parameter.
// For example, given:
//
//	var values = map[string]any{
//		"name":"Lucas",
//		"age":22,
//		"country":"Argentina",
//	}
//	order := []string{"age", "country", "name"}
//
// parseMapValuesToSluce(values, ) would return -> [22,"Argentina", "Lucas"]
func MapValuesToSlice(mapToParse map[string]any, order []string) []any {
	var slice []any
	for i := range order {
		slice = append(slice, mapToParse[order[i]])
	}
	return slice
}

// MapKeysToSlice returns an slice from the keys of a map, ordered by alphabet.
// For example, given:
//
//	var values = map[string]any{
//		"name":"Lucas",
//		"age":22,
//		"country":"Argentina",
//	}
//
// MapKeysToSlice(values) would return:
// []string{"age", "country", "name"}
func MapKeysToSlice(mapToParse map[string]any) []string {
	var keysSlice []string
	for key := range mapToParse {
		keysSlice = append(keysSlice, key)
	}
	sort.Strings(keysSlice)
	return keysSlice
}

// parseSlicesToMap converts two slices into a map, usnig the elements of the
// first slice as keys and the elements of the second slice as values
// For example, given:
//
// keys := []string{"name", "age", "country"}
// values := []stirng{"Lucas", 22, "Argentina"}
// parseSlicesToMap(keys, values) would return:
//
//	map[stirng]any{
//		"name":"Lucas",
//		"age":22,
//		"country": "Lucas"
//	}
func parseSlicesToMap(keys []string, values []any) map[string]any {
	var response map[string]any = make(map[string]any)
	for i := range keys {
		response[keys[i]] = values[i]
	}
	return response
}
