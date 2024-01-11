package handlers

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/tutorials/go/crud/pkg/mocks"
	"github.com/tutorials/go/crud/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request){
	// read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// append to the books mock
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// send a status 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Successfully created book")
}