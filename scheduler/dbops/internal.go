package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("SELECT VIDEO_ID FROM VIDEO_DEL_REC LIMIT ?")
	var ids []string 
	if err != nil {
		return ids, err 
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error : %v", err)
		return ids, err 
	}

	for rows.Next() {
		var id string 
		if err := rows.Scan(&id); err != nil {
			return ids, err 
		}
		ids = append(ids, id)
	}

	defer stmtOut.Close()

	return ids, nil 
}

func DeleteVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM VIDEO_DEL_REC WHERE VIDEO_ID=?")
	if err != nil {
		return err 
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err 
	}

	defer stmtDel.Close()

	return nil 
}