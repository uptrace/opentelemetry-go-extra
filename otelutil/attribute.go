package otelutil

import (
	"encoding/json"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
)

func Attribute(key string, value interface{}) attribute.KeyValue {
	switch value := value.(type) {
	case nil:
		return attribute.String(key, "<nil>")
	case string:
		return attribute.String(key, value)
	case int:
		return attribute.Int(key, value)
	case int64:
		return attribute.Int64(key, value)
	case uint64:
		return attribute.Int64(key, int64(value))
	case float64:
		return attribute.Float64(key, value)
	case bool:
		return attribute.Bool(key, value)
	case fmt.Stringer:
		return attribute.String(key, value.String())
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Array:
		rv = rv.Slice(0, rv.Len())
		fallthrough
	case reflect.Slice:
		switch rv.Type().Elem().Kind() {
		case reflect.Bool:
			return attribute.BoolSlice(key, boolSlice(rv))
		case reflect.Int:
			return attribute.IntSlice(key, intSlice(rv))
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return attribute.Int64Slice(key, int64Slice(rv))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return attribute.Int64Slice(key, uint64ToInt64Slice(rv))
		case reflect.Float32, reflect.Float64:
			return attribute.Float64Slice(key, float64Slice(rv))
		case reflect.String:
			return attribute.StringSlice(key, stringSlice(rv))
		default:
			return attribute.KeyValue{Key: attribute.Key(key)}
		}
	case reflect.Bool:
		return attribute.Bool(key, rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return attribute.Int64(key, rv.Int())
	case reflect.Float64:
		return attribute.Float64(key, rv.Float())
	case reflect.String:
		return attribute.String(key, rv.String())
	}
	if b, err := json.Marshal(value); b != nil && err == nil {
		return attribute.String(key, string(b))
	}
	return attribute.String(key, fmt.Sprint(value))
}

func boolSlice(rv reflect.Value) []bool {
	res := make([]bool, rv.Len())
	for i := range res {
		res[i] = rv.Index(i).Bool()
	}
	return res
}

func intSlice(rv reflect.Value) []int {
	res := make([]int, rv.Len())
	for i := range res {
		res[i] = int(rv.Index(i).Int())
	}
	return res
}

func int64Slice(rv reflect.Value) []int64 {
	res := make([]int64, rv.Len())
	for i := range res {
		res[i] = rv.Index(i).Int()
	}
	return res
}

func uint64ToInt64Slice(rv reflect.Value) []int64 {
	res := make([]int64, rv.Len())
	for i := range res {
		res[i] = int64(rv.Index(i).Uint())
	}
	return res
}

func float64Slice(rv reflect.Value) []float64 {
	res := make([]float64, rv.Len())
	for i := range res {
		res[i] = rv.Index(i).Float()
	}
	return res
}

func stringSlice(rv reflect.Value) []string {
	res := make([]string, rv.Len())
	for i := range res {
		res[i] = rv.Index(i).String()
	}
	return res
}

func LogValue(value interface{}) log.Value {
	switch value := value.(type) {
	case nil:
		return log.StringValue("<nil>")
	case string:
		return log.StringValue(value)
	case int:
		return log.IntValue(value)
	case int64:
		return log.Int64Value(value)
	case uint64:
		return log.Int64Value(int64(value))
	case float64:
		return log.Float64Value(value)
	case bool:
		return log.BoolValue(value)
	case fmt.Stringer:
		return log.StringValue(value.String())
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Array:
		rv = rv.Slice(0, rv.Len())
		fallthrough
	case reflect.Slice:
		values := make([]log.Value, rv.Len())
		for i := range values {
			values[i] = LogValue(rv.Index(i).Interface())
		}
		return log.SliceValue(values...)
	case reflect.Bool:
		return log.BoolValue(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return log.Int64Value(rv.Int())
	case reflect.Float64:
		return log.Float64Value(rv.Float())
	case reflect.String:
		return log.StringValue(rv.String())
	}
	if b, err := json.Marshal(value); err == nil {
		return log.StringValue(string(b))
	}
	return log.StringValue(fmt.Sprint(value))
}
