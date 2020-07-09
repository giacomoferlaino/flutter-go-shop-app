package app

import "database/sql"

// State contains the current application state
type State struct {
	Database *sql.DB
}
