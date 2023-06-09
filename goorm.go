package orm

import (
	"database/sql"

	"github.com/einsier/go-orm/dialect"
	"github.com/einsier/go-orm/logger"
	"github.com/einsier/go-orm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine create a instance of Engine, which connect to the database
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logger.Error(err)
		return
	}

	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		logger.Error(err)
		return
	}
	// make sure the specified dialect exists
	d, ok := dialect.GetDialect(driver)
	if !ok {
		logger.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: d}
	logger.Info("Connect database success")
	return
}

// Close close the database connection
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		logger.Error("Failed to close database")
	}
	logger.Info("Close database success")
}

// NewSession create a instance of Session
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}

// TxFunc defines the function type that can be executed in a transaction.
type TxFunc func(*session.Session) (interface{}, error)

// Transaction executes a function in a transaction.
func (e *Engine) Transaction(f TxFunc) (res interface{}, err error) {
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			s.Rollback()
			panic(p)
		} else if err != nil {
			s.Rollback()
		} else {
			defer func() {
				if err != nil {
					s.Rollback()
				}
			}()
			err = s.Commit() // err is nil; if Commit returns error update err
		}
	}()

	return f(s)
}
