package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Scheme   string `json:"scheme"`
			Opaque   string `json:"opaque"`
			User     string `json:"user"`
			Host     string `json:"host"`
			Path     string `json:"path"`
			RawQuery string `json:"raw_query"`
			Fragment string `json:"fragment"`
		}{
			Scheme:   r.URL.Scheme,
			Opaque:   r.URL.Opaque,
			User:     r.URL.User.String(),
			Host:     r.URL.Host,
			Path:     r.URL.Path,
			RawQuery: r.URL.RawQuery,
			Fragment: r.URL.Fragment,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
