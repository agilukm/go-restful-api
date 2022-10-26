package helper

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Filter(values url.Values, newStruct any, query string) string {
	metaValue := reflect.ValueOf(newStruct).Elem()

	for k, v := range values {
		field := metaValue.FieldByName(k)
		if field != (reflect.Value{}) {
			if strings.Contains(query, "where") {
				query = fmt.Sprintf("%s and %s like '%%%s%%' ", query, k, v[0])
			} else {
				query = fmt.Sprintf("%s where %s like '%%%s%%' ", query, k, v[0])
			}
		}
	}

	return query
}

func Sort(values url.Values, newStruct any, query string) string {
	sortBy := values.Get("sort")
	metaValue := reflect.ValueOf(newStruct).Elem()

	if sortBy == "" {
		sortBy = "id.desc"
	}

	field, order := validateAndReturnSortQuery(sortBy)

	queryField := metaValue.FieldByName(field)

	if queryField != (reflect.Value{}) {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, field, order)
	} else {
		query = fmt.Sprintf("%s ORDER BY id desc", query)
	}

	return query
}

func validateAndReturnSortQuery(sortBy string) (string, string) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", ""
	}
	field, order := splits[0], splits[1]

	if order != "desc" && order != "asc" {
		return "", ""
	}

	return field, order
}
