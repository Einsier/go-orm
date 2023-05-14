package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{}) (string, []interface{})

var generators = map[Type]generator{}

func init() {
	generators[INSERT] = genInsertClause
	generators[VALUES] = genValuesClause
	generators[SELECT] = genSelectClause
	generators[LIMIT] = genLimitClause
	generators[WHERE] = genWhereClause
	generators[ORDERBY] = genOrderByClause
	generators[UPDATE] = genUpdateClause
	generators[DELETE] = genDeleteClause
	generators[COUNT] = genCountClause
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

func genInsertClause(values ...interface{}) (string, []interface{}) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

func genValuesClause(values ...interface{}) (string, []interface{}) {
	// VALUES ($v1), ($v2), ...
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

func genSelectClause(values ...interface{}) (string, []interface{}) {
	// SELECT $fields FROM $tableName
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

func genLimitClause(values ...interface{}) (string, []interface{}) {
	// LIMIT $num
	return "LIMIT ?", values
}

func genWhereClause(values ...interface{}) (string, []interface{}) {
	// WHERE $desc
	desc := values[0]
	return fmt.Sprintf("WHERE %v", desc), values[1:]
}

func genOrderByClause(values ...interface{}) (string, []interface{}) {
	// ORDER BY $order
	order := values[0]
	return fmt.Sprintf("ORDER BY %s", order), []interface{}{}
}

func genUpdateClause(values ...interface{}) (string, []interface{}) {
	// UPDATE $tableName SET $set
	tableName := values[0]
	m := values[1].(map[string]interface{})
	var keys []string
	var vars []interface{}
	for k, v := range m {
		keys = append(keys, k+" = ?")
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %v", tableName, strings.Join(keys, ",")), vars
}

func genDeleteClause(values ...interface{}) (string, []interface{}) {
	// DELETE FROM $tableName
	tableName := values[0]
	return fmt.Sprintf("DELETE FROM %s", tableName), []interface{}{}
}

func genCountClause(values ...interface{}) (string, []interface{}) {
	// SELECT COUNT(*) FROM $tableName
	tableName := values[0]
	return genSelectClause(tableName, []string{"count(*)"})
}
