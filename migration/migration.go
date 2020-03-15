package migration

import (
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
)

type DatabaseRepository interface {
	Migrate() error
}

//TaiKhoanRepoImpl khởi tạo
type DatabaseRepoImpl struct {
	Db *gorm.DB
}

func (u *DatabaseRepoImpl) Migrate() error {
	err := u.Db.AutoMigrate(
	).Error

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}

//NewDatabaseRepo : khởi tạo
func NewDatabaseRepo(db *gorm.DB) DatabaseRepository {
	return &DatabaseRepoImpl{
		Db: db,
	}
}
