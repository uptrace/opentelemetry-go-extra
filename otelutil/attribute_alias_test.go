package otelutil

import (
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

// MyString is an alias used to reproduce the panic when passing a slice of a
// typed alias to Attribute. The underlying type is string so it should be
// handled like a []string, but the current implementation type asserts to
// []string directly and panics for alias types.
type MyString string

type (
	MyBool    bool
	MyInt32   int32
	MyUint16  uint16
	MyFloat32 float32
)

func TestAttributeAcceptsSliceAlias(t *testing.T) {
	cases := []struct {
		name     string
		value    interface{}
		wantType attribute.Type
		want     interface{}
	}{
		{
			name:     "string alias",
			value:    []MyString{"foo", "bar"},
			wantType: attribute.STRINGSLICE,
			want:     []string{"foo", "bar"},
		},
		{
			name:     "bool alias",
			value:    []MyBool{true, false},
			wantType: attribute.BOOLSLICE,
			want:     []bool{true, false},
		},
		{
			name:     "int32 alias",
			value:    []MyInt32{1, 2},
			wantType: attribute.INT64SLICE,
			want:     []int64{1, 2},
		},
		{
			name:     "uint16 alias",
			value:    []MyUint16{3, 4},
			wantType: attribute.INT64SLICE,
			want:     []int64{3, 4},
		},
		{
			name:     "float32 alias",
			value:    []MyFloat32{1.5, 2.5},
			wantType: attribute.FLOAT64SLICE,
			want:     []float64{1.5, 2.5},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Attribute panicked for %T: %v", tt.value, r)
				}
			}()

			kv := Attribute("alias", tt.value)
			if kv.Value.Type() != tt.wantType {
				t.Fatalf("unexpected attribute type: got %v want %v", kv.Value.Type(), tt.wantType)
			}

			switch tt.wantType {
			case attribute.STRINGSLICE:
				if got := kv.Value.AsStringSlice(); !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("strings differ: got %v want %v", got, tt.want)
				}
			case attribute.BOOLSLICE:
				if got := kv.Value.AsBoolSlice(); !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("bools differ: got %v want %v", got, tt.want)
				}
			case attribute.INT64SLICE:
				if got := kv.Value.AsInt64Slice(); !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("ints differ: got %v want %v", got, tt.want)
				}
			case attribute.FLOAT64SLICE:
				if got := kv.Value.AsFloat64Slice(); !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("floats differ: got %v want %v", got, tt.want)
				}
			default:
				t.Fatalf("unhandled type in test: %v", tt.wantType)
			}
		})
	}
}
