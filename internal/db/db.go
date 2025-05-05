package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "godatabase"
// )

// func Init() {
// 	var err error
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	DB, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 	}

// 	if err := DB.Ping(); err != nil {
// 		log.Fatal("Database unreachable:", err)
// 	}

// 	_, err = DB.Exec(`
//         CREATE TABLE IF NOT EXISTS users (
//             id SERIAL PRIMARY KEY,
//             username TEXT NOT NULL,
//             email TEXT NOT NULL,
//             password TEXT NOT NULL
//         );
//         CREATE TABLE IF NOT EXISTS items (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT,
// 		description TEXT,
// 		price REAL
//         );
//         CREATE TABLE IF NOT EXISTS carts (
// 		user_id INTEGER,
// 		item_id INTEGER
//         );
//     `)
// 	if err != nil {
// 		log.Fatal("Failed to create users or items table:", err)
// 	}
// }

func InitDB(dataSource string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	err = runSchemaFile("internal/db/schema.sql")
	if err != nil {
		log.Fatal("Failed to apply schema:", err)
	}
}

func runSchemaFile(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	schema, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(schema))
	return err
}

func Close() {
	DB.Close()
}
