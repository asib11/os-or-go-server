package handlers

import (
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("this is handler :- middle a print hobo")
}
