PopulateFieldsAndSets util for convert json to Struct
=====================

Overview
--------

PopulateFieldsAndSets is a Go function designed to populate the fields of a struct with values from a JSON object. It dynamically checks whether fields with corresponding JSON tags exist in the input JSON and assigns values to the struct fields, including nested structs and slices. Additionally, the function tracks the status of each field, setting a boolean flag indicating whether a field was populated from the JSON input.

Function Signature
------------------

Parameters
----------

*   jsonData (\[\]byte): A byte slice containing the JSON data to be parsed and used to populate the struct fields.
    
*   result (interface{}): A pointer to the struct that will have its fields populated based on the JSON data. The struct must have fields with JSON tags for the mapping to work.
    

Return Value
------------

*   error: The function returns an error if any issues arise while unmarshaling the JSON data or populating the struct fields. If the operation is successful, nil is returned.
    

Behavior
--------

1.  **Parsing JSON**:The function first unmarshals the provided jsonData into a map\[string\]interface{}. This allows the function to check if each field in the target struct matches a key in the JSON data.
    
2.  **Field Population**:The function then uses reflection to iterate over the fields of the provided struct. For each field, it checks if a corresponding key exists in the JSON data (using the JSON tag associated with the field). If a match is found, the field is populated with the corresponding value from the JSON. The following field types are supported:
    
    *   Basic types like string, int, float64, bool.
        
    *   Pointers to these types.
        
    *   Slices and arrays of these types.
        
    *   Maps.
        
    *   Nested structs (recursively populated).
        
3.  **Set Flags**:For each field, the function sets a flag named Set to true if the field was successfully populated. If a field is not present in the JSON data, the corresponding Set flag is set to false.
    
4.  **Type Conversion**:The function handles type conversions for numeric fields (e.g., int, int8, int16, int32, int64, float32, float64) to ensure that the correct type is assigned, even when the values are initially in a different type (e.g., float64 from JSON).
    

Example Usage
-------------
```
package main

import (
	"encoding/json"
	"fmt"
)

type TestChildStructData struct {
	ValueString string `json:"valuestring"`
}

type TestChildPtrStructData struct {
	ValueString *string `json:"valuestring"`
}

type TestPtrStructData struct {
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
	ValueStruct         *TestChildStructData    `json:"valuestruct"`
	ValuePtrStruct      *TestChildPtrStructData `json:"valueptrstruct"`

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


```

### Output:

Key Points
----------

*   **JSON Mapping**: The struct fields are mapped to the JSON fields using the struct tags (e.g., json:"name").
    
*   **Automatic Field Population**: The function automatically populates the fields with values from the JSON data.
    
*   **Flagging Fields**: The function creates flags to indicate whether each field was populated, which can be useful for further processing.
    
*   **Support for Nested Types**: The function supports deeply nested structs and arrays within structs, handling them recursively.
    

Error Handling
--------------

If the JSON data doesn't match the expected type (for example, trying to assign a string to an int field), or if there's an issue unmarshaling the JSON data, the function will return an error. This error can be handled by the caller to manage exceptions and ensure data integrity.

Limitations
-----------

*   The function assumes that the JSON data provided matches the structure and type of the target struct (other than simple type conversions). Complex or mismatched structures may cause runtime errors.
    
*   If the JSON data does not provide a field for a struct's property, the Set flag will be set to false for that field.