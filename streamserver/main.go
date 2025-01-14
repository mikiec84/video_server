package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r 
	m.l = NewConnLimiter(cc)

	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too Many Request")
		return 
	}
	t := time.Now()
	m.r.ServeHTTP(w, r)
	d := time.Now().Sub(t)
	fmt.Println("http server during time -->: ", d.String())

	defer m.l.ReleaseConn()
	
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New() 

	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	router.GET("/testpage", testPageHandler)

	return router 
}

func main() {

	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)

}