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
	"github.com/lucastomic/syn-auth/internlas/injector"
	"github.com/lucastomic/syn-auth/internlas/parser"
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
func getRows(db database.Database, request *http.Request, columns []string) ([]map[string]any, error) {
	var res []map[string]any
	var err error
	vars := mux.Vars(request)
	whereClauses, hasWhereClauses := request.URL.Query()["where"]
	if hasWhereClauses {
		whereClausesParsed := parser.ParseWhereClauses(whereClauses)
		res, err = db.SelectWhere(
			vars["table"],
			columns,
			whereClausesParsed,
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
// supports url/par/{table}[?columns=col1,col2,colN][&where=whereclause][&where=whereclause][&where=whereclause]...
func getRowsHandler(writer http.ResponseWriter, r *http.Request) {
	db, err := injector.GetDatabase()
	if err != nil {
		writeError(writer, err)
		return
	}
	columns := getColumns(r)
	res, err := getRows(db, r, columns)

	if err != nil {
		writeError(writer, err)
		return
	}

	parsedRes, err := json.Marshal(res)
	if err != nil {
		writeError(writer, err)
		return
	}
	writer.WriteHeader(200)
	fmt.Fprintln(writer, string(parsedRes))
}

// pingHandler connects to the database and checks whether it works properly or not.
// If it doesn't it returns the error with a 300 status code.
// Otherwise, it returns a message exmplaning the datbase us currently active.
func pingHandler(writer http.ResponseWriter, r *http.Request) {
	db, err := injector.GetDatabase()

	if err != nil {
		writeError(writer, err)
		return
	}
	err = db.Ping()

	if err != nil {
		writeError(writer, err)
		return
	} else {
		writer.WriteHeader(200)
		fmt.Fprintln(writer, "Connected succesfully")
	}

}

// insertIntoHandler manage the POST requests made for insert into operations.
// Expects the table as a parameter in the URL and the columns and values in the request's body
// Currently, only supports one row at time TODO: support more than one row at time
// If there is an error it returns it with a error status code,
// if there isn't returns a success message with a 200 status code
func insertIntoHandler(writer http.ResponseWriter, r *http.Request) {
	db, err := injector.GetDatabase()
	table := mux.Vars(r)["table"]
	if err != nil {
		writeError(writer, err)
		return
	}
	body, err := parser.ReqBodyToMap(r)
	if err != nil {
		writeError(writer, err)
		return
	}
	err = db.InsertInto(table, body)
	if err != nil {
		writeError(writer, err)
		return
	}
	writer.WriteHeader(200)
	fmt.Fprintln(writer, "Inserted succesfully")
}

// ListenAndServe raises the API server and configure all handlers.
func ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/get/{table}", getRowsHandler)
	r.HandleFunc("/insert/{table}", insertIntoHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
