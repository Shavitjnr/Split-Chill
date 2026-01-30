package datastore

import (
	"xorm.io/xorm"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type DataStore struct {
	databases []*Database
}


func (s *DataStore) Count() int {
	return len(s.databases)
}


func (s *DataStore) Get(index int) *Database {
	return s.databases[index]
}


func (s *DataStore) Choose(key int64) *Database {
	return s.databases[0]
}


func (s *DataStore) Query(c core.Context, key int64) *xorm.Session {
	return s.Choose(key).NewSession(c)
}


func (s *DataStore) DoTransaction(key int64, c core.Context, fn func(sess *xorm.Session) error) (err error) {
	return s.Choose(key).DoTransaction(c, fn)
}


func (s *DataStore) SyncStructs(beans ...any) error {
	var err error

	for i := 0; i < len(s.databases); i++ {
		err = s.databases[i].engineGroup.Sync2(beans...)

		if err != nil {
			return err
		}
	}

	return err
}


func NewDataStore(databases ...*Database) (*DataStore, error) {
	if len(databases) < 1 {
		return nil, errs.ErrDatabaseIsNull
	}

	return &DataStore{
		databases: databases,
	}, nil
}
