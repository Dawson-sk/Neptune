package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marhsaling 'respondWithJson': ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("Access-Control-Allow-Origin, "*")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(code, msg, err)
	}
	
	respondWithJSON(w, code, map[string]string{"error": msg})
}
