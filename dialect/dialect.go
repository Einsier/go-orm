package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

type Dialect interface {
	// DataTypeOf return the name of data type in specific database
	DataTypeOf(typ reflect.Value) string
	// TableExistSQL return the sql statement of whether a table is exist
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect register a dialect to global variable dialectMap
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// GetDialect get a dialect from global variable dialectMap
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
