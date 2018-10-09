package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO VIDEO_DEL_REC(VIDEO_ID) VALUES(?)")
	if err != nil {
		return err 
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord : %v", err)
		return err 
	}

	defer stmtIns.Close()

	return nil
}