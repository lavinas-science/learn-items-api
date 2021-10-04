package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lavinas-science/learn-items-api/clients/elasticsearch"
	"github.com/lavinas-science/learn-items-api/domain/queries"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem = "item"
)

func (i *Item) Save() rest_errors.RestErr {
	r, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item")
	}
	i.Id = r.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	r, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprint("no item found with id", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}
	f, err := r.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}	
	if err = json.Unmarshal(f, i); err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}
	i.Id = r.Id
	return nil
}

func (i *Item) Search(q queries.EsQuery) ([]Item, rest_errors.RestErr) {
	s, err := elasticsearch.Client.Search(indexItems, q.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying search documents")
	}
	r := make([]Item, s.TotalHits())
	for j, i := range s.Hits.Hits {
		f, err := i.Source.MarshalJSON()
		if err != nil {
			return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
		}
		var it Item
		if err = json.Unmarshal(f, &it); err != nil {
			return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
		}
		it.Id = i.Id
		r[j] = it
	}
	if len(r) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return r, nil
}
