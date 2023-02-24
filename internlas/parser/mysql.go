package parser

import (
	"database/sql"
	"strconv"
)

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

// TODO: Terminar documentacion (Hace falta releer)
// parseToMYSQLType takes an interface{} (any) object which is underlying
// a *[]byte object, and parses it depending on the MYSQL type ("INT, VARCHAR, etc...")
// passed as parameter in string format.
// In case the function doesn't recognize the MYSQL Type, it parses it as string
func parseToMYSQLType(value any, mysqlType string) any {
	pointer := value.(*any)
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

// TODO: Terminar documentacion (Hace falta releer)
// parseValuesToMYSQLType takes a slice of interface{} (any) objects wihch are underlying to *byte[]
// objects but representing a string with another type ("12.32","2","realString"), and parse them
// all to his respsective type acording to the columnTypes passed as arguments.
//
// This means, having an interface{} slice []any{a,b,c} where a,b and c are *[]byte objects that represents
// each one to a different data type (for example 'a' is byte's slice pointer that represents a string, 'b' is a
// byte's slice pointer which parsed to string represents a float ("12.32") and 'b' is a byte's slice pointer
// and 'c' is a byte's slice which parsed to string represents a bool ("1")) and given columnTypes like
// ["VARCHAR", "DECIMAL", "TINYINT"]
func parseValuesToMYSQLType(valuesToParse []any, columnTypes []*sql.ColumnType) []any {
	var parsedValues []any
	for i, val := range valuesToParse {
		mysqlType := columnTypes[i].DatabaseTypeName()
		parsedVal := parseToMYSQLType(val, mysqlType)
		parsedValues = append(parsedValues, parsedVal)
	}
	return parsedValues
}

// ParseRowsToSlice takes a *sql.Rows object and converts it into a slice of maps.
// Every element of the array is underlying to his correct type
func ParseRowsToSlice(rows *sql.Rows) ([]map[string]any, error) {
	var response []map[string]any
	columns, _ := rows.Columns()
	values := getPointersSlice(len(columns))

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		columnTypes, _ := rows.ColumnTypes()

		response = append(response, parseSlicesToMap(columns, parsedValues))
	}
	return response, nil
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
