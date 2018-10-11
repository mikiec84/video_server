package defs 

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

//response 
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	Username string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

//response 
type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type CommentsInfo struct {
	Comments []*Comment `json:"comments"`
}

type DeleteVideo struct {
	Success bool `json:"success"`
	Id string `json:"id"`
}

type NewCommentRes struct {
	Success bool `json:"success"`
	Id string 
}

/*
// Data Model
type VideoInfo struct {
	Id string 
	AuthorId int 
	Name string 
	DisplayCtime string 
}
*/

type Comment struct {
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	AuthorId string `json:"author"`
	Content string `json:"content"`
	ICon string `json:"icon"`
	Time string `json:"time"`
}

type SimpleSession struct {
	Username string 
	TTL int64
}

type User struct {
	Id int 
	LoginName string 
	Pwd string 
}



