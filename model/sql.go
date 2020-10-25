package model

import "database/sql"

// DBFunc database functions
type DBFunc struct {
	dbConn *sql.DB
}
