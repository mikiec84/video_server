package main 

import (
	"log"
	"net/http"
	"go_dev/src/video_server/api/session"
	"go_dev/src/video_server/api/defs"
)

var HEADER_FILED_SESSION = "X-Session-Id" 
var HEADER_FILED_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false 
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false 
	}

	r.Header.Add(HEADER_FILED_UNAME, uname)

	return true 
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {

	uname := r.Header.Get(HEADER_FILED_UNAME)
	log.Printf("ValidateUser uname: %s", uname)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return false
	}

	return true
}
