package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var SQL_DB *sql.DB

func InitDb() {
	var err error
	SQL_DB, err = sql.Open("sqlite", "./sql-db/events.db")
	fmt.Println("=================")
	fmt.Println(SQL_DB)
	fmt.Println("=================")
	if err != nil {
		// panic("Could not connext to SQL DB \n" +  err ) // crash eit
		fmt.Println("Could not connect to sql", err)
		return
	}
	fmt.Println("Connection eastablished with SQL")

	SQL_DB.SetMaxOpenConns(10)
	SQL_DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUSerTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        );
        
	`

	result, err := SQL_DB.Exec(createUSerTable)
	if err != nil {
		fmt.Println("Could not create user Table", err)
		return
	}
	fmt.Println("USER ---- Table created ->>>", result)

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
		`

	result, err = SQL_DB.Exec(createEventTable)
	if err != nil {
		fmt.Println("Could not create event Table", err)
		return
	}
	fmt.Println("EVENTS  --- Table created ->>>", result)
}
