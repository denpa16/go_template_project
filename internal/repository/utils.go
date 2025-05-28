package repository

import (
	sq "github.com/Masterminds/squirrel"
	"reflect"
)

func GetDbFieldsWithValues(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	numFields := v.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		isEmpty := reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
		if !isEmpty {
			result[field.Tag.Get("db")] = fieldValue.Interface()
		}
	}
	return result
}

func AddWhereAnd(fields []string, query sq.SelectBuilder, dbFields map[string]interface{}) sq.SelectBuilder {
	for _, field := range fields {
		value, ok := dbFields[field]
		if ok {
			query = query.Where(sq.And{sq.Eq{field: value}})
		}
	}
	return query
}

func AddWhereOr(field string, query sq.SelectBuilder, dbFields map[string]interface{}) sq.SelectBuilder {
	value, ok := dbFields[field]
	if ok {
		query = query.Where(sq.Or{sq.Eq{field: value}})
	}
	return query
}
