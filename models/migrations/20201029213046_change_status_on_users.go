package migrations

/*
Remove this migration. Sqlite don't support DROP COLUMN command.
 */
func changeStatusOnUsers() error {
	up := mm.ChangeFuncWrap(
		`ALTER TABLE users DROP COLUMN status;`,
		`ALTER TABLE users ADD COLUMN status VARCHAR(15), ADD INDEX idx_users_status ( status );`,
	)
	down := mm.ChangeFuncWrap(
		`ALTER TABLE users DROP COLUMN status;`,
		`ALTER TABLE users ADD COLUMN status TINYINT;`,
	)
	return mm.Change(up, down)
}
