package parser

// removeLastChar removes the last char from a string and returns the resultant string
func removeLastChar(s string) string {
	return s[:len(s)-1]
}

// ColumnsFromSlice takes a slice and returns a string with his elements
// in the next string format:
// "key1,key2,...,keyN"
// For example:
//
//	var values = []string{"name", "country", "age"}
//	ColumnsFromSlice(values) returns -> 'name,country,age'
func ColumnsFromSlice(slice []string) string {
	var keysParsed string
	for _, key := range slice {
		keysParsed += key + ","
	}
	// We remove the last char here because this is an inconvenient comma (,)
	// For example: "name,age,country(,)"
	return removeLastChar(keysParsed)
}

// MapValuesToQuestionMarks convert the values from a map to a string with as
// question marks as values in the map.
// For example, given:
//
//	var values = map[string]any{
//		"name":"Lucas",
//		"age":22,
//		"country":"Argentina",
//	}
//	MapValuesToQuestionMarks(values) would return -> '?,?,?' because "values" has 3 values
func MapValuesToQuestionMarks(mapToParse map[string]any) string {
	var questionMarks string
	for range mapToParse {
		questionMarks += "?,"
	}
	// We remove the last char here because this is an inconvenient comma (,)
	// For example: "?,?,?(,)"
	return removeLastChar(questionMarks)
}
