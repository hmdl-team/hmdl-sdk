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
		logrus.Warning("Không có biến môi trường NAT_URI, nên module NAT sẽ không hoạt động.")
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

// Lấy giá trị đầu vào
// vPt là pointer của biến đầu vào sẽ được xuất ra
func (s *NatMessage) Bind(vPt interface{}) error {
	err := json.Unmarshal(s.Data, vPt)
	if err != nil {
		s.Respond(NewErrorResponse(http.StatusBadRequest, err).ToBytes())
		return err
	}
	return nil
}

// Hàm xử lý error và respond về kết quả
func (s *NatMessage) HandleError(err error) {
	if ae, ok := err.(*AppError); ok {
		s.Respond(NewErrorResponse(ae.StatusCode, err).ToBytes())
		return
	}
	s.Respond(NewErrorResponse(http.StatusInternalServerError, err).ToBytes())
}

// Hàm respond kết quả success về cho client thường sử dụng cho request-reply
// [data] data của kết quả trả về sẽ được bọc trong lớp response trước khi trả về.
func (s *NatMessage) Ok(data interface{}) {
	s.Respond(NewSuccessResponse(data).ToBytes())
}

// Hàm respond bad request về cho client thường sử dụng cho request-reply
func (s *NatMessage) BadRequest(err error) {
	s.Respond(NewErrorResponse(http.StatusBadRequest, err).ToBytes())
}

// Hàm respond conflict request về cho client thường sử dụng cho request-reply
func (s *NatMessage) Conflict(err error) {
	s.Respond(NewErrorResponse(http.StatusConflict, err).ToBytes())
}

// Hàm respond not found request về cho client thường sử dụng cho request-reply
func (s *NatMessage) NotFound(err error) {
	s.Respond(NewErrorResponse(http.StatusNotFound, err).ToBytes())
}

// Hàm respond unauthorized request về cho client thường sử dụng cho request-reply
func (s *NatMessage) Unauthorized(err error) {
	s.Respond(NewErrorResponse(http.StatusUnauthorized, err).ToBytes())
}

// Hàm respond internal server error request về cho client thường sử dụng cho request-reply
func (s *NatMessage) InternalServerError(err error) {
	s.Respond(NewErrorResponse(http.StatusInternalServerError, err).ToBytes())
}
