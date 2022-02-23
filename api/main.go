package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TableContent struct {
}

func main() {
	fmt.Println("server is running")
	r := mux.NewRouter()
	// GET
	r.HandleFunc("/incoming", getIncomingPostback)
	log.Fatal(http.ListenAndServe(":8000", r))
}
func getIncomingPostback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var incoming = currentIncoming("incoming")
	json.NewEncoder(w).Encode(incoming)
}
func currentIncoming(reqName string) []TableContent {
	var Incoming []TableContent

	db, err := sql.Open("localhost", "3306")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM incoming_postback", reqName)

	for rows.Next() {
		i := TableContent{}

		err = rows.Scan()
		if err != nil {
			log.Fatal(err)
		}
		Incoming = append(Incoming, i)
	}
	return Incoming
}
