package sdk

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var (
	Esc *elastic.Client
)

func ConnectElastic() error {
	url := os.Getenv("ES_URL")
	if url == "" {
		logrus.Warnln("Biến môi trường ES_URL chưa được khai báo, module Elastic sẽ không được sử dụng.")
		return nil
	}

	client, err := elastic.NewClient(
		elastic.SetURL(os.Getenv("ES_URL")),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),

		// elastic.SetRetrier(NewCustomRetrier()),

		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		//elastic.SetHeaders(http.Header{
		//	"X-Caller-Id": []string{"..."},
		//}),
	)
	if err != nil {
		return err
	}

	Esc = client

	logrus.Infof("Connected ElasticSearch: %s", url)

	return nil
}
