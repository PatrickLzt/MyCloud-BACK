package main

import (
	"database/sql"
	"fmt"
	"log"
	"mycloud-back/v1/models"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	var err error
	// global variable.
	models.DB, err = sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
