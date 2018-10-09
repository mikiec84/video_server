package main

import (
	"github.com/julienschmidt/httprouter"
	"go_dev/src/video_server/scheduler/dbops"
	_ "go_dev/src/video_server/scheduler/taskrunner"
	"net/http"
	"fmt"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	fmt.Println("vid: ", vid)
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")

	return
}
