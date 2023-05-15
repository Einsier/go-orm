package session

import "github.com/einsier/go-orm/llog"

// Begin start a transaction
func (s *Session) Begin() (err error) {
	llog.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		llog.Error(err)
	}
	return
}

// Commit commit a transaction
func (s *Session) Commit() (err error) {
	llog.Info("transaction commit")
	if err := s.tx.Commit(); err != nil {
		llog.Error(err)
	}
	return
}

// Rollback rollback a transaction
func (s *Session) Rollback() (err error) {
	llog.Info("transaction rollback")
	if err := s.tx.Rollback(); err != nil {
		llog.Error(err)
	}
	return
}
