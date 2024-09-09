package mapstructure

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"
	"time"
)

type Basic struct {
	Vstring     string
	Vint        int
	Vint8       int8
	Vint16      int16
	Vint32      int32
	Vint64      int64
	Vuint       uint
	Vbool       bool
	Vfloat      float64
	Vextra      string
	vsilent     bool
	Vdata       any
	VjsonInt    int
	VjsonUint   uint
	VjsonUint64 uint64
	VjsonFloat  float64
	VjsonNumber json.Number
}

type BasicPointer struct {
	Vstring     *string
	Vint        *int
	Vuint       *uint
	Vbool       *bool
	Vfloat      *float64
	Vextra      *string
	vsilent     *bool
	Vdata       *any
	VjsonInt    *int
	VjsonFloat  *float64
	VjsonNumber *json.Number
}

type BasicSquash struct {
	Test Basic `mapstructure:",squash"`
}

type Embedded struct {
	Basic
	Vunique string
}

type EmbeddedPointer struct {
	*Basic
	Vunique string
}

type EmbeddedSquash struct {
	Basic   `mapstructure:",squash"`
	Vunique string
}

type EmbeddedPointerSquash struct {
	*Basic  `mapstructure:",squash"`
	Vunique string
}

type BasicMapStructure struct {
	Vunique string     `mapstructure:"vunique"`
	Vtime   *time.Time `mapstructure:"time"`
}

type NestedPointerWithMapstructure struct {
	Vbar *BasicMapStructure `mapstructure:"vbar"`
}

type EmbeddedPointerSquashWithNestedMapstructure struct {
	*NestedPointerWithMapstructure `mapstructure:",squash"`
	Vunique                        string
}

type EmbeddedAndNamed struct {
	Basic
	Named   Basic
	Vunique string
}

type SliceAlias []string

type EmbeddedSlice struct {
	SliceAlias `mapstructure:"slice_alias"`
	Vunique    string
}

type ArrayAlias [2]string

type EmbeddedArray struct {
	ArrayAlias `mapstructure:"array_alias"`
	Vunique    string
}

type SquashOnNonStructType struct {
	InvalidSquashType int `mapstructure:",squash"`
}

type Map struct {
	Vfoo   string
	Vother map[string]string
}

type MapOfStruct struct {
	Value map[string]Basic
}

type Nested struct {
	Vfoo string
	Vbar Basic
}

type NestedPointer struct {
	Vfoo string
	Vbar *Basic
}

type NilInterface struct {
	W io.Writer
}

type NilPointer struct {
	Value *string
}

type Slice struct {
	Vfoo string
	Vbar []string
}

type SliceOfByte struct {
	Vfoo string
	Vbar []byte
}

type SliceOfAlias struct {
	Vfoo string
	Vbar SliceAlias
}

type SliceOfStruct struct {
	Value []Basic
}

type SlicePointer struct {
	Vbar *[]string
}

type Array struct {
	Vfoo string
	Vbar [2]string
}

type ArrayOfStruct struct {
	Value [2]Basic
}

type Func struct {
	Foo func() string
}

type Tagged struct {
	Extra string `mapstructure:"bar,what,what"`
	Value string `mapstructure:"foo"`
}

type Remainder struct {
	A     string
	Extra map[string]any `mapstructure:",remain"`
}

type StructWithOmitEmpty struct {
	VisibleStringField string         `mapstructure:"visible-string"`
	OmitStringField    string         `mapstructure:"omittable-string,omitempty"`
	VisibleIntField    int            `mapstructure:"visible-int"`
	OmitIntField       int            `mapstructure:"omittable-int,omitempty"`
	VisibleFloatField  float64        `mapstructure:"visible-float"`
	OmitFloatField     float64        `mapstructure:"omittable-float,omitempty"`
	VisibleSliceField  []any          `mapstructure:"visible-slice"`
	OmitSliceField     []any          `mapstructure:"omittable-slice,omitempty"`
	VisibleMapField    map[string]any `mapstructure:"visible-map"`
	OmitMapField       map[string]any `mapstructure:"omittable-map,omitempty"`
	NestedField        *Nested        `mapstructure:"visible-nested"`
	OmitNestedField    *Nested        `mapstructure:"omittable-nested,omitempty"`
}

