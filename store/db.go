package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"io/ioutil"
	"github.com/awt/litter/config"
	"path/filepath"
	"os"
)

type Leet struct {
	Body string
	Signature string
	From string
}

type LeetList struct {
	Collection []Leet
}

type Name struct {
	Name string
	ShortHash string
	BlockCount int64
}

var Config *config.Config
type sqlFunc func(db *sql.DB, args ...interface{})

func AddPendingName(name string, shortHash string, blockCount int64) {
	// Insert new name into the db
	withDB(func(db *sql.DB, args ...interface{}) {
		_, err := db.Exec("insert into names VALUES (null, ?, ?, ?, 'pending')", name, shortHash, blockCount)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func CreateLeet(body string) {
	withDB(func(db *sql.DB, args ...interface{}) {
		_, err := db.Exec("insert into leets VALUES (null, ?)", body)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func Follow(name string) {
	withDB(func(db *sql.DB, args ...interface{}) {
		_, err := db.Exec("insert into friends VALUES (null, ?)", name)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func Friends() (friends []interface{}, err error) {
	withDB(func(db *sql.DB, args ...interface{}) {
		rows, err := db.Query("select name from friends")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var name string
			rows.Scan(&name)
			friends = append(friends, name)
		}
		rows.Close()
	})
	return friends, err
}

func ImportLeet(leet map[string]interface{}) {
	CreateLeet(leet["body"].(string))
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

func MarkRegistered(name string) {
	withDB(func(db *sql.DB, args ...interface{}) {
		_, err := db.Exec("update names set state='registered' where name=?", name)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func PendingNames() (names []Name, err error) {

	withDB(func(db *sql.DB, args ...interface{}) {
		rows, err := db.Query("select name, short_hash, block_count from names where state='pending'")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var name Name
			rows.Scan(&name.Name, &name.ShortHash, &name.BlockCount)
			names = append(names, name)
		}
		rows.Close()
	})
	return names, err
}

func createTables(db *sql.DB) {
	createTable(db, "leets", `body TEXT NOT NULL`)
	createTable(db, "friends", `name TEXT NOT NULL`)
	createTable(db, "names", `name TEXT NOT NULL UNIQUE, short_hash TEXT NOT NULL, block_count INTEGER, state TEXT DEFAULT 'pending'`)
}

func createTable(db *sql.DB, name string, fields string) {
	query := fmt.Sprintf(`SELECT name FROM sqlite_master WHERE type='table' AND name='%s';`, name)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var nameField string
	for rows.Next() {
		rows.Scan(&nameField)
	}
	
	// Create tables if they don't exist;
	if "" == nameField {
		log.Printf("Creating %s table", name);
		statement := fmt.Sprintf("CREATE TABLE %s (id INTEGER PRIMARY KEY, %s)", name, fields)
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
		db.Exec("drop table IF EXISTS leets");
		db.Exec("drop table IF EXISTS friends");
	})
}

func LoadFixture(path string) []byte {

	// Read fixture file

	fullPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "awt", "litter", path)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Printf("Couldn't find json fixture at path: %s. Exiting.", fullPath);
		os.Exit(1)
	}

	return content
}
