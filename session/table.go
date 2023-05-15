package session

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/einsier/go-orm/logger"
	"github.com/einsier/go-orm/schema"
)

// Model specify the model you would like to operate with
func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable return the Schema of the model
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		logger.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable create a table according to the model
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, field.Name+" "+field.Type+" "+field.Tag)
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

// DropTable drop a table according to the model
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable if a table exists
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
