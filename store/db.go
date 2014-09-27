package store

import (
	"database/sql"
	//"database/sql/driver"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/awt/litter/config"
)

var Config *config.Config
type sqlFunc func(db *sql.DB, args ...interface{})

func Leet(body string) {
	withDB(func(db *sql.DB, args ...interface{}) {
		_, err := db.Exec("insert into leets VALUES (null, ?)", body)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func Leets() (leets []interface{}, err error) {
	// return leets with uids later than cut off datetime
	withDB(func(db *sql.DB, args ...interface{}) {
		rows, err := db.Query("select body from leets")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var body string
			rows.Scan(&body)
			leets = append(leets, body)
		}
		rows.Close()
	})
	return leets, err
}

func createTables(db *sql.DB) {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND 
	name='leets';`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		rows.Scan(&name)
	}
	defer rows.Close()
	
	// Create tables if they don't exist;
	if "" == name {
		log.Print("Initializing database...");
		statement := `CREATE TABLE leets (
			id INTEGER PRIMARY KEY,
			body TEXT NOT NULL
		)`

		_, err := db.Exec(statement);
		if err != nil {
			log.Fatal(err)
		}
	}
}

func withDB(f sqlFunc, args ...interface{}) {
	db, err := sql.Open("sqlite3", Config.Get("dbpath"))
	if err != nil {
		log.Fatal(err)
	}
	createTables(db)
	f(db, args...)
	defer db.Close()
}

func Reset() {
	log.Print("Resetting database..");
	withDB(func(db *sql.DB, args ...interface{}) {
		db.Exec("drop table leets");
	})
}
