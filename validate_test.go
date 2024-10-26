package go_validate

import (
	"encoding/json"
	"testing"
)

func TestValidateInt(t *testing.T) {
	type testRequest struct {
		IntVal   int    `json:"int_val"`
		IntNil   *int   `json:"int_nil"`
		IntPtr   *int   `json:"int_ptr"`
		Int8Val  int8   `json:"int8_val"`
		Int8Ptr  *int8  `json:"int8_ptr"`
		Int16Val int16  `json:"int16_val"`
		Int16Ptr *int16 `json:"int16_ptr"`
		Int32Val int32  `json:"int32_val"`
		Int32Ptr *int32 `json:"int32_ptr"`
		Int64Val int64  `json:"int64_val"`
		Int64Ptr *int64 `json:"int64_ptr"`
	}

	var data map[string]any
	r1text := `{"int_val": 1, "int_nil": null, "int_ptr": 2, "int8_val": 2, "int8_ptr": 3, "int16_val": 3, "int16_ptr": 4, "int32_val": 4, "int32_ptr": 5, "int64_val": 5, "int64_ptr": 6}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.IntVal != 1 {
		t.Errorf("int_val expected 1, got %d", tr1.IntVal)
	}
	if tr1.IntNil != nil {
		t.Errorf("int_nil expected nil, got %v", tr1.IntNil)
	}
	if tr1.IntPtr == nil || *tr1.IntPtr != 2 {
		t.Errorf("int_ptr expected 2, got %v", tr1.IntPtr)
	}
	if tr1.Int8Val != 2 {
		t.Errorf("int8_val expected 2, got %d", tr1.Int8Val)
	}
	if tr1.Int8Ptr == nil || *tr1.Int8Ptr != 3 {
		t.Errorf("int8_ptr expected 3, got %v", tr1.Int8Ptr)
	}
	if tr1.Int16Val != 3 {
		t.Errorf("int16_val expected 3, got %d", tr1.Int16Val)
	}
	if tr1.Int16Ptr == nil || *tr1.Int16Ptr != 4 {
		t.Errorf("int16_ptr expected 4, got %v", tr1.Int16Ptr)
	}
	if tr1.Int32Val != 4 {
		t.Errorf("int32_val expected 4, got %d", tr1.Int32Val)
	}
	if tr1.Int32Ptr == nil || *tr1.Int32Ptr != 5 {
		t.Errorf("int32_ptr expected 5, got %v", tr1.Int32Ptr)
	}
	if tr1.Int64Val != 5 {
		t.Errorf("int64_val expected 5, got %d", tr1.Int64Val)
	}
	if tr1.Int64Ptr == nil || *tr1.Int64Ptr != 6 {
		t.Errorf("int64_ptr expected 6, got %v", tr1.Int64Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"int_val": "str1", "int_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.IntVal != 0 {
		t.Errorf("int_val expected 0, got %d", tr1.IntVal)
	}
	if tr1.IntPtr != nil {
		t.Errorf("int_ptr expected nil, got %v", tr1.IntPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "int_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "int_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateUInt(t *testing.T) {
	type testRequest struct {
		UIntVal   uint    `json:"int_val"`
		UIntNil   *uint   `json:"int_nil"`
		UIntPtr   *uint   `json:"int_ptr"`
		UInt8Val  uint8   `json:"int8_val"`
		UInt8Ptr  *uint8  `json:"int8_ptr"`
		UInt16Val uint16  `json:"int16_val"`
		UInt16Ptr *uint16 `json:"int16_ptr"`
		UInt32Val uint32  `json:"int32_val"`
		UInt32Ptr *uint32 `json:"int32_ptr"`
		UInt64Val uint64  `json:"int64_val"`
		UInt64Ptr *uint64 `json:"int64_ptr"`
	}

	var data map[string]any
	r1text := `{"int_val": 1, "int_nil": null, "int_ptr": 2, "int8_val": 2, "int8_ptr": 3, "int16_val": 3, "int16_ptr": 4, "int32_val": 4, "int32_ptr": 5, "int64_val": 5, "int64_ptr": 6}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.UIntVal != 1 {
		t.Errorf("int_val expected 1, got %d", tr1.UIntVal)
	}
	if tr1.UIntNil != nil {
		t.Errorf("int_nil expected nil, got %v", tr1.UIntNil)
	}
	if tr1.UIntPtr == nil || *tr1.UIntPtr != 2 {
		t.Errorf("int_ptr expected 2, got %v", tr1.UIntPtr)
	}
	if tr1.UInt8Val != 2 {
		t.Errorf("int8_val expected 2, got %d", tr1.UInt8Val)
	}
	if tr1.UInt8Ptr == nil || *tr1.UInt8Ptr != 3 {
		t.Errorf("int8_ptr expected 3, got %v", tr1.UInt8Ptr)
	}
	if tr1.UInt16Val != 3 {
		t.Errorf("int16_val expected 3, got %d", tr1.UInt16Val)
	}
	if tr1.UInt16Ptr == nil || *tr1.UInt16Ptr != 4 {
		t.Errorf("int16_ptr expected 4, got %v", tr1.UInt16Ptr)
	}
	if tr1.UInt32Val != 4 {
		t.Errorf("int32_val expected 4, got %d", tr1.UInt32Val)
	}
	if tr1.UInt32Ptr == nil || *tr1.UInt32Ptr != 5 {
		t.Errorf("int32_ptr expected 5, got %v", tr1.UInt32Ptr)
	}
	if tr1.UInt64Val != 5 {
		t.Errorf("int64_val expected 5, got %d", tr1.UInt64Val)
	}
	if tr1.UInt64Ptr == nil || *tr1.UInt64Ptr != 6 {
		t.Errorf("int64_ptr expected 6, got %v", tr1.UInt64Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"int_val": "str1", "int_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.UIntVal != 0 {
		t.Errorf("int_val expected 0, got %d", tr1.UIntVal)
	}
	if tr1.UIntPtr != nil {
		t.Errorf("int_ptr expected nil, got %v", *tr1.UIntPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "int_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "int_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateByteAndRune(t *testing.T) {
	type testRequest struct {
		ByteVal byte  `json:"byte_val"`
		ByteNil *byte `json:"byte_nil"`
		BytePtr *byte `json:"byte_ptr"`
		RuneVal rune  `json:"rune_val"`
		RunePtr *rune `json:"rune_ptr"`
	}

	var data map[string]any
	r1text := `{"byte_val": 1, "byte_nil": null, "byte_ptr": 2, "rune_val": 2, "rune_ptr": 3}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.ByteVal != 1 {
		t.Errorf("byte_val expected 1, got %d", tr1.ByteVal)
	}
	if tr1.ByteNil != nil {
		t.Errorf("byte_nil expected nil, got %v", tr1.ByteNil)
	}
	if tr1.BytePtr == nil || *tr1.BytePtr != 2 {
		t.Errorf("byte_ptr expected 2, got %v", tr1.BytePtr)
	}
	if tr1.RuneVal != 2 {
		t.Errorf("rune_val expected 2, got %d", tr1.RuneVal)
	}
	if tr1.RunePtr == nil || *tr1.RunePtr != 3 {
		t.Errorf("rune_ptr expected 3, got %v", tr1.RunePtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"byte_val": "str1", "byte_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.ByteVal != 0 {
		t.Errorf("byte_val expected 0, got %d", tr1.ByteVal)
	}
	if tr1.BytePtr != nil {
		t.Errorf("byte_ptr expected nil, got %v", *tr1.BytePtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "byte_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "byte_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateFloat(t *testing.T) {
	type testRequest struct {
		Float64Val float64  `json:"float64_val"`
		Float64Nil *float64 `json:"float64_nil"`
		Float64Ptr *float64 `json:"float64_ptr"`
		Float32Val float32  `json:"float32_val"`
		Float32Ptr *float32 `json:"float32_ptr"`
	}

	var data map[string]any
	r1text := `{"float64_val": 1.23, "float64_nil": null, "float64_ptr": 2.34, "float32_val": 2.55, "float32_ptr": 3.45}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.Float64Val != 1.23 {
		t.Errorf("float64_val expected 1.23, got %f", tr1.Float64Val)
	}
	if tr1.Float64Nil != nil {
		t.Errorf("float64_nil expected nil, got %v", tr1.Float64Nil)
	}
	if tr1.Float64Ptr == nil || *tr1.Float64Ptr != 2.34 {
		t.Errorf("float64_ptr expected 2.34, got %v", tr1.Float64Ptr)
	}
	if tr1.Float32Val != 2.55 {
		t.Errorf("float32_val expected 2.55, got %f", tr1.Float32Val)
	}
	if tr1.Float32Ptr == nil || *tr1.Float32Ptr != 3.45 {
		t.Errorf("float32_ptr expected 3.45, got %v", tr1.Float32Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"float64_val": "str1", "float64_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.Float64Val != 0 {
		t.Errorf("float64_val expected 0, got %f", tr1.Float64Val)
	}
	if tr1.Float64Ptr != nil {
		t.Errorf("float64_ptr expected nil, got %v", *tr1.Float64Ptr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "float64_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "float64_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateBool(t *testing.T) {
	type testRequest struct {
		BoolVal bool  `json:"bool_val"`
		BoolNil *bool `json:"bool_nil"`
		BoolPtr *bool `json:"bool_ptr"`
	}

	var data map[string]any
	r1text := `{"bool_val": true, "bool_nil": null, "bool_ptr": true}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.BoolVal != true {
		t.Errorf("bool_val expected true, got %v", tr1.BoolVal)
	}
	if tr1.BoolNil != nil {
		t.Errorf("bool_nil expected nil, got %v", tr1.BoolNil)
	}
	if tr1.BoolPtr == nil || *tr1.BoolPtr != true {
		t.Errorf("bool_ptr expected true, got %v", tr1.BoolPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"bool_val": "str1", "bool_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.BoolVal != false {
		t.Errorf("bool_val expected false, got %v", tr1.BoolVal)
	}
	if tr1.BoolPtr != nil {
		t.Errorf("bool_ptr expected nil, got %v", *tr1.BoolPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "bool_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "bool_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateString(t *testing.T) {
	type testRequest struct {
		StringVal string  `json:"string_val"`
		StringNil *string `json:"string_nil"`
		StringPtr *string `json:"string_ptr"`
	}

	var data map[string]any
	r1text := `{"string_val": "str1", "string_nil": null, "string_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.StringVal != "str1" {
		t.Errorf("string_val expected true, got %v", tr1.StringVal)
	}
	if tr1.StringNil != nil {
		t.Errorf("string_nil expected nil, got %v", tr1.StringNil)
	}
	if tr1.StringPtr == nil || *tr1.StringPtr != "str2" {
		t.Errorf("string_ptr expected true, got %v", tr1.StringPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"string_val": [], "string_ptr": ["a"]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.StringVal != "" {
		t.Errorf("string_val expected false, got %v", tr1.StringVal)
	}
	if tr1.StringPtr != nil {
		t.Errorf("string_ptr expected nil, got %v", *tr1.StringPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 2 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "string_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "string_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateTypes(t *testing.T) {
	type testIntType int
	const (
		intType0 testIntType = iota
		intType1
		intType2
	)

	type testStrType string
	const (
		strType0 testStrType = ""
		strType1 testStrType = "string1"
		strType2 testStrType = "string2"
	)

	type testRequest struct {
		IntVal    testIntType  `json:"int_val"`
		IntNil    *testIntType `json:"int_nil"`
		IntPtr    *testIntType `json:"int_ptr"`
		StringVal testStrType  `json:"string_val"`
		StringNil *testStrType `json:"string_nil"`
		StringPtr *testStrType `json:"string_ptr"`
	}

	var data map[string]any
	r1text := `{"int_val": 1, "int_nil": null, "int_ptr": 2, "string_val": "string1", "string_nil": null, "string_ptr": "string2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.IntVal != intType1 {
		t.Errorf("int_val expected intType1, got %v", tr1.IntVal)
	}
	if tr1.IntNil != nil {
		t.Errorf("int_nil expected nil, got %v", tr1.IntNil)
	}
	if tr1.IntPtr == nil || *tr1.IntPtr != intType2 {
		t.Errorf("int_ptr expected intType2, got %v", tr1.IntPtr)
	}
	if tr1.StringVal != strType1 {
		t.Errorf("string_val expected strType1, got %v", tr1.StringVal)
	}
	if tr1.StringNil != nil {
		t.Errorf("string_nil expected nil, got %v", tr1.StringNil)
	}
	if tr1.StringPtr == nil || *tr1.StringPtr != strType2 {
		t.Errorf("string_ptr expected strType2, got %v", tr1.StringPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"int_val": "string1", "int_ptr": "string2", "string_val": 1, "string_ptr": 2}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.IntVal != intType0 {
		t.Errorf("int_val expected intType0, got %v", tr1.IntVal)
	}
	if tr1.IntPtr != nil {
		t.Errorf("int_ptr expected nil, got %v", *tr1.IntPtr)
	}
	if tr1.StringVal != strType0 {
		t.Errorf("string_val expected false, got %v", tr1.StringVal)
	}
	if tr1.StringPtr != nil {
		t.Errorf("string_ptr expected nil, got %v", *tr1.StringPtr)
	}
	if len(errors) != 4 {
		t.Errorf("there should be only 4 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "int_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "int_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
		if errors[2].Attribute != "string_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[3].Attribute != "string_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateStructs(t *testing.T) {
	type testSubStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type testStruct struct {
		ID     int            `json:"id"`
		Name   string         `json:"name"`
		Sub    testSubStruct  `json:"sub"`
		SubPtr *testSubStruct `json:"sub_ptr"`
	}

	type testRequest struct {
		ObjVal testStruct  `json:"obj_val"`
		ObjNil *testStruct `json:"obj_nil"`
		ObjPtr *testStruct `json:"obj_ptr"`
	}

	var data map[string]any
	r1text := `{"obj_val": {"id": 1, "name": "Name1", "sub": {"id": 2, "name": "Sub1"}}, "obj_nil": null, "obj_ptr": {"id": 2, "name": "Name2", "sub": {"id": 3, "name": "Sub2"}, "sub_ptr": {"id": 4, "name": "Sub3"}}}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors := Run[testRequest](data, map[string][]any{})
	if tr1.ObjVal.ID != 1 {
		t.Errorf("obj_val.id expected 1, got %v", tr1.ObjVal.ID)
	}
	if tr1.ObjVal.Name != "Name1" {
		t.Errorf("obj_val.name expected name, got %v", tr1.ObjVal.Name)
	}
	if tr1.ObjVal.Sub.ID != 2 {
		t.Errorf("obj_val.sub.id expected 2, got %v", tr1.ObjVal.Sub.ID)
	}
	if tr1.ObjVal.Sub.Name != "Sub1" {
		t.Errorf("obj_val.sub.name expected Sub1, got %v", tr1.ObjVal.Sub.Name)
	}
	if tr1.ObjVal.SubPtr != nil {
		t.Errorf("obj_val.sub_ptr expected nil, got %v", *tr1.ObjVal.SubPtr)
	}
	if tr1.ObjNil != nil {
		t.Errorf("obj_nil.id expected nil, got %v", tr1.ObjNil)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.ID != 2 {
		t.Errorf("obj_nil.id expected 2, got %v", tr1.ObjPtr)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.Name != "Name2" {
		t.Errorf("obj_nil.name expected Name2, got %v", tr1.ObjPtr)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.Sub.ID != 3 {
		t.Errorf("obj_nil.sub.id expected 3, got %v", tr1.ObjPtr)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.Sub.Name != "Sub2" {
		t.Errorf("obj_nil.sub.name expected Sub2, got %v", tr1.ObjPtr)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.SubPtr == nil || tr1.ObjPtr.SubPtr.ID != 4 {
		t.Errorf("obj_nil.sub_ptr.id expected 4, got %v", tr1.ObjPtr)
	}
	if tr1.ObjPtr == nil || tr1.ObjPtr.SubPtr == nil || tr1.ObjPtr.SubPtr.Name != "Sub3" {
		t.Errorf("obj_nil.sub_ptr.name expected Sub3, got %v", tr1.ObjPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"obj_val": 1, "obj_ptr": "string2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr1, errors = Run[testRequest](data, map[string][]any{})
	if tr1.ObjVal.ID != 0 || tr1.ObjVal.Name != "" {
		t.Errorf("obj_val expected empty, got %v", tr1.ObjVal)
	}
	if tr1.ObjNil != nil {
		t.Errorf("obj_nil expected nil, got %v", *tr1.ObjNil)
	}
	if tr1.ObjPtr != nil {
		t.Errorf("obj_ptr expected nil, got %v", *tr1.ObjPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be only 4 errors. %+v", errors)
	} else {
		if errors[0].Attribute != "obj_val" || errors[0].Name != "format" {
			t.Errorf("errors[0] expected attributes, got %v", errors[0])
		}
		if errors[1].Attribute != "obj_ptr" || errors[1].Name != "format" {
			t.Errorf("errors[1] expected attributes, got %v", errors[1])
		}
	}
}

func TestValidateSliceInt(t *testing.T) {
	type testRequest struct {
		IntVal   []int    `json:"int_val"`
		IntNil   []int    `json:"int_nil"`
		IntPtr   []*int   `json:"int_ptr"`
		Int8Val  []int8   `json:"int8_val"`
		Int8Ptr  []*int8  `json:"int8_ptr"`
		Int16Val []int16  `json:"int16_val"`
		Int16Ptr []*int16 `json:"int16_ptr"`
		Int32Val []int32  `json:"int32_val"`
		Int32Ptr []*int32 `json:"int32_ptr"`
		Int64Val []int64  `json:"int64_val"`
		Int64Ptr []*int64 `json:"int64_ptr"`
	}

	var data map[string]any
	r1text := `{"int_val": [1, 2], "int_nil": null, "int_ptr": [2, 3], "int8_val": [3, 4], "int8_ptr": [4, 5], "int16_val": [5, 6], "int16_ptr": [6, 7], "int32_val": [7, 8], "int32_ptr": [8, 9], "int64_val": [9, 10], "int64_ptr": [10, 11]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 2 || tr.IntVal[0] != 1 || tr.IntVal[1] != 2 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if tr.IntNil != nil {
		t.Errorf("int_ptr expected nil, got %v", tr.IntNil)
	}
	if tr.IntPtr == nil || tr.IntPtr[0] == nil || *tr.IntPtr[0] != 2 || tr.IntPtr[1] == nil || *tr.IntPtr[1] != 3 {
		t.Errorf("int_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(tr.Int8Val) != 2 || tr.Int8Val[0] != 3 || tr.Int8Val[1] != 4 {
		t.Errorf("the data does not match. %+v", tr.Int8Val)
	}
	if tr.Int8Ptr == nil || tr.Int8Ptr[0] == nil || *tr.Int8Ptr[0] != 4 || tr.Int8Ptr[1] == nil || *tr.Int8Ptr[1] != 5 {
		t.Errorf("int8_ptr expected nil, got %v", tr.Int8Ptr)
	}
	if len(tr.Int16Val) != 2 || tr.Int16Val[0] != 5 || tr.Int16Val[1] != 6 {
		t.Errorf("the data does not match. %+v", tr.Int16Val)
	}
	if tr.Int16Ptr == nil || tr.Int16Ptr[0] == nil || *tr.Int16Ptr[0] != 6 || tr.Int16Ptr[1] == nil || *tr.Int16Ptr[1] != 7 {
		t.Errorf("int16_ptr expected nil, got %v", tr.Int16Ptr)
	}
	if len(tr.Int32Val) != 2 || tr.Int32Val[0] != 7 || tr.Int32Val[1] != 8 {
		t.Errorf("the data does not match. %+v", tr.Int32Val)
	}
	if tr.Int32Ptr == nil || tr.Int32Ptr[0] == nil || *tr.Int32Ptr[0] != 8 || tr.Int32Ptr[1] == nil || *tr.Int32Ptr[1] != 9 {
		t.Errorf("int32_ptr expected nil, got %v", tr.Int32Ptr)
	}
	if len(tr.Int64Val) != 2 || tr.Int64Val[0] != 9 || tr.Int64Val[1] != 10 {
		t.Errorf("the data does not match. %+v", tr.Int64Val)
	}
	if tr.Int64Ptr == nil || tr.Int64Ptr[0] == nil || *tr.Int64Ptr[0] != 10 || tr.Int64Ptr[1] == nil || *tr.Int64Ptr[1] != 11 {
		t.Errorf("int64_ptr expected nil, got %v", tr.Int64Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"int_val": 6, "int_nil": [1, "a", 3.5], "int_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if len(tr.IntNil) != 3 || tr.IntNil[0] != 1 || tr.IntNil[1] != 0 || tr.IntNil[2] != 3 {
		t.Errorf("the data does not match. %+v", tr.IntNil)
	}
	if tr.IntPtr != nil {
		t.Errorf("int_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(errors) != 3 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceUInt(t *testing.T) {
	type testRequest struct {
		IntVal   []uint    `json:"uint_val"`
		IntNil   []uint    `json:"uint_nil"`
		IntPtr   []*uint   `json:"uint_ptr"`
		Int8Val  []uint8   `json:"uint8_val"`
		Int8Ptr  []*uint8  `json:"uint8_ptr"`
		Int16Val []uint16  `json:"uint16_val"`
		Int16Ptr []*uint16 `json:"uint16_ptr"`
		Int32Val []uint32  `json:"uint32_val"`
		Int32Ptr []*uint32 `json:"uint32_ptr"`
		Int64Val []uint64  `json:"uint64_val"`
		Int64Ptr []*uint64 `json:"uint64_ptr"`
	}

	var data map[string]any
	r1text := `{"uint_val": [1, 2], "uint_nil": null, "uint_ptr": [2, 3], "uint8_val": [3, 4], "uint8_ptr": [4, 5], "uint16_val": [5, 6], "uint16_ptr": [6, 7], "uint32_val": [7, 8], "uint32_ptr": [8, 9], "uint64_val": [9, 10], "uint64_ptr": [10, 11]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 2 || tr.IntVal[0] != 1 || tr.IntVal[1] != 2 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if tr.IntNil != nil {
		t.Errorf("uint_ptr expected nil, got %v", tr.IntNil)
	}
	if tr.IntPtr == nil || tr.IntPtr[0] == nil || *tr.IntPtr[0] != 2 || tr.IntPtr[1] == nil || *tr.IntPtr[1] != 3 {
		t.Errorf("uint_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(tr.Int8Val) != 2 || tr.Int8Val[0] != 3 || tr.Int8Val[1] != 4 {
		t.Errorf("the data does not match. %+v", tr.Int8Val)
	}
	if tr.Int8Ptr == nil || tr.Int8Ptr[0] == nil || *tr.Int8Ptr[0] != 4 || tr.Int8Ptr[1] == nil || *tr.Int8Ptr[1] != 5 {
		t.Errorf("uint8_ptr expected nil, got %v", tr.Int8Ptr)
	}
	if len(tr.Int16Val) != 2 || tr.Int16Val[0] != 5 || tr.Int16Val[1] != 6 {
		t.Errorf("the data does not match. %+v", tr.Int16Val)
	}
	if tr.Int16Ptr == nil || tr.Int16Ptr[0] == nil || *tr.Int16Ptr[0] != 6 || tr.Int16Ptr[1] == nil || *tr.Int16Ptr[1] != 7 {
		t.Errorf("uint16_ptr expected nil, got %v", tr.Int16Ptr)
	}
	if len(tr.Int32Val) != 2 || tr.Int32Val[0] != 7 || tr.Int32Val[1] != 8 {
		t.Errorf("the data does not match. %+v", tr.Int32Val)
	}
	if tr.Int32Ptr == nil || tr.Int32Ptr[0] == nil || *tr.Int32Ptr[0] != 8 || tr.Int32Ptr[1] == nil || *tr.Int32Ptr[1] != 9 {
		t.Errorf("uint32_ptr expected nil, got %v", tr.Int32Ptr)
	}
	if len(tr.Int64Val) != 2 || tr.Int64Val[0] != 9 || tr.Int64Val[1] != 10 {
		t.Errorf("the data does not match. %+v", tr.Int64Val)
	}
	if tr.Int64Ptr == nil || tr.Int64Ptr[0] == nil || *tr.Int64Ptr[0] != 10 || tr.Int64Ptr[1] == nil || *tr.Int64Ptr[1] != 11 {
		t.Errorf("uint64_ptr expected nil, got %v", tr.Int64Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"uint_val": 6, "uint_nil": [1, "a", 3.5], "uint_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if len(tr.IntNil) != 3 || tr.IntNil[0] != 1 || tr.IntNil[1] != 0 || tr.IntNil[2] != 3 {
		t.Errorf("the data does not match. %+v", tr.IntNil)
	}
	if tr.IntPtr != nil {
		t.Errorf("uint_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(errors) != 3 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceByteAndRune(t *testing.T) {
	type testRequest struct {
		ByteVal []byte  `json:"byte_val"`
		BytePtr []*byte `json:"byte_ptr"`
		RuneVal []rune  `json:"rune_val"`
		RunePtr []*rune `json:"rune_ptr"`
	}

	var data map[string]any
	r1text := `{"byte_val": [1, 2], "byte_ptr": [2, 3], "rune_val": [3, 4], "rune_ptr": [4, 5]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.ByteVal) != 2 || tr.ByteVal[0] != 1 || tr.ByteVal[1] != 2 {
		t.Errorf("the data does not match. %+v", tr.ByteVal)
	}
	if tr.BytePtr == nil || tr.BytePtr[0] == nil || *tr.BytePtr[0] != 2 || tr.BytePtr[1] == nil || *tr.BytePtr[1] != 3 {
		t.Errorf("byte_ptr expected nil, got %v", tr.BytePtr)
	}
	if len(tr.RuneVal) != 2 || tr.RuneVal[0] != 3 || tr.RuneVal[1] != 4 {
		t.Errorf("the data does not match. %+v", tr.RuneVal)
	}
	if tr.RunePtr == nil || tr.RunePtr[0] == nil || *tr.RunePtr[0] != 4 || tr.RunePtr[1] == nil || *tr.RunePtr[1] != 5 {
		t.Errorf("rune_ptr expected nil, got %v", tr.RunePtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"byte_val": 6, "byte_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.ByteVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.ByteVal)
	}
	if tr.BytePtr != nil {
		t.Errorf("byte_ptr expected nil, got %v", tr.BytePtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceFloat(t *testing.T) {
	type testRequest struct {
		Float64Val []float64  `json:"float64_val"`
		Float64Ptr []*float64 `json:"float64_ptr"`
		Float32Val []float32  `json:"float32_val"`
		Float32Ptr []*float32 `json:"float32_ptr"`
	}

	var data map[string]any
	r1text := `{"float64_val": [1.2, 2.3], "float64_ptr": [2.4, 3.4], "float32_val": [3.5, 4.5], "float32_ptr": [4.6, 5.6]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.Float64Val) != 2 || tr.Float64Val[0] != 1.2 || tr.Float64Val[1] != 2.3 {
		t.Errorf("the data does not match. %+v", tr.Float64Val)
	}
	if tr.Float64Ptr == nil || tr.Float64Ptr[0] == nil || *tr.Float64Ptr[0] != 2.4 || tr.Float64Ptr[1] == nil || *tr.Float64Ptr[1] != 3.4 {
		t.Errorf("byte_ptr expected nil, got %v", tr.Float64Ptr)
	}
	if len(tr.Float32Val) != 2 || tr.Float32Val[0] != 3.5 || tr.Float32Val[1] != 4.5 {
		t.Errorf("the data does not match. %+v", tr.Float32Val)
	}
	if tr.Float32Ptr == nil || tr.Float32Ptr[0] == nil || *tr.Float32Ptr[0] != 4.6 || tr.Float32Ptr[1] == nil || *tr.Float32Ptr[1] != 5.6 {
		t.Errorf("rune_ptr expected nil, got %v", tr.Float32Ptr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"float64_val": 6, "float64_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.Float64Val) != 0 {
		t.Errorf("the data does not match. %+v", tr.Float64Val)
	}
	if tr.Float64Ptr != nil {
		t.Errorf("byte_ptr expected nil, got %v", tr.Float64Ptr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceBool(t *testing.T) {
	type testRequest struct {
		BoolVal []bool  `json:"bool_val"`
		BoolPtr []*bool `json:"bool_ptr"`
	}

	var data map[string]any
	r1text := `{"bool_val": [true, false], "bool_ptr": [false, true]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.BoolVal) != 2 || tr.BoolVal[0] != true || tr.BoolVal[1] != false {
		t.Errorf("the data does not match. %+v", tr.BoolVal)
	}
	if tr.BoolPtr == nil || tr.BoolPtr[0] == nil || *tr.BoolPtr[0] != false || tr.BoolPtr[1] == nil || *tr.BoolPtr[1] != true {
		t.Errorf("bool_ptr expected nil, got %v", tr.BoolPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"bool_val": 6, "bool_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.BoolVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.BoolVal)
	}
	if tr.BoolPtr != nil {
		t.Errorf("bool_ptr expected nil, got %v", tr.BoolPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceString(t *testing.T) {
	type testRequest struct {
		StringVal []string  `json:"string_val"`
		StringPtr []*string `json:"string_ptr"`
	}

	var data map[string]any
	r1text := `{"string_val": ["a", "b"], "string_ptr": ["c", "d"]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.StringVal) != 2 || tr.StringVal[0] != "a" || tr.StringVal[1] != "b" {
		t.Errorf("the data does not match. %+v", tr.StringVal)
	}
	if tr.StringPtr == nil || tr.StringPtr[0] == nil || *tr.StringPtr[0] != "c" || tr.StringPtr[1] == nil || *tr.StringPtr[1] != "d" {
		t.Errorf("string_ptr expected nil, got %v", tr.StringPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"string_val": 6, "string_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.StringVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.StringVal)
	}
	if tr.StringPtr != nil {
		t.Errorf("string_ptr expected nil, got %v", tr.StringPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceType(t *testing.T) {
	type testIntType int
	const (
		intType0 testIntType = iota
		intType1
		intType2
	)

	type testRequest struct {
		IntVal []testIntType  `json:"int_val"`
		IntPtr []*testIntType `json:"int_ptr"`
	}

	var data map[string]any
	r1text := `{"int_val": [1, 2], "int_ptr": [2, 0]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 2 || tr.IntVal[0] != intType1 || tr.IntVal[1] != intType2 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if tr.IntPtr == nil || tr.IntPtr[0] == nil || *tr.IntPtr[0] != intType2 || tr.IntPtr[1] == nil || *tr.IntPtr[1] != intType0 {
		t.Errorf("int_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"int_val": 6, "int_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.IntVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.IntVal)
	}
	if tr.IntPtr != nil {
		t.Errorf("int_ptr expected nil, got %v", tr.IntPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceStructs(t *testing.T) {
	type testSubStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type testStruct struct {
		ID     int              `json:"id"`
		Name   string           `json:"name"`
		Sub    []testSubStruct  `json:"sub"`
		SubPtr []*testSubStruct `json:"sub_ptr"`
	}

	type testRequest struct {
		ObjVal []testStruct  `json:"obj_val"`
		ObjPtr []*testStruct `json:"obj_ptr"`
	}

	var data map[string]any
	r1text := `{"obj_val": [{"id": 1, "name": "Name1", "sub": [{"id": 2, "name": "Sub1"}, {"id": 3, "name": "Sub2"}]}, {"id": 3, "name": "Name3", "sub": []}], "obj_ptr": [{"id": 2, "name": "Name2", "sub_ptr": [{"id": 4, "name": "Sub4"}, {"id": 5, "name": "Sub5"}]}]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.ObjVal) != 2 || tr.ObjVal[0].ID != 1 || tr.ObjVal[0].Name != "Name1" || len(tr.ObjVal[0].Sub) != 2 || tr.ObjVal[0].Sub[0].ID != 2 || tr.ObjVal[0].Sub[0].Name != "Sub1" || tr.ObjVal[1].ID != 3 {
		t.Errorf("the data does not match. %+v", tr.ObjVal)
	}
	if tr.ObjPtr == nil {
		t.Errorf("obj_ptr expected nil, got %v", tr.ObjPtr)
	} else {
		if len(tr.ObjPtr) != 1 || tr.ObjPtr[0] == nil || (*tr.ObjPtr[0]).ID != 2 || (*tr.ObjPtr[0]).Name != "Name2" || len((*tr.ObjPtr[0]).SubPtr) != 2 || (*(*tr.ObjPtr[0]).SubPtr[0]).ID != 4 {
			t.Errorf("obj_ptr expected nil, got %#v", tr.ObjPtr[0])
		}
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"obj_val": 6, "obj_ptr": "aa"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(tr.ObjVal) != 0 {
		t.Errorf("the data does not match. %+v", tr.ObjVal)
	}
	if tr.ObjPtr != nil {
		t.Errorf("obj_ptr expected nil, got %v", tr.ObjPtr)
	}
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateSliceSlice(t *testing.T) {
	type testRequest struct {
		StringVal [][]string  `json:"string_val"`
		StringPtr [][]*string `json:"string_ptr"`
	}

	var data map[string]any
	r1text := `{"string_val": [["a", "1"], ["b"]], "string_ptr": [["c"], ["d", "4"]]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.StringVal) != 2 || len(tr.StringVal[0]) != 2 || tr.StringVal[0][0] != "a" || tr.StringVal[0][1] != "1" || len(tr.StringVal[1]) != 1 || tr.StringVal[1][0] != "b" {
		t.Errorf("the data does not match. %+v", tr.StringVal)
	}
	if len(tr.StringPtr) != 2 || len(tr.StringPtr[0]) != 1 || *tr.StringPtr[0][0] != "c" || len(tr.StringPtr[1]) != 2 || *tr.StringPtr[1][0] != "d" || *tr.StringPtr[1][1] != "4" {
		t.Errorf("string_ptr expected nil, got %v", tr.StringPtr)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}

	data = make(map[string]any)
	r1text = `{"string_val": [6], "string_ptr": ["aa"]}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors = Run[testRequest](data, map[string][]any{})
	if len(errors) != 2 {
		t.Errorf("there should be 2 errors. %+v", errors)
	}
}

func TestValidateMap(t *testing.T) {
	type testRequest struct {
		Map1 map[string]int   `json:"map1"`
		Map2 map[string]*int  `json:"map2"`
		Map3 map[string][]int `json:"map3"`
	}

	var data map[string]any
	r1text := `{"map1": {"1": 2, "4": 5}, "map2": {"d":11, "e":12, "f":13}, "map3": {"k":[1, 2, 3]}}`
	_ = json.Unmarshal([]byte(r1text), &data)

	tr, errors := Run[testRequest](data, map[string][]any{})
	if len(tr.Map1) != 2 || tr.Map1["1"] != 2 || tr.Map1["4"] != 5 {
		t.Errorf("the data does not match. %+v", tr.Map1)
	}
	if len(tr.Map2) != 3 || *tr.Map2["d"] != 11 || *tr.Map2["e"] != 12 || *tr.Map2["f"] != 13 {
		t.Errorf("the data does not match. %+v", tr.Map2)
	}
	if len(tr.Map3) != 1 || len(tr.Map3["k"]) != 3 || tr.Map3["k"][0] != 1 || tr.Map3["k"][1] != 2 || tr.Map3["k"][2] != 3 {
		t.Errorf("the data does not match. %+v", tr.Map3)
	}
	if len(errors) > 0 {
		t.Errorf("there were errors when converting the data. %+v", errors)
	}
}
