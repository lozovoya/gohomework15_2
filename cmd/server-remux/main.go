package main

import (
	"encoding/json"
	"log"
	"net/http"

	"netology/webinar-rest/pkg/core"
	"netology/webinar-rest/pkg/remux"
)

func main() {
	rmux := remux.New()
	rmux.RegisterPlain(remux.GET, "/bands", http.HandlerFunc(bandsHandler))

	log.Fatal(http.ListenAndServe(":8000", rmux))
}

func bandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.ErrNotSupported.Error(), http.StatusBadRequest)
		return
	}
	bands := core.Bands()
	data, err := json.Marshal(bands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
