package driver 

import (
	"database/sql"  // Database SQL package to perform queries
	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
)

func DBConn() (db *sql.DB) {

	dbDriver := "mysql"   // Database driver
	dbUser := "root"      // Mysql username
	dbPass := "" // Mysql password
	dbName := "hikari-spring"   // Mysql schema

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}

	// Return db object to be used by other functions
	return db
}