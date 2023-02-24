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
	values := map[string]any{
		"date":          "2023-12-06",
		"payedAtMoment": true,
		"amount":        12,
		"userID":        1,
		"product":       "besugo",
		"caliberName":   "1-3kg",
	}
	err = db.InsertInto("sale", values)
	if err != nil {
		responseWriter.WriteHeader(300)
		fmt.Fprintln(responseWriter, err)
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))

}
