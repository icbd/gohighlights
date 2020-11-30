package migrations

func addIndexOnUsers() error {
	up := mm.ChangeFuncWrap(
		`CREATE UNIQUE INDEX idx_users_email_deletedAt ON users (email,deleted_at);`,
	)
	down := mm.ChangeFuncWrap(
		`DROP INDEX idx_users_email_deletedAt ON users;`,
	)
	return mm.Change(up, down)
}

func addIndexOnSessions() error {
	up := mm.ChangeFuncWrap(
		`CREATE UNIQUE INDEX idx_sessions_token_deletedAt ON sessions (token,deleted_at);`,
	)
	down := mm.ChangeFuncWrap(
		`DROP INDEX idx_sessions_token_deletedAt ON sessions;`,
	)
	return mm.Change(up, down)
}

func addIndexOnMarks() error {
	up := mm.ChangeFuncWrap(
		`CREATE UNIQUE INDEX idx_marks_userID_hashKey_deletedAt ON marks (user_id,hash_key,deleted_at);`,
	)
	down := mm.ChangeFuncWrap(
		`DROP INDEX idx_marks_userID_hashKey_deletedAt ON marks;`,
	)
	return mm.Change(up, down)
}

func addIndexOnComments() error {
	up := mm.ChangeFuncWrap(
		`CREATE UNIQUE INDEX idx_comments_markID_userID_deletedAt ON comments (mark_id,user_id,deleted_at);`,
	)
	down := mm.ChangeFuncWrap(
		`DROP INDEX idx_comments_markID_userID_deletedAt ON comments;`,
	)
	return mm.Change(up, down)
}
