package store

import ( 
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Exec() {
	db, err := sql.Open("sqlite3", "./litter.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("select *")
	defer db.Close()
}
