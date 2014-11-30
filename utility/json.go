package utility

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ReadJson(r *http.Request, data interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}
	return true
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	// prefix := []byte(")]}',\n")
	if d, err := json.Marshal(data); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		// d = append(prefix, d...)
		w.Header().Set("Content-Length", strconv.Itoa(len(d)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}
}
