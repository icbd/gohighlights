package migrations

import (
	"github.com/icbd/gohighlights/models"
	mgr "github.com/icbd/gorm-migration"
)

var mm *mgr.MigrationManger

func init() {
	mm = mgr.NewMigrationManger(models.DB(), mgr.Check)
	mm.RegisterFunctions(
		createUsers,
		addAvatarToUsers,
		createSessions,
		createMarks,
		createComments,
		addIndexOnUsers,
		addIndexOnSessions,
		addIndexOnMarks,
		addIndexOnComments,
	)
}

func Migrate(t mgr.MigrateType) {
	mm.Type = t
	mm.Migrate()
}
