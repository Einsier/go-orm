package session

import (
	"database/sql"
	"strings"

	"github.com/einsier/go-orm/clause"
	"github.com/einsier/go-orm/dialect"
	"github.com/einsier/go-orm/logger"
	"github.com/einsier/go-orm/schema"
)

// Session keep a pointer to sql.DB and provides all execution of all
type Session struct {
	db       *sql.DB
	tx       *sql.Tx // transaction
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	// sql and sqlVars are used to store the sql statement and its variables.
	sql     strings.Builder
	sqlVars []interface{}
}

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec execs a sql statement with given sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		logger.Error(err)
	}
	return
}

// QueryRow executes a sql statement with given sqlVars and returns a single row
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows executes a sql statement with given sqlVars and returns multiple rows
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		logger.Error(err)
	}
	return
}
