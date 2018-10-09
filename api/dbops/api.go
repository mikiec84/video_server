package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"go_dev/src/video_server/api/defs"
	"go_dev/src/video_server/api/utils"
	"time"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:zhaomeiping@tcp(192.168.189.155:3306)/VIDEO_DB?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return dbConn
}

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO USERS (LOGIN_NAME, PWD) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err 
	}

	defer stmtIns.Close()

	return nil
}

func GetUserCredential(loginName string) (string, error) {

	stmtOut, err := dbConn.Prepare("SELECT PWD FROM USERS WHERE LOGIN_NAME=?");
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string 
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", nil
	}
	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM USERS WHERE LOGIN_NAME=? AND PWD=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err 
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return nil 
	}

	defer stmtDel.Close();

	return nil
}

func GetUser(loginName string) (int, error) {

	stmtOut, err := dbConn.Prepare("SELECT ID FROM USERS WHERE LOGIN_NAME=?");
	if err != nil {
		log.Printf("%s", err)
		return -1, err
	}

	var idx int
	err = stmtOut.QueryRow(loginName).Scan(&idx)
	if err != nil && err != sql.ErrNoRows {
		return idx, nil
	}
	defer stmtOut.Close()

	return idx, nil
}

func AddVideoInfo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err 
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmtIns, err := dbConn.Prepare(`INSERT INTO VIDEO_INFO(ID, AUTHOR_ID, NAME, DISPLAY_CTIME) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err 
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {

	stmtOut, err := dbConn.Prepare(`SELECT AUTHOR_ID, NAME, DISPLAY_CTIME FROM VIDEO_INFO WHERE ID=?`)
	if err != nil {
		return nil, err
	}

	var aid int 
	var name string 
	var dct string 
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err 
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	defer stmtOut.Close()

	return res, nil
}

func ListVideoInfo(aid int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT ID, NAME, DISPLAY_CTIME FROM VIDEO_INFO WHERE AUTHOR_ID= ?`)

	var res []*defs.VideoInfo
	rows, err := stmtOut.Query(aid)
	if err != nil {
		return res, nil 
	}

	for rows.Next() {
		var id, name, displayctime string
		if err := rows.Scan(&id, &name, &displayctime); err != nil {
			return res, err 
		}

		v := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: displayctime}
		res = append(res, v)
	}

	defer stmtOut.Close()

	return res, err

} 

func DeleteVideoInfo(vid string, aid int) error {

	stmtDel, err := dbConn.Prepare(`DELETE FROM VIDEO_INFO WHERE ID=? AND AUTHOR_ID=?`)
	if err != nil {
		return nil 
	}

	_, err = stmtDel.Exec(vid, aid)
	if err != nil {
		return nil 
	}
	defer stmtDel.Close()

	return nil
}


func AddComments(vid string, aid int, content string) (string, error) {

	id, err := utils.NewUUID()
	if err != nil {
		return "", err
	}

	stmtIns, err := dbConn.Prepare(`INSERT INTO COMMENTS(ID, VIDEO_ID, AUTHOR_ID, CONTENT) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return "", err
	}

	log.Println(id, vid, aid, content)
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		log.Printf("%s", err)
		return "", err 
	}
	defer stmtIns.Close()

	return id, nil 
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT COMMENTS.ID, USERS.LOGIN_NAME, 
	COMMENTS.CONTENT FROM COMMENTS 
	INNER JOIN USERS ON COMMENTS.AUTHOR_ID = USERS.ID
	WHERE COMMENTS.VIDEO_ID=? AND COMMENTS.TIME > FROM_UNIXTIME(?) 
	AND COMMENTS.TIME <= NOW()`) //FROM_UNIXTIME(?)

	//log.Println("vid", vid, "from", from, "to", to)

	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from)//, to
	if err != nil {
		return res, nil 
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err 
		}

		c := &defs.Comment{Id: id, VideoId: vid, AuthorId: name, Content: content}
		res = append(res, c)
	}

	defer stmtOut.Close()

	return res, err

}
