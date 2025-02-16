package main

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	switch b.Method { //
	case  http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("httpClient : %v\n", err);
			return 
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("httpClient : %v\n", err);
			return 
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("httpClient : %v\n", err);
			return 
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		w.WriteHeader(500)
		io.WriteString(w, string(re))

		return 
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
