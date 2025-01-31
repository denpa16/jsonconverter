package main

import (
	"fmt"
	"jsonconverter/field_setter"
)

type TestChildStructData struct {
	ValueString         string                 `json:"valuestring"`
	ValueInt            int                    `json:"valueint"`
	ValueInt8           int8                   `json:"valueint8"`
	ValueInt16          int16                  `json:"valueint16"`
	ValueInt32          int32                  `json:"valueint32"`
	ValueInt64          int64                  `json:"valueint64"`
	ValueFloat32        float32                `json:"valuefloat32"`
	ValueFloat64        float64                `json:"valuefloat64"`
	Active              bool                   `json:"active"`
	ValueSliceString    []string               `json:"valueslicestring"`
	ValueSliceInt       []int                  `json:"valuesliceint"`
	ValueSliceInt8      []int8                 `json:"valuesliceint8"`
	ValueSliceInt16     []int16                `json:"valuesliceint16"`
	ValueSliceInt32     []int32                `json:"valuesliceint32"`
	ValueSliceInt64     []int64                `json:"valuesliceint64"`
	ValueSliceFloat32   []float32              `json:"valueslicefloat32"`
	ValueSliceFloat64   []float64              `json:"valueslicefloat64"`
	ValueSliceInterface []interface{}          `json:"valuesliceinterface"`
	ValueMapStringKey   map[string]interface{} `json:"valuemapstringkey"`

	SetValueString         bool `json:"-"`
	SetValueInt            bool `json:"-"`
	SetValueInt8           bool `json:"-"`
	SetValueInt16          bool `json:"-"`
	SetValueInt32          bool `json:"-"`
	SetValueInt64          bool `json:"-"`
	SetValueFloat32        bool `json:"-"`
	SetValueFloat64        bool `json:"-"`
	SetActive              bool `json:"-"`
	SetValueSliceString    bool `json:"-"`
	SetValueSliceInt       bool `json:"-"`
	SetValueSliceInt8      bool `json:"-"`
	SetValueSliceInt16     bool `json:"-"`
	SetValueSliceInt32     bool `json:"-"`
	SetValueSliceInt64     bool `json:"-"`
	SetValueSliceFloat32   bool `json:"-"`
	SetValueSliceFloat64   bool `json:"-"`
	SetValueSliceInterface bool `json:"-"`
	SetValueMapStringKey   bool `json:"-"`
}

type TestChildPtrStructData struct {
	ValueString         *string                   `json:"valuestring"`
	ValueInt            *int                      `json:"valueint"`
	ValueInt8           *int8                     `json:"valueint8"`
	ValueInt16          *int16                    `json:"valueint16"`
	ValueInt32          *int32                    `json:"valueint32"`
	ValueInt64          *int64                    `json:"valueint64"`
	ValueFloat32        *float32                  `json:"valuefloat32"`
	ValueFloat64        *float64                  `json:"valuefloat64"`
	Active              *bool                     `json:"active"`
	ValueSliceString    *[]string                 `json:"valueslicestring"`
	ValueSliceInt       *[]int                    `json:"valuesliceint"`
	ValueSliceInt8      *[]int8                   `json:"valuesliceint8"`
	ValueSliceInt16     *[]int16                  `json:"valuesliceint16"`
	ValueSliceInt32     *[]int32                  `json:"valuesliceint32"`
	ValueSliceInt64     *[]int64                  `json:"valuesliceint64"`
	ValueSliceFloat32   *[]float32                `json:"valueslicefloat32"`
	ValueSliceFloat64   *[]float64                `json:"valueslicefloat64"`
	ValueSliceInterface *[]interface{}            `json:"valuesliceinterface"`
	ValueMapStringKey   *map[string]interface{}   `json:"valuemapstringkey"`
	ValueStruct         *TestChildStructData      `json:"valuestruct"`
	ValuePtrStruct      *TestChildPtrStructData   `json:"valueptrstruct"`
	ValueSliceStruct    *[]TestChildStructData    `json:"valueslicestruct"`
	ValueSlicePtrStruct *[]TestChildPtrStructData `json:"valuesliceptrstruct"`

	SetValueString         bool `json:"-"`
	SetValueInt            bool `json:"-"`
	SetValueInt8           bool `json:"-"`
	SetValueInt16          bool `json:"-"`
	SetValueInt32          bool `json:"-"`
	SetValueInt64          bool `json:"-"`
	SetValueFloat32        bool `json:"-"`
	SetValueFloat64        bool `json:"-"`
	SetActive              bool `json:"-"`
	SetValueSliceString    bool `json:"-"`
	SetValueSliceInt       bool `json:"-"`
	SetValueSliceInt8      bool `json:"-"`
	SetValueSliceInt16     bool `json:"-"`
	SetValueSliceInt32     bool `json:"-"`
	SetValueSliceInt64     bool `json:"-"`
	SetValueSliceFloat32   bool `json:"-"`
	SetValueSliceFloat64   bool `json:"-"`
	SetValueSliceInterface bool `json:"-"`
	SetValueMapStringKey   bool `json:"-"`
	SetValueStruct         bool `json:"-"`
	SetValuePtrStruct      bool `json:"-"`
	SetValueSliceStruct    bool `json:"-"`
	SetValueSlicePtrStruct bool `json:"-"`
}

