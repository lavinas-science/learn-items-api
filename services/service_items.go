package services

import (
	"github.com/lavinas-science/learn-items-api/domain/items"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}
type itemsService struct{}

func (s *itemsService) Create(it items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := it.Save(); err != nil {
		return nil, err
	}
	return &it, nil
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewNotImplementedError("not implemented")
}
