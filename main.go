package main

import (
	"encoding/json"
	"fmt"
)

type ChildStructData struct {
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

type ChildPtrStructData struct {
	ValueString         *string                 `json:"valuestring"`
	ValueInt            *int                    `json:"valueint"`
	ValueInt8           *int8                   `json:"valueint8"`
	ValueInt16          *int16                  `json:"valueint16"`
	ValueInt32          *int32                  `json:"valueint32"`
	ValueInt64          *int64                  `json:"valueint64"`
	ValueFloat32        *float32                `json:"valuefloat32"`
	ValueFloat64        *float64                `json:"valuefloat64"`
	Active              *bool                   `json:"active"`
	ValueSliceString    *[]string               `json:"valueslicestring"`
	ValueSliceInt       *[]int                  `json:"valuesliceint"`
	ValueSliceInt8      *[]int8                 `json:"valuesliceint8"`
	ValueSliceInt16     *[]int16                `json:"valuesliceint16"`
	ValueSliceInt32     *[]int32                `json:"valuesliceint32"`
	ValueSliceInt64     *[]int64                `json:"valuesliceint64"`
	ValueSliceFloat32   *[]float32              `json:"valueslicefloat32"`
	ValueSliceFloat64   *[]float64              `json:"valueslicefloat64"`
	ValueSliceInterface *[]interface{}          `json:"valuesliceinterface"`
	ValueMapStringKey   *map[string]interface{} `json:"valuemapstringkey"`
	ValueStruct         *ChildStructData        `json:"valuestruct"`
	ValuePtrStruct      *ChildPtrStructData     `json:"valueptrstruct"`
	ValueSliceStruct    *[]ChildStructData      `json:"valueslicestruct"`
	ValueSlicePtrStruct *[]ChildPtrStructData   `json:"valuesliceptrstruct"`

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

type PtrStructData struct {
	ValueString         *string                 `json:"valuestring"`
	ValueInt            *int                    `json:"valueint"`
	ValueInt8           *int8                   `json:"valueint8"`
	ValueInt16          *int16                  `json:"valueint16"`
	ValueInt32          *int32                  `json:"valueint32"`
	ValueInt64          *int64                  `json:"valueint64"`
	ValueFloat32        *float32                `json:"valuefloat32"`
	ValueFloat64        *float64                `json:"valuefloat64"`
	Active              *bool                   `json:"active"`
	ValueSliceString    *[]string               `json:"valueslicestring"`
	ValueSliceInt       *[]int                  `json:"valuesliceint"`
	ValueSliceInt8      *[]int8                 `json:"valuesliceint8"`
	ValueSliceInt16     *[]int16                `json:"valuesliceint16"`
	ValueSliceInt32     *[]int32                `json:"valuesliceint32"`
	ValueSliceInt64     *[]int64                `json:"valuesliceint64"`
	ValueSliceFloat32   *[]float32              `json:"valueslicefloat32"`
	ValueSliceFloat64   *[]float64              `json:"valueslicefloat64"`
	ValueSliceInterface *[]interface{}          `json:"valuesliceinterface"`
	ValueMapStringKey   *map[string]interface{} `json:"valuemapstringkey"`
	ValueStruct         *ChildStructData        `json:"valuestruct"`
	ValuePtrStruct      *ChildPtrStructData     `json:"valueptrstruct"`
	ValueSliceStruct    *[]ChildStructData      `json:"valueslicestruct"`
	ValueSlicePtrStruct *[]ChildPtrStructData   `json:"valuesliceptrstruct"`

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

type StructData struct {
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
	ValueStruct         ChildStructData        `json:"valuestruct"`
	ValueSliceStruct    []ChildStructData      `json:"valueslicestruct"`
	ValueSlicePtrStruct []ChildPtrStructData   `json:"valuesliceptrstruct"`
	ValueJsonRaw        *json.RawMessage       `json:"valuejsonraw"`

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
	SetValueJsonRaw        bool `json:"-"`
}

func main() {
	// Пример JSON с различными типами данных
	data := `{
		"valuestring": "example_string",
		"valueint": 42,
		"valueint8": 8,
		"valueint16": 16,
		"valueint32": 32,
		"valueint64": 64,
		"valuefloat32": 32.32,
		"valuefloat64": 64.64,
		"active": true,
		"valueslicestring": ["string1", "string2"],
		"valuesliceint": [1, 2, 3],
		"valuesliceint8": [8, 7, 6],
		"valuesliceint16": [16, 15],
		"valuesliceint32": [32, 31],
		"valuesliceint64": [64, 63],
		"valueslicefloat32": [1.1, 2.2],
		"valueslicefloat64": [3.3, 4.4],
		"valuesliceinterface": ["a", 1, 3.14, true],
		"valuemapstringkey": {
		  "key1": "value1",
		  "key2": 2
		},
		"valuestruct": {
		  "valuestring": "child_struct_string",
		  "valueint": 100,
		  "valueint8": 1,
		  "valueint16": 16,
		  "valueint32": 32,
		  "valueint64": 64,
		  "valuefloat32": 10.10,
		  "valuefloat64": 20.20,
		  "active": true,
		  "valueslicestring": ["child_string1"],
		  "valuesliceint": [10],
		  "valuesliceint8": [8],
		  "valuesliceint16": [16],
		  "valuesliceint32": [32],
		  "valuesliceint64": [64],
		  "valueslicefloat32": [3.14],
		  "valueslicefloat64": [2.71],
		  "valuesliceinterface": ["child_interface", 42],
		  "valuemapstringkey": {
			"child_key1": "child_value1"
		  }
		},
		"valueslicestruct": [
		  {
			"valuestring": "struct_in_slice_1",
			"valueint": 500,
			"valueint8": 2,
			"valueint16": 100,
			"valueint32": 200,
			"valueint64": 400,
			"valuefloat32": 50.50,
			"valuefloat64": 100.100,
			"active": false,
			"valueslicestring": ["child_string_in_slice"],
			"valuesliceint": [200],
			"valuesliceint8": [15],
			"valuesliceint16": [50],
			"valuesliceint32": [100],
			"valuesliceint64": [200],
			"valueslicefloat32": [5.55],
			"valueslicefloat64": [9.99],
			"valuesliceinterface": ["another_item", 30],
			"valuemapstringkey": {
			  "child_key_in_slice": "child_value_in_slice"
			}
		  },
		  {
			"valuestring": "struct_in_slice_2",
			"valueint": 600,
			"valueint8": 3,
			"valueint16": 200,
			"valueint32": 400,
			"valueint64": 800,
			"valuefloat32": 60.60,
			"valuefloat64": 120.120,
			"active": true,
			"valueslicestring": ["child_string_in_slice_2"],
			"valuesliceint": [300],
			"valuesliceint8": [18],
			"valuesliceint16": [60],
			"valuesliceint32": [150],
			"valuesliceint64": [250],
			"valueslicefloat32": [6.66],
			"valueslicefloat64": [10.10],
			"valuesliceinterface": ["another_item_2", 100],
			"valuemapstringkey": {
			  "child_key_in_slice_2": "child_value_in_slice_2"
			}
		  }
		],
		"valuesliceptrstruct": [
		  {
			"valuestring": "ptr_struct_in_slice_1",
			"valueint": 700,
			"valueint8": 4,
			"valueint16": 400,
			"valueint32": 800,
			"valueint64": 1600,
			"valuefloat32": 70.70,
			"valuefloat64": 140.140,
			"active": false,
			"valueslicestring": ["ptr_string_in_slice_1"],
			"valuesliceint": [400],
			"valuesliceint8": [20],
			"valuesliceint16": [80],
			"valuesliceint32": [200],
			"valuesliceint64": [400],
			"valueslicefloat32": [7.77],
			"valueslicefloat64": [11.11],
			"valuesliceinterface": ["ptr_item_1", 200],
			"valuemapstringkey": {
			  "ptr_key_in_slice_1": "ptr_value_in_slice_1"
			}
		  }
		]
	  }`

	result1 := &StructData{}
	result2 := &PtrStructData{}

	testData := []byte(data)

	err := PopulateFieldsAndSets(testData, result1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = PopulateFieldsAndSets(testData, result2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result3 := &StructData{}

	ddata := `{
		"valuejsonraw": {"id": 1}
	}`

	testDdata := []byte(ddata)

	err = PopulateFieldsAndSets(testDdata, result3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("###########")
	fmt.Println(result3.ValueJsonRaw)
	fmt.Println(result3.SetValueJsonRaw)
	fmt.Println("###########")

}
