package otelzap

import (
	"encoding/json"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap/zapcore"
)

// arrayEncoder implements zapcore.ArrayEncoder and collects the marshaled
// values so that they can be reported as a typed attribute ([]bool, []int64,
// []float64 or []string).
//
// Arrays that mix several types, as well as nested arrays and objects, are
// rendered as strings, which is the only representation OTel attributes
// support for heterogeneous elements.
type arrayEncoder struct {
	values []interface{}
}

var _ zapcore.ArrayEncoder = (*arrayEncoder)(nil)

// keyValue converts the collected values into an attribute.
func (e *arrayEncoder) keyValue(key string) attribute.KeyValue {
	if len(e.values) == 0 {
		return attribute.StringSlice(key, []string{})
	}

	if vs, ok := toInt64Slice(e.values); ok {
		return attribute.Int64Slice(key, vs)
	}
	if vs, ok := toFloat64Slice(e.values); ok {
		return attribute.Float64Slice(key, vs)
	}
	if vs, ok := toStringSlice(e.values); ok {
		return attribute.StringSlice(key, vs)
	}
	if vs, ok := toBoolSlice(e.values); ok {
		return attribute.BoolSlice(key, vs)
	}

	vs := make([]string, len(e.values))
	for i, v := range e.values {
		vs[i] = fmt.Sprint(v)
	}
	return attribute.StringSlice(key, vs)
}

// objectKeyValue converts marshaled object fields into a JSON string attribute,
// because OTel attributes have no map type.
func objectKeyValue(key string, fields map[string]interface{}) attribute.KeyValue {
	if bs, err := json.Marshal(fields); err == nil {
		return attribute.String(key, string(bs))
	}
	return attribute.String(key, fmt.Sprint(fields))
}

func toInt64Slice(values []interface{}) ([]int64, bool) {
	vs := make([]int64, len(values))
	for i, v := range values {
		n, ok := v.(int64)
		if !ok {
			return nil, false
		}
		vs[i] = n
	}
	return vs, true
}

func toFloat64Slice(values []interface{}) ([]float64, bool) {
	vs := make([]float64, len(values))
	for i, v := range values {
		switch n := v.(type) {
		case int64:
			vs[i] = float64(n)
		case float64:
			vs[i] = n
		default:
			return nil, false
		}
	}
	return vs, true
}

func toStringSlice(values []interface{}) ([]string, bool) {
	vs := make([]string, len(values))
	for i, v := range values {
		s, ok := v.(string)
		if !ok {
			return nil, false
		}
		vs[i] = s
	}
	return vs, true
}

func toBoolSlice(values []interface{}) ([]bool, bool) {
	vs := make([]bool, len(values))
	for i, v := range values {
		b, ok := v.(bool)
		if !ok {
			return nil, false
		}
		vs[i] = b
	}
	return vs, true
}

func (e *arrayEncoder) AppendBool(v bool)             { e.values = append(e.values, v) }
func (e *arrayEncoder) AppendByteString(v []byte)     { e.values = append(e.values, string(v)) }
func (e *arrayEncoder) AppendComplex64(v complex64)   { e.values = append(e.values, fmt.Sprint(v)) }
func (e *arrayEncoder) AppendComplex128(v complex128) { e.values = append(e.values, fmt.Sprint(v)) }
func (e *arrayEncoder) AppendDuration(v time.Duration) {
	e.values = append(e.values, v.String())
}
func (e *arrayEncoder) AppendFloat32(v float32) { e.values = append(e.values, float64(v)) }
func (e *arrayEncoder) AppendFloat64(v float64) { e.values = append(e.values, v) }
func (e *arrayEncoder) AppendInt(v int)         { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendInt8(v int8)       { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendInt16(v int16)     { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendInt32(v int32)     { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendInt64(v int64)     { e.values = append(e.values, v) }
func (e *arrayEncoder) AppendString(v string)   { e.values = append(e.values, v) }
func (e *arrayEncoder) AppendTime(v time.Time) {
	e.values = append(e.values, v.Format(time.RFC3339Nano))
}
func (e *arrayEncoder) AppendUint(v uint)       { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendUint8(v uint8)     { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendUint16(v uint16)   { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendUint32(v uint32)   { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendUint64(v uint64)   { e.values = append(e.values, int64(v)) }
func (e *arrayEncoder) AppendUintptr(v uintptr) { e.values = append(e.values, int64(v)) }

func (e *arrayEncoder) AppendArray(v zapcore.ArrayMarshaler) error {
	sub := new(arrayEncoder)
	if err := v.MarshalLogArray(sub); err != nil {
		return err
	}
	bs, err := json.Marshal(sub.values)
	if err != nil {
		return err
	}
	e.values = append(e.values, string(bs))
	return nil
}

func (e *arrayEncoder) AppendObject(v zapcore.ObjectMarshaler) error {
	enc := zapcore.NewMapObjectEncoder()
	if err := v.MarshalLogObject(enc); err != nil {
		return err
	}
	bs, err := json.Marshal(enc.Fields)
	if err != nil {
		return err
	}
	e.values = append(e.values, string(bs))
	return nil
}

func (e *arrayEncoder) AppendReflected(v interface{}) error {
	switch v.(type) {
	case bool, int64, float64, string:
		e.values = append(e.values, v)
	default:
		e.values = append(e.values, fmt.Sprint(v))
	}
	return nil
}
