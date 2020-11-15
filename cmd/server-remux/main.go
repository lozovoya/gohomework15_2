package main

import (
	"encoding/json"
	"log"
	"net/http"
	"remogithub.com/lozovoya/gohomework15_2/pkg/core"
	"remogithub.com/lozovoya/gohomework15_2/pkg/remux"
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
