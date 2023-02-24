package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucastomic/syn-auth/internlas/parser"
)

// Relevant information to connect with the databse
var user = "root"
var pass = "secret"
var host = "tcp(mysql)"
var databaseName = "naturalYSalvaje"

// MYSQLDB is the MYSQL implementation of the databse interface
type MYSQLDB struct {
	db *sql.DB
}

// GetMYSQL returns an instance of MYSQLDB
// In case of error it returns an empty MYSQLDB and the eror as second parameter
func GetMYSQL() (MYSQLDB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s", user, pass, host, databaseName),
	)
	if err != nil {
		return MYSQLDB{}, err
	}
	return MYSQLDB{db}, nil
}

// InsertInto inserts a row in the table with values specified as arguments.
// If there is any error, it returns it. Otherwise it returns nil
// Before converting the body into the statement, we have to order it. Because,
// when we convert a map into a slice, the order changes every time, so the values
// wouldn't correspond to the columns if we didn't oreder them.
func (mysql MYSQLDB) InsertInto(table string, body map[string]any) error {
	orderedColumns := parser.MapKeysToSlice(body)
	columns := parser.ColumnsFromSlice(orderedColumns)
	questionMarks := parser.MapValuesToQuestionMarks(body)
	valuesSlice := parser.MapValuesToSlice(body, orderedColumns)

	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table, columns, questionMarks)
	statement, err := mysql.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(valuesSlice...)
	if err != nil {
		return err
	}
	defer statement.Close()
	return nil
}

func (mysql MYSQLDB) Select(table string, columns []string) ([]byte, error) {
	var response []map[string]any

	parsedColumns := parser.ColumnsFromSlice(columns)
	query := fmt.Sprintf("SELECT %v FROM %v.%v", parsedColumns, databaseName, table)
	rows, err := mysql.db.Query(query)
	if err != nil {
		return []byte{}, err
	}

	values := make([]any, len(columns))
	for i := range values {
		var v any
		values[i] = &v
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return []byte{}, err
		}

		response = append(response, parser.ParseSlicesToMap(columns, values))

	}
	parsedResponse, err := json.Marshal(response)
	return parsedResponse, nil
}
