package goorm

import (
	"database/sql"

	"github.com/einsier/go-orm/llog"
	"github.com/einsier/go-orm/session"
)

type Engine struct {
	db *sql.DB
}

// NewEngine create a instance of Engine, which connect to the database
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		llog.Error(err)
		return
	}

	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		llog.Error(err)
		return
	}

	e = &Engine{db: db}
	llog.Info("Connect database success")
	return
}

// Close close the database connection
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		llog.Error("Failed to close database")
	}
	llog.Info("Close database success")
}

// NewSession create a instance of Session
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
