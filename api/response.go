package main

import (
	"io"
	"net/http"
	"encoding/json"
	"go_dev/src/video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}

