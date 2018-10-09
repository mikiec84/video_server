package dbops

import (
	"testing"
	_ "go_dev/src/video_server/api/defs"
	"strconv"
	"time"
	"fmt"
)

var (
	tempvid string
	tempaid int 
)

//init(dblogin, truncate tables) --> run tests --> clear  

func clearTables() {
	dbConn.Exec("truncate USERS")
	dbConn.Exec("truncate VIDEO_INFO")
	dbConn.Exec("truncate COMMENTS")
	dbConn.Exec("truncate SESSIONS")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("wangbojing", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
	tempaid = 1
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("wangbojing")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("wangbojing", "123")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("wangbojing")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting of User: %v", err)
	}
}


func TestVideoWorkFlow(t *testing.T) {
	t.Run("AddUser", testAddUser)
	t.Run("AddVideoInfo", testAddVideoInfo)
	t.Run("GetVideoInfo", testGetVideoInfo)
	t.Run("DelVideoInfo", testDelVideoInfo)
	t.Run("RegetVideoInfo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddVideoInfo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err) 
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDelVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid, tempaid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if !(err == nil && vi == nil) {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	cid, err := AddComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
	if len(cid) == 0 {
		t.Errorf("Error of CID: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d %v \n", i, ele)
	}
}

