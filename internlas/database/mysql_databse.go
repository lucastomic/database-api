package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucastomic/syn-auth/internlas/parser"
)

// Relevant information to connect to the databse
var user = "root"
var pass = "secret"
var host = "tcp(mysql)"
var databaseName = "naturalYSalvaje"

// MYSQLDB is the MYSQL implementation of the database interface
// Allows to make requests to a MYSQL database
type MYSQLDB struct {
	db *sql.DB
}

// GetMYSQLDB returns an instance of MYSQLDB
// In case of error it returns an empty MYSQLDB and the eror as second parameter
// It connects automaticaly to a DB with the information in the file
func GetMYSQLDB() (MYSQLDB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s", user, pass, host, databaseName),
	)
	if err != nil {
		return MYSQLDB{}, err
	}
	return GetMYSQLDBWithDB(db)
}

// GetMYSQLDBWithDB connects to a mysql db specified as argument.
// This method is made for testing and is not recommended in production code
// unless the need of using another DB more than the main one
func GetMYSQLDBWithDB(db *sql.DB) (MYSQLDB, error) {
	return MYSQLDB{db}, nil
}

// Ping returns checks whether the databse is still alive.
// If it isn't it returns the error. If it is, it returns nil
func (mysql MYSQLDB) Ping() error {
	return mysql.Ping()
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

// Select brings the values from the table and columns specified as argumen in json format
// For example,
// mysql.Select("animal", []string{"legs", "mammal", "name"})
// Could return something like this:
// [
//
//		{
//			name:"Dog"
//			legs: 4,
//			mammal: true
//		},
//	 	{
//	 	name:"Snake",
//			legs: 0,
//			mammal: false
//		},
//
// ]
func (mysql MYSQLDB) Select(table string, columns []string) (string, error) {
	var response []map[string]any

	rows, err := mysql.getRows(table, columns)
	if err != nil {
		return "", err
	}
	response, err = parser.ParseRowsToMapSlice(rows)
	if err != nil {
		return "", err
	}
	parsedResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(parsedResponse), nil
}

// getRows retrive the rows from the table with the rows specified as argument
// For example:
// mysql.getRows("animal", []string{"legs", "fur"}) returns
// the rows with the columns "legs" and "fur" from the table "animal"
func (mysql MYSQLDB) getRows(table string, columns []string) (*sql.Rows, error) {
	parsedColumns := parser.ColumnsFromSlice(columns)
	query := fmt.Sprintf("SELECT %v FROM %v.%v", parsedColumns, databaseName, table)
	rows, err := mysql.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
