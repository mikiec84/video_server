package main

import (
	"net/http"
	"io"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}