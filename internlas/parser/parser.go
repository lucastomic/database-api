package parser

import (
	"sort"
	"strings"
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

// SliceToString takes a slice and returns a string with his elements
// in the next string format:
// "key1,key2,...,keyN"
// For example:
//
// var values = []string{"name", "country", "age"}
// SliceToString(values) returns -> 'name,country,age'
func SliceToString(slice []string) string {
	var keysParsed string
	for _, key := range slice {
		keysParsed += key + ","
	}
	// We remove the last char here because this is an inconvenient comma (,)
	// For example: "name,age,country(,)"
	return removeLastChar(keysParsed)
}

// StringToSlice takes a string with the next froamt: "value1,value2,...,valueN" and converts
// it into a slice like this: []stirng{"value1", "value2",...,"valueN"}
// For example:
//
// str := "banana,apple,potato"
// StringToSlice(str) -> []string{"banana", "apple", "potato"}
func StringToSlice(str string) []string {
	return strings.Split(str, ",")
}
