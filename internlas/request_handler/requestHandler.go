// requesthandler is the package for the handling of all the HTTP requests.
// His responsability is raise the server and handle the incoming requests
package requesthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucastomic/syn-auth/internlas/database"
)

// errorStatus is the status code returned when there is an error
var errorStatus int = 300

// writeError writes a response with the error status code and the message of the errors passed as argument
func writeError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(errorStatus)
	fmt.Fprintln(writer, err)
}

// getColumns returns the columns in the url query of the http request passed as argument.
// It returns them as a strings slice with the format []string{"column1,column2,...,columnN"}.
// For its purpose it could also be returned as []string{"column1", "column2",...,"columnN"} and it would work though
// If there is no column specified as query, it will return a slice with only one element: "*"
func getColumns(request *http.Request) []string {
	columns, ok := request.URL.Query()["columns"]
	if !ok {
		columns = []string{"*"}
	}
	return columns
}

// getRows retrive the rows from the db according to according to the information in the request.
// Only retrive the columns specified as argument.
// If no columns are specified as query, it retrive all the columns.
// If no where caluses are specified as query, it retrive all the rows of the table.
// TODO:Currently, getRowsHandler only support one where clauses at time.
func getRows(db database.MYSQLDB, request *http.Request, columns []string) ([]map[string]any, error) {
	var res []map[string]any
	var err error
	vars := mux.Vars(request)
	whereClauses, hasWhereClauses := request.URL.Query()["where"]
	if hasWhereClauses {
		res, err = db.SelectWhere(
			vars["table"],
			columns,
			whereClauses[0], /*Should support more than one where clauses here*/
		)
	} else {
		res, err = db.Select(vars["table"], columns)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

// getRowsHandler manage the select requests.
// It expects two queries, where clauses and the columns to retrive.
// It parses and writes de response as a JSON string.
// If there is any error (connecting to database or making the request) it writes
// it with the error status code
// This method supports this kind of requests:
// url/par/{table}[?columns=col1,col2,colN][&where=whereclause]
// TODO:Should support url/par/{table}[?columns=col1,col2,colN][&where=whereclause][&where=whereclause][&where=whereclause]...
func getRowsHandler(writer http.ResponseWriter, r *http.Request) {
	db, err := database.GetMYSQLDB()
	if err != nil {
		writeError(writer, err)
	}
	columns := getColumns(r)
	res, err := getRows(db, r, columns)

	if err != nil {
		writeError(writer, err)
	}

	parsedRes, err := json.Marshal(res)
	if err != nil {
		writeError(writer, err)
	}
	fmt.Fprintln(writer, string(parsedRes))
}

// pingHandler connects to the database and checks whether it works properly or not.
// If it doesn't it returns the error with a 300 status code.
// Otherwise, it returns a message exmplaning the datbase us currently active.
func pingHandler(writer http.ResponseWriter, r *http.Request) {
	db, err := database.GetMYSQLDB()

	if err != nil {
		writeError(writer, err)
	}
	err = db.Ping()

	if err != nil {
		writeError(writer, err)
	} else {
		fmt.Fprintln(writer, "Connected succesfully")
	}

}
func ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/get/{table}", getRowsHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
