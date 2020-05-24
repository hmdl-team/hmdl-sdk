package queue

import (
	"encoding/json"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"hmdl-user-service/repository"
)

type UserSubscribeService struct {
	RepoNhanVien      repository.NhanVienRepository
	RepoThemSoHeThong repository.DmThamSoHeThongRepo
}

func (u *UserSubscribeService) Subscribes() {
	sub, err := sdk.Nat.Subscribe("ping", func(msg *nats.Msg) {
		logrus.Errorf(string(msg.Data))
	})
	if err != nil {
		logrus.Errorf("Queue %s | subject %s | %s", sub.Queue, sub.Subject, err.Error())
	}

	subNhanVien, err := sdk.Nat.Subscribe("getDanhSachNhanVien", func(msg *nats.Msg) {
		dsNhanVien, err := u.RepoNhanVien.GetAll()
		dataSend, err := json.Marshal(dsNhanVien)
		if err != nil {
			logrus.Errorf(err.Error())
		}
		err = sdk.Nat.Publish(msg.Reply, dataSend)
		if err != nil {
			logrus.Errorf(err.Error())
		}

	})

	if err != nil {
		logrus.Errorf("Queue %s | subject %s | %s", subNhanVien.Queue, subNhanVien.Subject, err.Error())
	}
}
