package migration

import (
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"hmdl-user-service/models/data_user"
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
		data_user.DmThamSoHeThong{},
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