type TypeConversionResult struct {
	IntToFloat         float32
	IntToUint          uint
	IntToBool          bool
	IntToString        string
	UintToInt          int
	UintToFloat        float32
	UintToBool         bool
	UintToString       string
	BoolToInt          int
	BoolToUint         uint
	BoolToFloat        float32
	BoolToString       string
	FloatToInt         int
	FloatToUint        uint
	FloatToBool        bool
	FloatToString      string
	SliceUint8ToString string
	StringToSliceUint8 []byte
	ArrayUint8ToString string
	StringToInt        int
	StringToUint       uint
	StringToBool       bool
	StringToFloat      float32
	StringToStrSlice   []string
	StringToIntSlice   []int
	StringToStrArray   [1]string
	StringToIntArray   [1]int
	SliceToMap         map[string]any
	MapToSlice         []any
	ArrayToMap         map[string]any
	MapToArray         [1]any
}

func ptr[T any](v T) *T { return &v }

func TestDecode_Basic(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vstring": "foo",
		"vint":    42,
		"Vuint":   42,
		"vbool":   true,
		"Vfloat":  42.42,
		"vsilent": true,
		"vdata":   42,
	}

	var result Basic
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}

	if result.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Vstring)
	}

	if result.Vint != 42 {
		t.Errorf("vint value should be 42: %#v", result.Vint)
	}

	if result.Vuint != 42 {
		t.Errorf("vuint value should be 42: %#v", result.Vuint)
	}

	if result.Vbool != true {
		t.Errorf("vbool value should be true: %#v", result.Vbool)
	}

	if result.Vfloat != 42.42 {
		t.Errorf("vfloat value should be 42.42: %#v", result.Vfloat)
	}

	if result.Vextra != "" {
		t.Errorf("vextra value should be empty: %#v", result.Vextra)
	}

	if result.vsilent != false {
		t.Error("vsilent should not be set, it is unexported")
	}

	if result.Vdata != 42 {
		t.Errorf("vdata should be 42: %#v", result.Vdata)
	}
}

func TestDecode_Basic_IntWithFloat(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vint": float64(42),
	}

	var result Basic
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}
}

func TestDecode_Basic_Merge(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vint": 42,
	}

	var result Basic
	result.Vuint = 100
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}

	expected := Basic{
		Vint:  42,
		Vuint: 100,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad: %#v", result)
	}
}

// Test for issue #46.
func TestDecode_Basic_Struct(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vdata": map[string]any{
			"vstring": "foo",
		},
	}

	var result, inner Basic
	result.Vdata = &inner
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}
	expected := Basic{
		Vdata: &Basic{
			Vstring: "foo",
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad: %#v", result)
	}
}

func TestDecode_Basic_interfaceStruct(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vstring": "foo",
	}

	var iface any = &Basic{}
	err := Decode(input, &iface)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}

	expected := &Basic{
		Vstring: "foo",
	}
	if !reflect.DeepEqual(iface, expected) {
		t.Fatalf("bad: %#v", iface)
	}
}

// Issue 187
func TestDecode_Basic_interfaceStructNonPtr(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vstring": "foo",
	}

	var iface any = Basic{}
	err := Decode(input, &iface)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}

	expected := Basic{
		Vstring: "foo",
	}
	if !reflect.DeepEqual(iface, expected) {
		t.Fatalf("bad: %#v", iface)
	}
}

func TestDecode_BasicSquash(t *testing.T) {
	t.Parallel()

	input := map[string]any{
		"vstring": "foo",
	}

	var result BasicSquash
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err.Error())
	}

	if result.Test.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Test.Vstring)
	}
}

func TestDecodeFrom_BasicSquash(t *testing.T) {
	t.Parallel()

	input := BasicSquash{
		Test: Basic{
			Vstring: "foo",
		},
	}

	var result map[string]any
	err := Decode(input, &result)
	if err != nil {
		t.Fatalf("got an err: %s", err)
	}
}