type TestPtrStructData struct {
	ValueString         *string                   `json:"valuestring"`
	ValueInt            *int                      `json:"valueint"`
	ValueInt8           *int8                     `json:"valueint8"`
	ValueInt16          *int16                    `json:"valueint16"`
	ValueInt32          *int32                    `json:"valueint32"`
	ValueInt64          *int64                    `json:"valueint64"`
	ValueFloat32        *float32                  `json:"valuefloat32"`
	ValueFloat64        *float64                  `json:"valuefloat64"`
	Active              *bool                     `json:"active"`
	ValueSliceString    *[]string                 `json:"valueslicestring"`
	ValueSliceInt       *[]int                    `json:"valuesliceint"`
	ValueSliceInt8      *[]int8                   `json:"valuesliceint8"`
	ValueSliceInt16     *[]int16                  `json:"valuesliceint16"`
	ValueSliceInt32     *[]int32                  `json:"valuesliceint32"`
	ValueSliceInt64     *[]int64                  `json:"valuesliceint64"`
	ValueSliceFloat32   *[]float32                `json:"valueslicefloat32"`
	ValueSliceFloat64   *[]float64                `json:"valueslicefloat64"`
	ValueSliceInterface *[]interface{}            `json:"valuesliceinterface"`
	ValueMapStringKey   *map[string]interface{}   `json:"valuemapstringkey"`
	ValueStruct         *TestChildStructData      `json:"valuestruct"`
	ValuePtrStruct      *TestChildPtrStructData   `json:"valueptrstruct"`
	ValueSliceStruct    *[]TestChildStructData    `json:"valueslicestruct"`
	ValueSlicePtrStruct *[]TestChildPtrStructData `json:"valuesliceptrstruct"`

	SetValueString         bool `json:"-"`
	SetValueInt            bool `json:"-"`
	SetValueInt8           bool `json:"-"`
	SetValueInt16          bool `json:"-"`
	SetValueInt32          bool `json:"-"`
	SetValueInt64          bool `json:"-"`
	SetValueFloat32        bool `json:"-"`
	SetValueFloat64        bool `json:"-"`
	SetActive              bool `json:"-"`
	SetValueSliceString    bool `json:"-"`
	SetValueSliceInt       bool `json:"-"`
	SetValueSliceInt8      bool `json:"-"`
	SetValueSliceInt16     bool `json:"-"`
	SetValueSliceInt32     bool `json:"-"`
	SetValueSliceInt64     bool `json:"-"`
	SetValueSliceFloat32   bool `json:"-"`
	SetValueSliceFloat64   bool `json:"-"`
	SetValueSliceInterface bool `json:"-"`
	SetValueMapStringKey   bool `json:"-"`
	SetValueStruct         bool `json:"-"`
	SetValuePtrStruct      bool `json:"-"`
	SetValueSliceStruct    bool `json:"-"`
	SetValueSlicePtrStruct bool `json:"-"`
}

