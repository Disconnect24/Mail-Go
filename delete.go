package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// Delete handles delete requests of mail.
func Delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM `mails` WHERE `sent` = 1 AND `recipient_id` = ? ORDER BY `timestamp` ASC LIMIT ?")
	if err != nil {
		// Welp, that went downhill fast.
		w.Write([]byte(GenNormalErrorCode(440, "Database error.")))
		log.Fatal(err)
	}

	isVerified, err := Auth(r.Form)
	if err != nil {
		fmt.Fprintf(w, GenNormalErrorCode(666, "Something weird happened."))
		log.Printf("Error deleting: %v", err)
		return
	} else if !isVerified {
		fmt.Fprintf(w, GenNormalErrorCode(240, "An authentication error occurred."))
		return
	}
	log.Println(isVerified, err)

	wiiID := r.Form.Get("mlid")
	delnum := r.Form.Get("delnum")
	_, err = stmt.Exec(wiiID, delnum)

	if err != nil {
		log.Fatal(err)
		w.Write([]byte(fmt.Sprint(GenNormalErrorCode(541, "Issue deleting mail from the database."))))
	} else {
		w.Write([]byte(fmt.Sprint(GenNormalErrorCode(100, "Success."),
			"deletenum=", delnum)))
	}
}
