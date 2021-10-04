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
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string)(*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	r, err := c.client.Index().Type(docType).Index(index).BodyJson(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error elasticSearch indexing %s", index), err)
		return nil, err
	}
	return r, nil
}

func (c *esClient) Get (index string, docType string, id string) (*elastic.GetResult, error)  {
	ctx := context.Background()
	r, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error elasticSearch indexing %s", index), err)
		return nil, err
	}
	return r, nil
}

func (c *esClient) setClient(ec *elastic.Client) {
	c.client = ec
}

func (c *esClient) Search(i string, q elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	r, err := c.client.Search(i).Query(q).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when to search document %s", i), err)
		return nil, err
	}
	return r, nil
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
