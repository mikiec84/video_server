package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go_dev/src/video_server/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r 
	return m
}

func (m middleWareHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session 
	validateUserSession(r)
	m.r.ServeHTTP(w, r)

}
//192.168.189.134:8080/comments/videos/02894ee6-a92b-b5fa-d814-6bc6669207da
func RegisterHandlers() *httprouter.Router {

	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:user_name", GetUserInfo)
	
	router.POST("/videos/:user_name", AddNewVideo)
	router.GET("/videos/:user_name", ListAllVideos)
	router.DELETE("/videos/:user_name/:vid-id", DeleteVideo)

	router.POST("/comments/videos/:vid-id", PostComment)
	router.GET("/comments/videos/:vid-id", ShowComments)

	return router
}

func Perpare() {
	session.LoadSessionFromDB()
}

func main() {
	Perpare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}


//handler -> validation{1.request, 2.user} -> 
//1. data model
//2. error handling