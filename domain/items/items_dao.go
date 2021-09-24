package items

import (
	"github.com/lavinas-science/learn-items-api/clients/elasticsearch"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	r, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to saver item")
	}
	i.Id = r.Id
	return nil
}
