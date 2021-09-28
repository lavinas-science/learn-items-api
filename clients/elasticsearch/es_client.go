package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/lavinas-science/learn-utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	r, err := c.client.Index().Type("item").Index(index).BodyJson(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error elasticSearch indexing %s", index), err)
		return nil, err
	}
	return r, nil
}

func (c *esClient) setClient(ec *elastic.Client) {
	c.client = ec
}

func Init() {
	logR := logger.GetLogger()
	cl, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(logR),
		elastic.SetInfoLog(logR),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(cl)
}
