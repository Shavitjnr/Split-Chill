package datastore

import (
	"xorm.io/xorm"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type Database struct {
	databaseType string
	engineGroup  *xorm.EngineGroup
}


func (db *Database) NewSession(c core.Context) *xorm.Session {
	return db.engineGroup.Context(NewXOrmContextAdapter(c))
}


func (db *Database) DoTransaction(c core.Context, fn func(sess *xorm.Session) error) (err error) {
	sess := db.engineGroup.NewSession()

	if c != nil {
		sess.Context(NewXOrmContextAdapter(c))
	}

	defer sess.Close()

	if err = sess.Begin(); err != nil {
		return err
	}

	if err = fn(sess); err != nil {
		_ = sess.Rollback()
		return err
	}

	if err = sess.Commit(); err != nil {
		return err
	}

	return nil
}


func (db *Database) SetSavePoint(sess *xorm.Session, savePointName string) error {
	if db.databaseType == settings.PostgresDbType {
		_, err := sess.Exec("SAVEPOINT " + savePointName)
		return err
	}

	return nil
}


func (db *Database) RollbackToSavePoint(sess *xorm.Session, savePointName string) error {
	if db.databaseType == settings.PostgresDbType {
		_, err := sess.Exec("ROLLBACK TO SAVEPOINT " + savePointName)
		return err
	}

	return nil
}
