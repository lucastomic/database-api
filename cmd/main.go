package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucastomic/syn-auth/internlas/database"
)

func HomeHandler(responseWriter http.ResponseWriter, r *http.Request) {
	db, err := database.GetMYSQLDB()
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, "No pudo conextarse")
	}
	res, err := db.Select("sale", []string{"date", "payedAtMoment"})
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, err)
	}
	parsedRes, err := json.Marshal(res)
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, err)
	}
	fmt.Fprintln(responseWriter, string(parsedRes))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))

}
