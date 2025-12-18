package migrations

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		dob DATE NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}
