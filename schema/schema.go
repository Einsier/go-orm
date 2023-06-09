package schema

import (
	"go/ast"
	"reflect"

	"github.com/einsier/go-orm/dialect"
)

// Field represent a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represent a table of database
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse a struct to a Schema instance
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("goorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// RecordValues return the values of dest, like ["Tom", 18]
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var values []interface{}
	for _, field := range schema.Fields {
		values = append(values, destValue.FieldByName(field.Name).Interface())
	}
	return values
}
