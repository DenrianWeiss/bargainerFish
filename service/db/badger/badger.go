package badger

import (
	"encoding/json"
	"github.com/DenrianWeiss/barginerFish/model"
	"github.com/dgraph-io/badger/v4"
	"strconv"
)

type BadgerProvider struct {
}

var badgerDb *badger.DB

func (b BadgerProvider) Init() (err error) {
	badgerDb, err = badger.Open(badger.DefaultOptions("./db/badger"))
	if err != nil {
		return err
	}
	return nil
}

func (b BadgerProvider) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := badgerDb.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

func (b BadgerProvider) GetRecentItems() (items []model.Item, err error) {
	result := make([]model.Item, 0)
	err = badgerDb.View(func(txn *badger.Txn) error {
		id, err := txn.Get([]byte("id"))
		if id == nil {
			return err
		}
		valueCopy, err := id.ValueCopy(nil)
		if err != nil {
			return err
		}
		idInt, _ := strconv.ParseInt(string(valueCopy), 10, 64)
		for i := idInt; i > 0 && i > i-20; i-- {
			item, err := txn.Get([]byte(strconv.FormatInt(i, 10)))
			if err != nil {
				continue
			}
			valueCopy, err = item.ValueCopy(nil)
			if err != nil {
				continue
			}
			itemM := model.Item{}
			json.Unmarshal(valueCopy, &itemM)
			result = append(result, itemM)
		}
		return nil
	})
	return result, err
}
func (b BadgerProvider) GetItemById(id uint) (item model.Item, err error) {
	var result model.Item
	err = badgerDb.View(func(txn *badger.Txn) error {
		dbItem, err := txn.Get([]byte(strconv.Itoa(int(id))))
		if err != nil {
			return err
		}
		valueCopy, err := dbItem.ValueCopy(nil)
		if err != nil {
			return err
		}
		err = json.Unmarshal(valueCopy, &result)
		if err != nil {
			return err
		}
		return err
	})
	return result, err
}
func (b BadgerProvider) CreateItem(item model.Item) (err error) {
	err = badgerDb.Update(func(txn *badger.Txn) error {
		id, err := txn.Get([]byte("id"))
		if err != nil {
			txn.Set([]byte("id"), []byte("0"))
			id, err = txn.Get([]byte("id"))
		}
		valueCopy, err := id.ValueCopy(nil)
		if err != nil {
			return err
		}
		idInt, _ := strconv.ParseInt(string(valueCopy), 10, 64)
		item.Id = uint(idInt + 1)
		err = txn.Set([]byte("id"), []byte(strconv.FormatInt(idInt+1, 10)))
		if err != nil {
			return err
		}
		itemBytes, err := json.Marshal(item)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(strconv.Itoa(int(item.Id))), itemBytes)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
func (b BadgerProvider) UpdateItem(item model.Item) (err error) {
	err = badgerDb.Update(func(txn *badger.Txn) error {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			return err
		}
		err = txn.Set([]byte(strconv.Itoa(int(item.Id))), itemBytes)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (b BadgerProvider) DelItem(id uint) (err error) {
	err = badgerDb.Update(func(txn *badger.Txn) error {
		err = txn.Delete([]byte(strconv.Itoa(int(id))))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
