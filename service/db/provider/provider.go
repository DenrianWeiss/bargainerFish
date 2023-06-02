package provider

import "github.com/DenrianWeiss/barginerFish/model"

type DbProvider interface {
	GetRecentItems() (items []model.Item, err error)
	GetItemById(id uint) (item model.Item, err error)
	CreateItem(item model.Item) (err error)
	UpdateItem(item model.Item) (err error)
	DelItem(id uint) (err error)
}
