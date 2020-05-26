package sdk

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	Nat     *nats.Conn
	NatJson *nats.EncodedConn
)

func ConnectNat() error {
	uri := os.Getenv("NAT_URI")

	// ignore if NAT URI not set
	if uri == "" {
		return nil
	}

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

	logrus.Infof("Connected NAT: %s", uri)

	return nil
}



type NatMessage struct {
	*nats.Msg
}

// constructor for NatMessage
func NatSender(msg *nats.Msg) *NatMessage {
	return &NatMessage{Msg: msg}
}

func (s *NatMessage) Input(vPt interface{}) bool {
	err := json.Unmarshal(s.Data, vPt)
	if err != nil {
		_ = s.Respond(NewErrorResponse(http.StatusBadRequest, err).ToBytes())
		return false
	}
	return true
}

func (s *NatMessage) HandleError(err error) {
	if ae, ok := err.(*AppError); ok {
		_ = s.Respond(NewErrorResponse(ae.StatusCode, err).ToBytes())
		return
	}
	_ = s.Respond(NewErrorResponse(http.StatusInternalServerError, err).ToBytes())
}

func (s *NatMessage) Ok(data interface{}) {
	_ = s.Respond(NewSuccessResponse(data).ToBytes())
}