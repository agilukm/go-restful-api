package utils

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
