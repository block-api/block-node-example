package migration

import "github.com/block-api/block-node/db"

const createUserTable string = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
`

const dropUserTable string = `DROP TABLE IF EXISTS test;`

func UserMigration() db.SQLMigration {
	return db.NewSQLMigration("user", createUserTable, dropUserTable)
}
