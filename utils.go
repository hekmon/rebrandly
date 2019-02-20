package rebrandly

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func convertStructToURLQuery(params interface{}) (values url.Values, err error) {
	// Prepare
	if params == nil || (reflect.ValueOf(params).Kind() == reflect.Ptr && reflect.ValueOf(params).IsNil()) {
		return
	}
	refValue := reflect.Indirect(reflect.ValueOf(params))
	refType := refValue.Type()
	if refType.Kind() != reflect.Struct {
		err = fmt.Errorf("params is not a struct or a pointer of struct: %v", reflect.TypeOf(params))
		return
	}
	values = make(url.Values, refType.NumField())
	// Build
	var (
		field reflect.Value
		key   string
	)
	for i := 0; i < refType.NumField(); i++ {
		field = refValue.Field(i)
		if field.IsNil() {
			continue
		}
		key = refType.Field(i).Tag.Get("urlQuery")
		switch typedValue := field.Elem().Interface().(type) {
		case bool:
			if typedValue {
				values.Add(key, "true")
			} else {
				values.Add(key, "false")
			}
		case DomainsType:
			values.Add(key, string(typedValue))
		case OrderBy:
			values.Add(key, string(typedValue))
		case OrderDir:
			values.Add(key, string(typedValue))
		case int:
			values.Add(key, strconv.Itoa(typedValue))
		case string:
			values.Add(key, typedValue)
		default:
			err = fmt.Errorf("elem id %d with urlQuery key '%s' is not supported: %v", i, key, reflect.TypeOf(typedValue))
			return
		}
	}
	return
}