type TestStructData struct {
	ValueString         string                   `json:"valuestring"`
	ValueInt            int                      `json:"valueint"`
	ValueInt8           int8                     `json:"valueint8"`
	ValueInt16          int16                    `json:"valueint16"`
	ValueInt32          int32                    `json:"valueint32"`
	ValueInt64          int64                    `json:"valueint64"`
	ValueFloat32        float32                  `json:"valuefloat32"`
	ValueFloat64        float64                  `json:"valuefloat64"`
	Active              bool                     `json:"active"`
	ValueSliceString    []string                 `json:"valueslicestring"`
	ValueSliceInt       []int                    `json:"valuesliceint"`
	ValueSliceInt8      []int8                   `json:"valuesliceint8"`
	ValueSliceInt16     []int16                  `json:"valuesliceint16"`
	ValueSliceInt32     []int32                  `json:"valuesliceint32"`
	ValueSliceInt64     []int64                  `json:"valuesliceint64"`
	ValueSliceFloat32   []float32                `json:"valueslicefloat32"`
	ValueSliceFloat64   []float64                `json:"valueslicefloat64"`
	ValueSliceInterface []interface{}            `json:"valuesliceinterface"`
	ValueMapStringKey   map[string]interface{}   `json:"valuemapstringkey"`
	ValueStruct         TestChildStructData      `json:"valuestruct"`
	ValueSliceStruct    []TestChildStructData    `json:"valueslicestruct"`
	ValueSlicePtrStruct []TestChildPtrStructData `json:"valuesliceptrstruct"`

	SetValueString         bool `json:"-"`
	SetValueInt            bool `json:"-"`
	SetValueInt8           bool `json:"-"`
	SetValueInt16          bool `json:"-"`
	SetValueInt32          bool `json:"-"`
	SetValueInt64          bool `json:"-"`
	SetValueFloat32        bool `json:"-"`
	SetValueFloat64        bool `json:"-"`
	SetActive              bool `json:"-"`
	SetValueSliceString    bool `json:"-"`
	SetValueSliceInt       bool `json:"-"`
	SetValueSliceInt8      bool `json:"-"`
	SetValueSliceInt16     bool `json:"-"`
	SetValueSliceInt32     bool `json:"-"`
	SetValueSliceInt64     bool `json:"-"`
	SetValueSliceFloat32   bool `json:"-"`
	SetValueSliceFloat64   bool `json:"-"`
	SetValueSliceInterface bool `json:"-"`
	SetValueMapStringKey   bool `json:"-"`
	SetValueStruct         bool `json:"-"`
	SetValueSliceStruct    bool `json:"-"`
	SetValueSlicePtrStruct bool `json:"-"`
}

func main() {
	jsonData := []byte(`{
		"valuestring": "test",
		"valueint8": 8,
		"valueint16": 16,
		"valueint32": 32,
		"valueint64": 64,
		"valuefloat32": 3.14,
		"valuefloat64": 6.28,
		"active": true,
		"valueslicestring": ["a", "b", "c"],
		"valuesliceint": [1, 2, 3],
		"valuesliceint8": [4, 5, 6],
		"valuesliceint16": [7, 8, 9],
		"valuesliceint32": [10, 11, 12],
		"valuesliceint64": [13, 14, 15],
		"valueslicefloat32": [1.1, 2.2, 3.3],
		"valueslicefloat64": [4.4, 5.5, 6.6],
		"valuesliceinterface": [1, "two", 3.0],
		"valuemapstringkey": {
			"key1": "value1",
			"key2": 2,
			"key3": 3.0
		},
		"valuestruct": {
			"valuestring": "nested",
			"valueint": 99
		},
		"valueslicestruct": [
			{
				"valuestring": "nested1",
				"valueint": 100
			},
			{
				"valuestring": "nested2",
				"valueint": 200
			}
		],
		"valuesliceptrstruct": [
			{
				"valuestring": "nestedPtr1",
				"valueint": 300
			},
			{
				"valuestring": "nestedPtr2",
				"valueint": 400
			}
		]
	}`)

	testPtrStruct := &TestPtrStructData{}

	field_setter.FillFields(testPtrStruct, jsonData)
	fmt.Println(*testPtrStruct.ValueStruct)
	fmt.Println(testPtrStruct.SetValueString)
}
