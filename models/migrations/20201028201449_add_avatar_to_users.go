package migrations

func addAvatarToUsers() error {
	type User struct {
		Avatar string `gorm:"type:text;comment:Avatar URL"`
	}
	return mm.ChangeColumn(&User{}, "Avatar")
}
