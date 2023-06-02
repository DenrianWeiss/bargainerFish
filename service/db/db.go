package db

import (
	"github.com/DenrianWeiss/barginerFish/service/db/badger"
	"github.com/DenrianWeiss/barginerFish/service/db/provider"
)

var db provider.DbProvider

func GetDb() provider.DbProvider {
	return db
}

func init() {
	badgerI := badger.BadgerProvider{}
	err := badgerI.Init()
	if err == nil {
		db = badgerI
	}
}
