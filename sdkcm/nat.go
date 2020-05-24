package sdkcm

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Nat     *nats.Conn
	NatJson *nats.EncodedConn
)

func ConnectNat() error {
	uri := os.Getenv("NAT_URI")
	nc, err := nats.Connect(uri)
	if err != nil {
		return err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	Nat = nc
	NatJson = ec

	logrus.Infof("Connected to %s", uri)

	return nil
}
