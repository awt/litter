package store

import ( 
	"database/sql"
	//"database/sql/driver"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/awt/litter/config"
)

var Config *config.Config

func Leet(body string) {
	exec("insert into leets VALUES (null, ?)", body)
}

func exec(statement string, args ...interface{}) {
	log.Print(args);
	db, err := sql.Open("sqlite3", Config.Get("dbpath"))
	if err != nil {
		log.Fatal(err)
	}
	createTables(db)
	_, err = db.Exec(statement, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

//func withDB() {
	//db, err := sql.Open("sqlite3", Config.Get("dbpath"))
	//if err != nil {
		//log.Fatal(err)
	//}
	//createTables(db)

	//_, err = db.Exec(statement, args...)
	//if err != nil {
		//log.Fatal(err)
	//}
	//defer db.Close()
//}

func Reset() {

}
