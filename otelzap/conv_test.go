package otelzap

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// attrValue returns the value of the attribute with the given key.
func attrValue(t *testing.T, kvs []attribute.KeyValue, key string) attribute.Value {
	t.Helper()
	for _, kv := range kvs {
		if string(kv.Key) == key {
			return kv.Value
		}
	}
	t.Fatalf("attribute %q not found in %v", key, kvs)
	return attribute.Value{}
}

func TestAppendFieldArrayMarshaler(t *testing.T) {
	tests := []struct {
		name  string
		field zapcore.Field
		want  interface{}
	}{
		{
			name:  "strings",
			field: zap.Strings("key", []string{"foo", "bar"}),
			want:  []string{"foo", "bar"},
		},
		{
			name:  "empty",
			field: zap.Strings("key", nil),
			want:  []string{},
		},
		{
			name:  "bools",
			field: zap.Bools("key", []bool{true, false}),
			want:  []bool{true, false},
		},
		{
			name:  "int64s",
			field: zap.Int64s("key", []int64{1, -2}),
			want:  []int64{1, -2},
		},
		{
			name:  "uint32s",
			field: zap.Uint32s("key", []uint32{1, 2}),
			want:  []int64{1, 2},
		},
		{
			name:  "float64s",
			field: zap.Float64s("key", []float64{1.5, -2.5}),
			want:  []float64{1.5, -2.5},
		},
		{
			name:  "durations",
			field: zap.Durations("key", []time.Duration{time.Millisecond, time.Second}),
			want:  []string{"1ms", "1s"},
		},
		{
			name:  "times",
			field: zap.Times("key", []time.Time{time.Unix(0, 0).UTC()}),
			want:  []string{"1970-01-01T00:00:00Z"},
		},
		{
			name:  "byte strings",
			field: zap.ByteStrings("key", [][]byte{[]byte("foo")}),
			want:  []string{"foo"},
		},
		{
			name:  "mixed",
			field: zap.Array("key", mixedArray{}),
			want:  []string{"foo", "1"},
		},
		{
			name:  "objects",
			field: zap.Objects("key", []testObject{{Name: "foo", Count: 1}}),
			want:  []string{`{"count":1,"name":"foo"}`},
		},
		{
			name:  "nested arrays",
			field: zap.Array("key", nestedArray{}),
			want:  []string{`["foo","bar"]`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kvs := appendField(nil, tt.field)
			if len(kvs) != 1 {
				t.Fatalf("got %d attributes, want 1: %v", len(kvs), kvs)
			}

			got := attrValue(t, kvs, "key")
			if kvs[0].Key != "key" {
				t.Fatalf("unexpected key: %q", kvs[0].Key)
			}
			if !reflect.DeepEqual(got.AsInterface(), tt.want) {
				t.Fatalf("got %#v, want %#v", got.AsInterface(), tt.want)
			}
		})
	}
}

func TestAppendFieldObjectMarshaler(t *testing.T) {
	kvs := appendField(nil, zap.Object("key", testObject{Name: "foo", Count: 1}))
	if len(kvs) != 1 {
		t.Fatalf("got %d attributes, want 1: %v", len(kvs), kvs)
	}
	if got, want := attrValue(t, kvs, "key").AsString(), `{"count":1,"name":"foo"}`; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAppendFieldArrayMarshalerError(t *testing.T) {
	kvs := appendField(nil, zap.Array("key", errArray{}))
	if len(kvs) != 1 {
		t.Fatalf("got %d attributes, want 1: %v", len(kvs), kvs)
	}
	got, ok := kvs[0].Value.AsString(), string(kvs[0].Key) == "key_error"
	if !ok || got != "otelzap: unable to marshal array: boom" {
		t.Fatalf("got key=%q value=%q", kvs[0].Key, got)
	}
}

func TestConvertFieldsOtherTypes(t *testing.T) {
	err := errors.New("boom")
	kvs := convertFields([]zapcore.Field{
		zap.String("string", "foo"),
		zap.Int("int", 1),
		zap.Bool("bool", true),
		zap.Error(err),
	})
	if got := attrValue(t, kvs, "string").AsString(); got != "foo" {
		t.Fatalf("string: got %q", got)
	}
	if got := attrValue(t, kvs, "int").AsInt64(); got != 1 {
		t.Fatalf("int: got %d", got)
	}
	if got := attrValue(t, kvs, "bool").AsBool(); !got {
		t.Fatalf("bool: got %v", got)
	}
	if got := attrValue(t, kvs, "exception.message").AsString(); got != "boom" {
		t.Fatalf("exception.message: got %q", got)
	}
}

type testObject struct {
	Name  string
	Count int
}

func (o testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", o.Name)
	enc.AddInt("count", o.Count)
	return nil
}

type mixedArray struct{}

func (mixedArray) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	enc.AppendString("foo")
	enc.AppendInt(1)
	return nil
}

type nestedArray struct{}

func (nestedArray) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return enc.AppendArray(zap.Strings("", []string{"foo", "bar"}).Interface.(zapcore.ArrayMarshaler))
}

type errArray struct{}

func (errArray) MarshalLogArray(zapcore.ArrayEncoder) error {
	return errors.New("boom")
}
