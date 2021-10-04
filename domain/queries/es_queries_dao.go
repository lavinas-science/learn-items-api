package queries

import (
	"github.com/olivere/elastic"
)

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	eQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		eQueries = append(eQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(eQueries...)
	return query
}