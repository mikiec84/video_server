package dbops

import (
	"strconv"
	"sync"
	"log"
	"database/sql"
	"go_dev/src/video_server/api/defs"
)



func InsertSession(sid string, ttl int64, uname string) error {

	ttlstr := strconv.FormatInt(ttl, 10)

	stmtIns, err := dbConn.Prepare(`INSERT INTO SESSIONS(SESSION_ID, TTL, LOGIN_NAME) VALUES(?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err 
	}
	defer stmtIns.Close()

	return nil
}

func DeleteSession(sid string) error {

	stmtDel, err := dbConn.Prepare(`DELETE FROM SESSIONS WHERE SESSION_ID=?`)
	if err != nil {
		return err 
	}

	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err 
	}

	defer stmtDel.Close()
	
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}

	stmtOut, err := dbConn.Prepare(`SELECT TTL, LOGIN_NAME FROM SESSIONS WHERE SESSION_ID=?`)
	if err != nil {
		return nil, err 
	}

	var ttl, uname string 
	stmtOut.QueryRow().Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err 
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		return nil, err 
	} else {
		ss.TTL = res
		ss.Username = uname 
	}
	defer stmtOut.Close()

	return ss, nil 
}

func RetrieveAllSessions() (*sync.Map, error) {

	m := &sync.Map{}

	stmtOut, err := dbConn.Prepare(`SELECT SESSION_ID, TTL, LOGIN_NAME FROM SESSIONS`)
	if err != nil {
		return nil, err 
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err 
	}

	for rows.Next() {
		var id string 
		var ttlstr string  
		var uname string 

		if err := rows.Scan(&id, &ttlstr, &uname); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break 
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 != nil {
			ss := &defs.SimpleSession{Username: uname, TTL: ttl}
			m.Store(id, ss)
			log.Printf(" session id : %s, ttl : %d", id, ss.TTL)
		}

	}

	defer stmtOut.Close()

	return m, nil 

}

