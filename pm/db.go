package pm

import (
	"database/sql"
	"fmt"
	"os"

	"bytespace.network/rpsclient/print"
	_ "github.com/mattn/go-sqlite3"
)

// DB.go
// =====
// This file contains all functionality concerning RPS's local sqlite database
const dbPath = "./packages/pkginfo.db"
const defTbl = `
	CREATE TABLE Packages (
		ID           INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Name         TEXT,
		Version      TEXT,
		File         TEXT,
		Dependencies TEXT
	);`

var db *sql.DB

func InitializeDB() {
	// if theres no database file
	if _, err := os.Stat(dbPath); err != nil {
		// create one!
		CreateDB()
	}

	_db, err := sql.Open("sqlite3", "./packages/pkginfo.db")
	if err != nil {
		print.PrintCF(print.Red, "Could not open local database file '%s'!", dbPath)
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// store this database connection for later use
	db = _db
}

func CreateDB() {
	print.PrintC(print.Yellow, "No local RPS package database could be found. Generating a new one...")

	// create the file
	os.Create(dbPath)

	// open the file as a db
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		print.PrintCF(print.Red, "Could not open local database file '%s'!", dbPath)
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// create the default package table
	db.Exec(defTbl)

	// close the database to save changes
	db.Close()

	// fancy print
	print.PrintCF(print.Green, "Local database has been created at '%s'!", dbPath)
}

func ErrorDB(err error) {
	if err == nil {
		return
	}

	print.PrintC(print.Red, "Communication with local database failed!")
	fmt.Println(err.Error())
	os.Exit(-1)
}

func GetPackage(row *sql.Rows) Package {
	id, name, version, file, dep := GetRow(row)
	return Package{id, name, version, file, dep}
}

func GetRow(row *sql.Rows) (id int, name string, version string, file string, dep string) {
	row.Scan(&id, &name, &version, &file, &dep)
	return
}
