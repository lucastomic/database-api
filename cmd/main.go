package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucastomic/syn-auth/internlas/database"
)

func HomeHandler(responseWriter http.ResponseWriter, r *http.Request) {
	db, err := database.GetMYSQL()
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, "No pudo conextarse")
	}
	res, err := db.Select("sale", []string{"date", "payedAtMoment"})
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, err)
	}
	fmt.Fprintln(responseWriter, string(res))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))

}
