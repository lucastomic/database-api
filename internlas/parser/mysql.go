package parser

import (
	"database/sql"
	"strconv"
)

// removeLastChar removes the last char from a string and returns the resultant string
func removeLastChar(s string) string {
	return s[:len(s)-1]
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

// parseFromMYSQLType takes an interface{} (any) object which is underlying
// a *[]byte object, and parses it depending on the MYSQL type ("INT, VARCHAR, etc...")
// passed as parameter in string format.
// In case the function doesn't recognize the MYSQL Type, it parses it as string
func parseFromMYSQLType(value any, mysqlType string) any {

	pointer := value.(*any)
	if *pointer == nil {
		return nil
	}
	byteSlice := (*pointer).([]byte)
	strValue := string(byteSlice)
	var response any

	switch mysqlType {
	case "INT":
		response, _ = strconv.Atoi(strValue)
	case "TINYINT":
		response = strValue == "1"
	case "DECIMAL":
		response, _ = strconv.ParseFloat(strValue, 2)
	default:
		response = strValue
	}
	return response
}

// parseValuesFromMYSQLType takes a slice of interface{} (any) objects wihch are underlying to *byte[]
// objects but representing a string with another type ("12.32","2","realString"), and parse them
// all to his respsective type acording to the columnTypes passed as arguments.
//
// This means, having an interface{} slice []any{a,b,c} where a,b and c are *[]byte objects that represents
// each one to a different data type (for example 'a' is byte's slice pointer that represents a string, 'b' is a
// byte's slice pointer which parsed to string represents a float ("12.32")
// and 'c' is a byte's slice which parsed to string represents a bool ("1")) and given columnTypes like
// ["VARCHAR", "DECIMAL", "TINYINT"] parseValuesFromMYSQLType(ourSlice, ourMYSQLTypes) would return
// a slice with the next types: []any{string,float,bool}
func parseValuesFromMYSQLType(valuesToParse []any, columnTypes []*sql.ColumnType) []any {
	var parsedValues []any
	for i, val := range valuesToParse {
		mysqlType := columnTypes[i].DatabaseTypeName()
		parsedVal := parseFromMYSQLType(val, mysqlType)
		parsedValues = append(parsedValues, parsedVal)
	}
	return parsedValues
}

// getMapFromRow converts the current row from the rows passed as argument (with "the current row"
// we refer to the row which will be retrived from the rows iterator) to a map.
// If there is an error it retruns it as second value.
func getMapFromRow(rows *sql.Rows) (map[string]any, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := getPointersSlice(len(columns))
	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}
	columnTypes, _ := rows.ColumnTypes()
	parsedValues := parseValuesFromMYSQLType(values, columnTypes)
	return parseSlicesToMap(columns, parsedValues), nil
}

// ParseRowsToMapSlice takes a *sql.Rows object and converts it into a slice of maps,
// where each row is a map element in the slice.
// Every element of the map is underlying to his correct type
// If there is an error retriving one of the lines, instead of this line (parsed to map)
// it adds a map explaing the error.
func ParseRowsToMapSlice(rows *sql.Rows) ([]map[string]any, error) {
	if rows == nil {
		return []map[string]any{}, nil
	}
	var response []map[string]any
	for rows.Next() {
		if newMap, err := getMapFromRow(rows); err != nil {
			errorMessage := getErrorMessageAsMap(err)
			response = append(response, errorMessage)
		} else {
			response = append(response, newMap)
		}
	}
	return response, nil
}

// getErrorMessageAsMap returns the error specified as arguemnt as a map with the next format:
//
//	{
//		"message": "error retriving this row",
//		"error":   err,
//	}
func getErrorMessageAsMap(err error) map[string]any {
	return map[string]any{
		"message": "error retriving this row",
		"error":   err,
	}
}

// getPointersSlice returns an already initialized slice of pointers ([]any whose all elements are pointers)
// with as many poiners as the integer passed as argument
func getPointersSlice(length int) []any {
	values := make([]any, length)
	for i := range values {
		var v any
		values[i] = &v
	}
	return values
}
