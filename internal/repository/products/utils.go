package products

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
			tag := field.Tag.Get("db")
			if tag != "" {
				result[tag] = fieldValue.Interface()
			}
		}
	}
	return result
}

func SelectBuilderAddWhereAnd(fields []string, query sq.SelectBuilder, dbFields map[string]interface{}) sq.SelectBuilder {
	for _, field := range fields {
		value, ok := dbFields[field]
		if ok {
			query = query.Where(sq.And{sq.Eq{field: value}})
		}
	}
	return query
}

func SelectBuilderAddWhereOr(field string, query sq.SelectBuilder, dbFields map[string]interface{}) sq.SelectBuilder {
	value, ok := dbFields[field]
	if ok {
		query = query.Where(sq.Or{sq.Eq{field: value}})
	}
	return query
}

func UpdateBuilderAddWhereAnd(fields []string, query sq.UpdateBuilder, dbFields map[string]interface{}) sq.UpdateBuilder {
	for _, field := range fields {
		value, ok := dbFields[field]
		if ok {
			query = query.Where(sq.And{sq.Eq{field: value}})
		}
	}
	return query
}

func UpdateBuilderAddWhereOr(field string, query sq.UpdateBuilder, dbFields map[string]interface{}) sq.UpdateBuilder {
	value, ok := dbFields[field]
	if ok {
		query = query.Where(sq.Or{sq.Eq{field: value}})
	}
	return query
}

func DeleteBuilderAddWhereAnd(fields []string, query sq.DeleteBuilder, dbFields map[string]interface{}) sq.DeleteBuilder {
	for _, field := range fields {
		value, ok := dbFields[field]
		if ok {
			query = query.Where(sq.And{sq.Eq{field: value}})
		}
	}
	return query
}

func DeleteBuilderAddWhereOr(field string, query sq.DeleteBuilder, dbFields map[string]interface{}) sq.DeleteBuilder {
	value, ok := dbFields[field]
	if ok {
		query = query.Where(sq.Or{sq.Eq{field: value}})
	}
	return query
}
