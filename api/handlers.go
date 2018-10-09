package main 


import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go_dev/src/video_server/api/defs"
	"go_dev/src/video_server/api/dbops"
	"go_dev/src/video_server/api/session"
	"go_dev/src/video_server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	fmt.Println(string(res))
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return 
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}


func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	fmt.Println(string(res))

	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	uname := p.ByName("user_name")
	if uname != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return 
	}

	log.Printf("%s", ubody.Username)
	pwd, err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return 
	}

	id := session.GenerateNewSessionId(ubody.Username)
	log.Printf("SessionId %s", id)
	si := &defs.SignedIn{Success: true, SessionId: id}

	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !ValidateUser(w, r) {
		log.Printf("Unathorized user\n")
		return 
	}

	var idx int 
	uname := p.ByName("user_name")
	idx, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	}

	//log.Println("GetUserInfo idx ", idx)
	ui := &defs.UserInfo{Id: idx}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

	//io.WriteString(w, uname)
}


func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
/*
 {
	"author_id": 1,
	"name":"123456.mp4"
 }
 */
 
	if !ValidateUser(w, r) {
		log.Printf("Unathorized User\n")
		return 
	} 

	uname := p.ByName("user_name")
	idx, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return 
	}
	log.Println("username: ", uname, "userid" , idx)

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}

	log.Println(string(res))

	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	vi, err := dbops.AddVideoInfo(idx, nvbody.Name)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
	}

	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}


func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Printf("ListAllVideos")

	if !ValidateUser(w, r) {
		log.Printf("Unathorized User\n")
		return 
	}
	

	uname := p.ByName("user_name")
	idx, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return 
	}
	log.Println("username: ", uname, "userid" , idx)

	vis, err := dbops.ListVideoInfo(idx)
	if err != nil {
		log.Printf("Error in ListVideoInfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	}

	videos := &defs.VideosInfo{Videos: vis}
	//log.Printf("videos ")
	if resp, err := json.Marshal(videos); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

/*
 
 */
func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	uname := p.ByName("user_name")
	vid := p.ByName("vid-id")
	//io.WriteString(w, uname)

	idx, err := dbops.GetUser(uname)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return 
	}

	err = dbops.DeleteVideoInfo(vid, idx)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	} 
	
	videos := defs.DeleteVideo{Success: true, Id: vid}
	if resp, err := json.Marshal(videos); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}
//IAM SSO Rbac,
//SOA
/*
{
	"author_id": 1,
	"content": "你好"
} */

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !ValidateUser(w, r) {
		log.Printf("Unathorized User\n")
		return 
	}
	log.Println("Enter PostComment")

	vid := p.ByName("vid-id")
	log.Println("PostComment", vid)

	res, _ := ioutil.ReadAll(r.Body)
	nvcomment := &defs.NewComment{}

	if err := json.Unmarshal(res, nvcomment); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	log.Println("vid", vid, "aid", nvcomment.AuthorId, "content", nvcomment.Content)

	cid, err := dbops.AddComments(vid, nvcomment.AuthorId, nvcomment.Content)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	}

	log.Println("NewCommentRes cid: ", cid)
	ncr := defs.NewCommentRes{Success: true, Id: cid}
	if resp, err := json.Marshal(ncr); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}


func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !ValidateUser(w, r) {
		log.Printf("Unathorized User\n")
		return 
	}

	vid := p.ByName("vid-id")
	cm, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec()) 
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return 
	}
/*
	log.Printf("ListComment %d\n", len(cm))
	for i := 0;i < len(cm);i ++ {
		log.Println(cm[i].Id, cm[i].AuthorId, cm[i].VideoId, cm[i].Content)
	}*/

	cms := &defs.CommentsInfo{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}
