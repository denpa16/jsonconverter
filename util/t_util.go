package util

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestTestPtrStructDataWithoutFilledFields(t *testing.T) {
	// 1. Тест: JSON без поля ValueString (string)
	t.Run("Field with string type, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueString)
		assert.False(t, result.SetValueString)
	})

	// 2. Тест: JSON без поля ValueInt (int)
	t.Run("Field with pointer to int, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueInt)
		assert.False(t, result.SetValueInt)
	})

	// 3. Тест: JSON без поля ValueInt8 (int8)
	t.Run("Field with pointer to int8, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueInt8)
		assert.False(t, result.SetValueInt8)
	})

	// 4. Тест: JSON без поля ValueInt16 (int16)
	t.Run("Field with pointer to int16, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueInt16)
		assert.False(t, result.SetValueInt16)
	})

	// 5. Тест: JSON без поля ValueInt32 (int32)
	t.Run("Field with pointer to int32, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueInt32)
		assert.False(t, result.SetValueInt32)
	})

	// 6. Тест: JSON без поля ValueInt64 (int64)
	t.Run("Field with pointer to int64, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueInt64)
		assert.False(t, result.SetValueInt64)
	})

	// 7. Тест: JSON без поля ValueFloat64 (float64)
	t.Run("Field with float64, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueFloat64)
		assert.False(t, result.SetValueFloat64)
	})

	// 8. Тест: JSON без поля Active (bool)
	t.Run("Field with pointer to bool, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.Active)
		assert.False(t, result.SetActive)
	})

	// 9. Тест: JSON без поля ValueSliceString (срез строк)
	t.Run("Field with pointer to []string, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceString)
		assert.False(t, result.SetValueSliceString)
	})

	// 10. Тест: JSON без поля ValueSliceInt (срез int)
	t.Run("Field with pointer to []int, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInt)
		assert.False(t, result.SetValueSliceInt)
	})

	// 11. Тест: JSON без поля ValueSliceInt8 (срез int8)
	t.Run("Field with pointer to []int8, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInt8)
		assert.False(t, result.SetValueSliceInt8)
	})

	// 12. Тест: JSON без поля ValueSliceInt16 (срез int16)
	t.Run("Field with pointer to []int16, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInt16)
		assert.False(t, result.SetValueSliceInt16)
	})

	// 13. Тест: JSON без поля ValueSliceInt32 (срез int32)
	t.Run("Field with pointer to []int32, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInt32)
		assert.False(t, result.SetValueSliceInt32)
	})

	// 14. Тест: JSON без поля ValueSliceInt64 (срез int64)
	t.Run("Field with pointer to []int64, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInt64)
		assert.False(t, result.SetValueSliceInt64)
	})

	// 15. Тест: JSON без поля ValueSliceInterface (срез интерфейсов)
	t.Run("Field with pointer to []interface{}, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueSliceInterface)
		assert.False(t, result.SetValueSliceInterface)
	})

	// 16. Тест: JSON без поля ValueMapStringKey (map[string]interface{})
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueMapStringKey)
		assert.False(t, result.SetValueMapStringKey)
	})

	// 17. Тест: JSON без поля ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueStruct)
		assert.False(t, result.SetValueStruct)
	})

	// 18. Тест: JSON без поля ValuePtrStruct (TestChildStructData)
	t.Run("Field with pointer to ValuePtrStruct, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValuePtrStruct)
		assert.False(t, result.SetValuePtrStruct)
	})

	// 18. Тест: JSON с заполненным полем ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueStruct)
		assert.False(t, result.SetValueStruct)
	})

	// 19. Тест: JSON с заполненным полем ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValuePtrStruct)
		assert.False(t, result.SetValuePtrStruct)
	})

	// 19. Тест: JSON с пустым объектом, без значений
	t.Run("Field with missing values (empty object)", func(t *testing.T) {
		testJsonData := []byte(`{}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Nil(t, result.ValueString)
		assert.Nil(t, result.ValueInt)
		assert.Nil(t, result.ValueInt8)
		assert.Nil(t, result.ValueInt16)
		assert.Nil(t, result.ValueInt32)
		assert.Nil(t, result.ValueInt64)
		assert.Nil(t, result.ValueFloat64)
		assert.Nil(t, result.Active)
		assert.Nil(t, result.ValueSliceString)
		assert.Nil(t, result.ValueSliceInt)
		assert.Nil(t, result.ValueSliceInt8)
		assert.Nil(t, result.ValueSliceInt16)
		assert.Nil(t, result.ValueSliceInt32)
		assert.Nil(t, result.ValueSliceInt64)
		assert.Nil(t, result.ValueSliceInterface)
		assert.Nil(t, result.ValueMapStringKey)
		assert.Nil(t, result.ValueStruct)

		// Проверяем, что флаги Set не установлены
		assert.False(t, result.SetValueString)
		assert.False(t, result.SetValueInt)
		assert.False(t, result.SetValueInt8)
		assert.False(t, result.SetValueInt16)
		assert.False(t, result.SetValueInt32)
		assert.False(t, result.SetValueInt64)
		assert.False(t, result.SetValueFloat64)
		assert.False(t, result.SetActive)
		assert.False(t, result.SetValueSliceString)
		assert.False(t, result.SetValueSliceInt)
		assert.False(t, result.SetValueSliceInt8)
		assert.False(t, result.SetValueSliceInt16)
		assert.False(t, result.SetValueSliceInt32)
		assert.False(t, result.SetValueSliceInt64)
		assert.False(t, result.SetValueSliceInterface)
		assert.False(t, result.SetValueMapStringKey)
		assert.False(t, result.SetValueStruct)
	})
}

func TestTestPtrStructDataWithFilledFields(t *testing.T) {
	// 1. Тест: JSON с заполненным полем ValueString (string)
	t.Run("Field with string type, filled field", func(t *testing.T) {
		ValueString := "John Doe"
		testJsonData := []byte(`{"valuestring": "` + ValueString + `"}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueString, result.ValueString)
		assert.True(t, result.SetValueString)
	})

	// 2. Тест: JSON с заполненным полем ValueInt (int)
	t.Run("Field with pointer to int, filled field", func(t *testing.T) {
		ValueInt := 30
		testJsonData := []byte(`{"valueint": ` + fmt.Sprintf("%d", ValueInt) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueInt, result.ValueInt)
		assert.True(t, result.SetValueInt)
	})

	// 3. Тест: JSON с заполненным полем ValueInt8 (int8)
	t.Run("Field with pointer to int8, filled field", func(t *testing.T) {
		ValueInt8 := int8(25)
		testJsonData := []byte(`{"valueint8": ` + fmt.Sprintf("%d", ValueInt8) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueInt8, result.ValueInt8)
		assert.True(t, result.SetValueInt8)
	})

	// 4. Тест: JSON с заполненным полем ValueInt16 (int16)
	t.Run("Field with pointer to int16, filled field", func(t *testing.T) {
		ValueInt16 := int16(40)
		testJsonData := []byte(`{"valueint16": ` + fmt.Sprintf("%d", ValueInt16) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueInt16, result.ValueInt16)
		assert.True(t, result.SetValueInt16)
	})

	// 5. Тест: JSON с заполненным полем ValueInt32 (int32)
	t.Run("Field with pointer to int32, filled field", func(t *testing.T) {
		ValueInt32 := int32(35)
		testJsonData := []byte(`{"valueint32": ` + fmt.Sprintf("%d", ValueInt32) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueInt32, result.ValueInt32)
		assert.True(t, result.SetValueInt32)
	})

	// 6. Тест: JSON с заполненным полем ValueInt64 (int64)
	t.Run("Field with pointer to int64, filled field", func(t *testing.T) {
		ValueInt64 := int64(50)
		testJsonData := []byte(`{"valueint64": ` + fmt.Sprintf("%d", ValueInt64) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueInt64, result.ValueInt64)
		assert.True(t, result.SetValueInt64)
	})

	// 7. Тест: JSON с заполненным полем ValueFloat64 (float64)
	t.Run("Field with float64, filled field", func(t *testing.T) {
		ValueFloat64 := 175.5
		testJsonData := []byte(`{"valuefloat64": ` + fmt.Sprintf("%f", ValueFloat64) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueFloat64, result.ValueFloat64)
		assert.True(t, result.SetValueFloat64)
	})

	// 8. Тест: JSON с заполненным полем Active (bool)
	t.Run("Field with pointer to bool, filled field", func(t *testing.T) {
		active := true
		testJsonData := []byte(`{"active": ` + strconv.FormatBool(active) + `}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &active, result.Active)
		assert.True(t, result.SetActive)
	})

	// 9. Тест: JSON с заполненным полем ValueSliceString (срез строк)
	t.Run("Field with pointer to []string, filled field", func(t *testing.T) {
		ValueSliceString := []string{"tag1", "tag2", "tag3"}
		testJsonData := []byte(`{"valueslicestring": ["tag1", "tag2", "tag3"]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceString, result.ValueSliceString)
		assert.True(t, result.SetValueSliceString)
	})

	// 10. Тест: JSON с заполненным полем ValueSliceInt (срез int)
	t.Run("Field with pointer to []int, filled field", func(t *testing.T) {
		ValueSliceInt := []int{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint": [1, 2, 3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceInt)
		assert.True(t, result.SetValueSliceInt)
	})

	// 11. Тест: JSON с заполненным полем ValueSliceInt8 (срез int8)
	t.Run("Field with pointer to []int8, filled field", func(t *testing.T) {
		ValueSliceInt := []int8{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint8": [1, 2, 3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceInt8)
		assert.True(t, result.SetValueSliceInt8)
	})

	// 12. Тест: JSON с заполненным полем ValueSliceInt16 (срез int16)
	t.Run("Field with pointer to []int16, filled field", func(t *testing.T) {
		ValueSliceInt := []int16{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint16": [1, 2, 3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceInt16)
		assert.True(t, result.SetValueSliceInt16)
	})

	// 13. Тест: JSON с заполненным полем ValueSliceInt32 (срез int32)
	t.Run("Field with pointer to []int32, filled field", func(t *testing.T) {
		ValueSliceInt := []int32{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint32": [1, 2, 3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceInt32)
		assert.True(t, result.SetValueSliceInt32)
	})

	// 14. Тест: JSON с заполненным полем ValueSliceInt64 (срез int64)
	t.Run("Field with pointer to []int64, filled field", func(t *testing.T) {
		ValueSliceInt := []int64{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint64": [1, 2, 3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceInt64)
		assert.True(t, result.SetValueSliceInt64)
	})

	// 15. Тест: JSON с заполненным полем ValueSliceFloat32 (срез float32)
	t.Run("Field with pointer to []float32, filled field", func(t *testing.T) {
		ValueSliceInt := []float32{1.1, 2.2, 3.3}
		testJsonData := []byte(`{"valueslicefloat32": [1.1, 2.2, 3.3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceFloat32)
		assert.True(t, result.SetValueSliceFloat32)
	})

	// 16. Тест: JSON с заполненным полем ValueSliceFloat64 (срез float64)
	t.Run("Field with pointer to []float64, filled field", func(t *testing.T) {
		ValueSliceInt := []float64{1.1, 2.2, 3.3}
		testJsonData := []byte(`{"valueslicefloat64": [1.1, 2.2, 3.3]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInt, result.ValueSliceFloat64)
		assert.True(t, result.SetValueSliceFloat64)
	})

	// 17. Тест: JSON с заполненным полем ValueSliceInterface (срез интерфейсов)
	t.Run("Field with pointer to []interface{}, filled field", func(t *testing.T) {
		ValueSliceInterface := []interface{}{float64(1), "string", true}
		testJsonData := []byte(`{"valuesliceinterface": [1, "string", true]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceInterface, result.ValueSliceInterface)
		assert.True(t, result.SetValueSliceInterface)
	})

	// 18. Тест: JSON с заполненным полем ValueMapStringKey (map[string]interface{})
	t.Run("Field with pointer to map, filled field", func(t *testing.T) {
		ValueMapStringKey := map[string]interface{}{
			"theme":         "dark",
			"languValueInt": "en",
		}
		testJsonData := []byte(`{"valuemapstringkey": {"theme": "dark", "languValueInt": "en"}}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueMapStringKey, result.ValueMapStringKey)
		assert.True(t, result.SetValueMapStringKey)
	})

	// 18. Тест: JSON с заполненным полем ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "example_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		testJsonData := []byte(`{"valuestruct": {"valuestring": "example_value_string"}}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueStruct, result.ValueStruct)
		assert.True(t, result.SetValueStruct)
	})

	// 19. Тест: JSON с заполненным полем ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "example_value_string"
		ValuePtrStruct := TestChildPtrStructData{
			ValueString: &stringValue,
		}
		testJsonData := []byte(`{"valueptrstruct": {"valuestring": "example_value_string"}}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValuePtrStruct, result.ValuePtrStruct)
		assert.True(t, result.SetValuePtrStruct)
	})

	// 20. Тест: JSON без поля ValueStruct (TestChildPtrStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "exmaple_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		ValueSliceStruct := []TestChildStructData{ValueStruct}
		testJsonData := []byte(`{"valueslicestruct": [{"valuestring": "exmaple_value_string"}]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSliceStruct, result.ValueSliceStruct)
		assert.True(t, result.SetValueSliceStruct)
	})

	// 21. Тест: JSON с заполненным полем ValueStructSlice (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "example_value_string"
		ValueStruct := TestChildPtrStructData{
			ValueString: &stringValue,
		}
		ValueSlicePtrStruct := []TestChildPtrStructData{ValueStruct}
		testJsonData := []byte(`{"valuesliceptrstruct": [{"valuestring": "example_value_string"}]}`)
		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueSlicePtrStruct, result.ValueSlicePtrStruct)
		assert.True(t, result.SetValueSlicePtrStruct)
	})

	// 20. Тест: JSON с заполненными всеми полями
	t.Run("Field with all fields filled", func(t *testing.T) {
		ValueString := "John Doe"
		ValueInt := 30
		ValueInt8 := int8(25)
		ValueInt16 := int16(40)
		ValueInt32 := int32(35)
		ValueInt64 := int64(50)
		ValueFloat64 := 175.5
		active := true
		ValueSliceString := []string{"tag1", "tag2", "tag3"}
		ValueSliceInt8 := []int8{1, 2, 3}
		ValueSliceInt16 := []int16{1, 2, 3}
		ValueSliceInt32 := []int32{1, 2, 3}
		ValueSliceInt64 := []int64{1, 2, 3}
		ValueSliceInterface := []interface{}{float64(1), "string", true}
		ValueMapStringKey := map[string]interface{}{
			"theme":         "dark",
			"languValueInt": "en",
		}
		stringValue := "example_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}

		testJsonData := []byte(`{
			"valuestring": "` + ValueString + `",
			"valueint": ` + fmt.Sprintf("%d", ValueInt) + `,
			"valueint8": ` + fmt.Sprintf("%d", ValueInt8) + `,
			"valueint16": ` + fmt.Sprintf("%d", ValueInt16) + `,
			"valueint32": ` + fmt.Sprintf("%d", ValueInt32) + `,
			"valueint64": ` + fmt.Sprintf("%d", ValueInt64) + `,
			"valuefloat64": ` + fmt.Sprintf("%f", ValueFloat64) + `,
			"active": ` + strconv.FormatBool(active) + `,
			"valueslicestring": ["tag1", "tag2", "tag3"],
			"valuesliceint8": [1, 2, 3],
			"valuesliceint16": [1, 2, 3],
			"valuesliceint32": [1, 2, 3],
			"valuesliceint64": [1, 2, 3],
			"valuesliceinterface": [1, "string", true],
			"valuemapstringkey": {"theme": "dark", "languValueInt": "en"},
			"valuestruct": {"valuestring": "example_value_string"}
		}`)

		result := &TestPtrStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, &ValueString, result.ValueString)
		assert.Equal(t, &ValueInt, result.ValueInt)
		assert.Equal(t, &ValueInt8, result.ValueInt8)
		assert.Equal(t, &ValueInt16, result.ValueInt16)
		assert.Equal(t, &ValueInt32, result.ValueInt32)
		assert.Equal(t, &ValueInt64, result.ValueInt64)
		assert.Equal(t, &ValueFloat64, result.ValueFloat64)
		assert.Equal(t, &active, result.Active)
		assert.Equal(t, &ValueSliceString, result.ValueSliceString)
		assert.Equal(t, &ValueSliceInt8, result.ValueSliceInt8)
		assert.Equal(t, &ValueSliceInt16, result.ValueSliceInt16)
		assert.Equal(t, &ValueSliceInt32, result.ValueSliceInt32)
		assert.Equal(t, &ValueSliceInt64, result.ValueSliceInt64)
		assert.Equal(t, &ValueSliceInterface, result.ValueSliceInterface)
		assert.Equal(t, &ValueMapStringKey, result.ValueMapStringKey)
		assert.Equal(t, &ValueStruct, result.ValueStruct)

		// Проверка флагов
		assert.True(t, result.SetValueString)
		assert.True(t, result.SetValueInt)
		assert.True(t, result.SetValueInt8)
		assert.True(t, result.SetValueInt16)
		assert.True(t, result.SetValueInt32)
		assert.True(t, result.SetValueInt64)
		assert.True(t, result.SetValueFloat64)
		assert.True(t, result.SetActive)
		assert.True(t, result.SetValueSliceString)
		assert.True(t, result.SetValueSliceInt8)
		assert.True(t, result.SetValueSliceInt16)
		assert.True(t, result.SetValueSliceInt32)
		assert.True(t, result.SetValueSliceInt64)
		assert.True(t, result.SetValueSliceInterface)
		assert.True(t, result.SetValueMapStringKey)
		assert.True(t, result.SetValueStruct)
	})
}

func TestTestStructDataWithFilledFields(t *testing.T) {
	// 1. Тест: JSON с заполненным полем ValueString (string)
	t.Run("Field with string type, filled field", func(t *testing.T) {
		ValueString := "John Doe"
		testJsonData := []byte(`{"valuestring": "` + ValueString + `"}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueString, result.ValueString)
		assert.True(t, result.SetValueString)
	})

	// 2. Тест: JSON с заполненным полем ValueInt (int)
	t.Run("Field with pointer to int, filled field", func(t *testing.T) {
		ValueInt := 30
		testJsonData := []byte(`{"valueint": ` + fmt.Sprintf("%d", ValueInt) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueInt, result.ValueInt)
		assert.True(t, result.SetValueInt)
	})

	// 3. Тест: JSON с заполненным полем ValueInt8 (int8)
	t.Run("Field with pointer to int8, filled field", func(t *testing.T) {
		ValueInt8 := int8(25)
		testJsonData := []byte(`{"valueint8": ` + fmt.Sprintf("%d", ValueInt8) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueInt8, result.ValueInt8)
		assert.True(t, result.SetValueInt8)
	})

	// 4. Тест: JSON с заполненным полем ValueInt16 (int16)
	t.Run("Field with pointer to int16, filled field", func(t *testing.T) {
		ValueInt16 := int16(40)
		testJsonData := []byte(`{"valueint16": ` + fmt.Sprintf("%d", ValueInt16) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueInt16, result.ValueInt16)
		assert.True(t, result.SetValueInt16)
	})

	// 5. Тест: JSON с заполненным полем ValueInt32 (int32)
	t.Run("Field with pointer to int32, filled field", func(t *testing.T) {
		ValueInt32 := int32(35)
		testJsonData := []byte(`{"valueint32": ` + fmt.Sprintf("%d", ValueInt32) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueInt32, result.ValueInt32)
		assert.True(t, result.SetValueInt32)
	})

	// 6. Тест: JSON с заполненным полем ValueInt64 (int64)
	t.Run("Field with pointer to int64, filled field", func(t *testing.T) {
		ValueInt64 := int64(50)
		testJsonData := []byte(`{"valueint64": ` + fmt.Sprintf("%d", ValueInt64) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueInt64, result.ValueInt64)
		assert.True(t, result.SetValueInt64)
	})

	// 7. Тест: JSON с заполненным полем ValueFloat64 (float32)
	t.Run("Field with float64, filled field", func(t *testing.T) {
		ValueFloat32 := float32(175.5)
		testJsonData := []byte(`{"valuefloat32": ` + fmt.Sprintf("%f", ValueFloat32) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueFloat32, result.ValueFloat32)
		assert.True(t, result.SetValueFloat32)
	})

	// 8. Тест: JSON с заполненным полем ValueFloat64 (float64)
	t.Run("Field with float64, filled field", func(t *testing.T) {
		ValueFloat64 := 175.5
		testJsonData := []byte(`{"valuefloat64": ` + fmt.Sprintf("%f", ValueFloat64) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueFloat64, result.ValueFloat64)
		assert.True(t, result.SetValueFloat64)
	})

	// 9. Тест: JSON с заполненным полем Active (bool)
	t.Run("Field with pointer to bool, filled field", func(t *testing.T) {
		active := true
		testJsonData := []byte(`{"active": ` + strconv.FormatBool(active) + `}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, active, result.Active)
		assert.True(t, result.SetActive)
	})

	// 10. Тест: JSON с заполненным полем ValueSliceString (срез строк)
	t.Run("Field with pointer to []string, filled field", func(t *testing.T) {
		ValueSliceString := []string{"tag1", "tag2", "tag3"}
		testJsonData := []byte(`{"valueslicestring": ["tag1", "tag2", "tag3"]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceString, result.ValueSliceString)
		assert.True(t, result.SetValueSliceString)
	})

	// 11. Тест: JSON с заполненным полем ValueSliceInt (срез int)
	t.Run("Field with pointer to []int, filled field", func(t *testing.T) {
		ValueSliceInt := []int{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint": [1, 2, 3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceInt)
		assert.True(t, result.SetValueSliceInt)
	})

	// 12. Тест: JSON с заполненным полем ValueSliceInt8 (срез int8)
	t.Run("Field with pointer to []int8, filled field", func(t *testing.T) {
		ValueSliceInt := []int8{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint8": [1, 2, 3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceInt8)
		assert.True(t, result.SetValueSliceInt8)
	})

	// 13. Тест: JSON с заполненным полем ValueSliceInt16 (срез int16)
	t.Run("Field with pointer to []int16, filled field", func(t *testing.T) {
		ValueSliceInt := []int16{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint16": [1, 2, 3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceInt16)
		assert.True(t, result.SetValueSliceInt16)
	})

	// 14. Тест: JSON с заполненным полем ValueSliceInt32 (срез int32)
	t.Run("Field with pointer to []int32, filled field", func(t *testing.T) {
		ValueSliceInt := []int32{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint32": [1, 2, 3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceInt32)
		assert.True(t, result.SetValueSliceInt32)
	})

	// 15. Тест: JSON с заполненным полем ValueSliceInt64 (срез int64)
	t.Run("Field with pointer to []int64, filled field", func(t *testing.T) {
		ValueSliceInt := []int64{1, 2, 3}
		testJsonData := []byte(`{"valuesliceint64": [1, 2, 3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceInt64)
		assert.True(t, result.SetValueSliceInt64)
	})

	// 16. Тест: JSON с заполненным полем ValueSliceFloat64 (срез float64)
	t.Run("Field with pointer to []float64, filled field", func(t *testing.T) {
		ValueSliceInt := []float64{1.1, 2.2, 3.3}
		testJsonData := []byte(`{"valueslicefloat64": [1.1, 2.2, 3.3]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInt, result.ValueSliceFloat64)
		assert.True(t, result.SetValueSliceFloat64)
	})

	// 17. Тест: JSON с заполненным полем ValueSliceInterface (срез интерфейсов)
	t.Run("Field with pointer to []interface{}, filled field", func(t *testing.T) {
		ValueSliceInterface := []interface{}{float64(1), "string", true}
		testJsonData := []byte(`{"valuesliceinterface": [1, "string", true]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceInterface, result.ValueSliceInterface)
		assert.True(t, result.SetValueSliceInterface)
	})

	// 18. Тест: JSON с заполненным полем ValueMapStringKey (map[string]interface{})
	t.Run("Field with pointer to map, filled field", func(t *testing.T) {
		ValueMapStringKey := map[string]interface{}{
			"theme":         "dark",
			"languValueInt": "en",
		}
		testJsonData := []byte(`{"valuemapstringkey": {"theme": "dark", "languValueInt": "en"}}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueMapStringKey, result.ValueMapStringKey)
		assert.True(t, result.SetValueMapStringKey)
	})

	// 19. Тест: JSON без поля ValueStruct (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "exmaple_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		testJsonData := []byte(`{"valuestruct": {"valuestring": "exmaple_value_string"}}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueStruct, result.ValueStruct)
		assert.True(t, result.SetValueStruct)
	})
	// 20. Тест: JSON без поля ValueStruct (TestChildPtrStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "exmaple_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		testJsonData := []byte(`{"valuestruct": {"valuestring": "exmaple_value_string"}}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueStruct, result.ValueStruct)
		assert.True(t, result.SetValueStruct)
	})

	// 21. Тест: JSON с заполненным полем ValueStructSlice (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "example_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		ValueSliceStruct := []TestChildStructData{ValueStruct}
		testJsonData := []byte(`{"valueslicestruct": [{"valuestring": "example_value_string"}]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSliceStruct, result.ValueSliceStruct)
		assert.True(t, result.SetValueSliceStruct)
	})

	// 22. Тест: JSON с заполненным полем ValueStructSlice (TestChildStructData)
	t.Run("Field with pointer to map, missing field", func(t *testing.T) {
		stringValue := "example_value_string"
		ValueStruct := TestChildPtrStructData{
			ValueString: &stringValue,
		}
		ValueSlicePtrStruct := []TestChildPtrStructData{ValueStruct}
		testJsonData := []byte(`{"valuesliceptrstruct": [{"valuestring": "example_value_string"}]}`)
		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueSlicePtrStruct, result.ValueSlicePtrStruct)
		assert.True(t, result.SetValueSlicePtrStruct)
	})

	// 23. Тест: JSON с заполненными всеми полями
	t.Run("Field with all fields filled", func(t *testing.T) {
		ValueString := "John Doe"
		ValueInt := 30
		ValueInt8 := int8(25)
		ValueInt16 := int16(40)
		ValueInt32 := int32(35)
		ValueInt64 := int64(50)
		ValueFloat64 := 175.5
		active := true
		ValueSliceString := []string{"tag1", "tag2", "tag3"}
		ValueSliceInt8 := []int8{1, 2, 3}
		ValueSliceInt16 := []int16{1, 2, 3}
		ValueSliceInt32 := []int32{1, 2, 3}
		ValueSliceInt64 := []int64{1, 2, 3}
		ValueSliceInterface := []interface{}{float64(1), "string", true}
		ValueMapStringKey := map[string]interface{}{
			"theme":         "dark",
			"languValueInt": "en",
		}
		stringValue := "example_value_string"
		ValueStruct := TestChildStructData{
			ValueString: stringValue,
		}
		ValuePtrStruct := TestChildPtrStructData{
			ValueString: &stringValue,
		}
		ValueSliceStruct := []TestChildStructData{ValueStruct}
		ValueSlicePtrStruct := []TestChildPtrStructData{ValuePtrStruct}

		testJsonData := []byte(fmt.Sprintf(`{
			"valuestring": "%s", 
			"valueint": %d,
			"valueint8": %d,
			"valueint16": %d,
			"valueint32": %d,
			"valueint64": %d,
			"valuefloat64": %f,
			"active": %t,
			"valueslicestring": ["tag1", "tag2", "tag3"],
			"valuesliceint8": [1, 2, 3],
			"valuesliceint8": [1, 2, 3],
			"valuesliceint16": [1, 2, 3],
			"valuesliceint32": [1, 2, 3],
			"valuesliceint64": [1, 2, 3],
			"valuesliceinterface": [1, "string", true],
			"valuemapstringkey": {"theme": "dark", "languValueInt": "en"},
			"valuestruct": {"valuestring": "example_value_string"},
			"valueslicestruct": [{"valuestring": "example_value_string"}],
			"valuesliceptrstruct": [{"valuestring": "example_value_string"}]
		}`, ValueString, ValueInt, ValueInt8, ValueInt16, ValueInt32, ValueInt64, ValueFloat64, active))

		result := &TestStructData{}
		err := PopulateFieldsAndSets(testJsonData, result)
		assert.NoError(t, err)
		assert.Equal(t, ValueString, result.ValueString)
		assert.Equal(t, ValueInt, result.ValueInt)
		assert.Equal(t, ValueInt8, result.ValueInt8)
		assert.Equal(t, ValueInt16, result.ValueInt16)
		assert.Equal(t, ValueInt32, result.ValueInt32)
		assert.Equal(t, ValueInt64, result.ValueInt64)
		assert.Equal(t, ValueFloat64, result.ValueFloat64)
		assert.Equal(t, active, result.Active)
		assert.Equal(t, ValueSliceString, result.ValueSliceString)
		assert.Equal(t, ValueSliceInt8, result.ValueSliceInt8)
		assert.Equal(t, ValueSliceInt16, result.ValueSliceInt16)
		assert.Equal(t, ValueSliceInt32, result.ValueSliceInt32)
		assert.Equal(t, ValueSliceInt64, result.ValueSliceInt64)
		assert.Equal(t, ValueSliceInterface, result.ValueSliceInterface)
		assert.Equal(t, ValueMapStringKey, result.ValueMapStringKey)
		assert.Equal(t, ValueStruct, result.ValueStruct)
		assert.Equal(t, ValueSliceStruct, result.ValueSliceStruct)
		assert.Equal(t, ValueSlicePtrStruct, result.ValueSlicePtrStruct)
		assert.True(t, result.SetValueString)
		assert.True(t, result.SetValueInt)
		assert.True(t, result.SetValueInt8)
		assert.True(t, result.SetValueInt16)
		assert.True(t, result.SetValueInt32)
		assert.True(t, result.SetValueInt64)
		assert.True(t, result.SetValueFloat64)
		assert.True(t, result.SetActive)
		assert.True(t, result.SetValueSliceString)
		assert.True(t, result.SetValueSliceInt8)
		assert.True(t, result.SetValueSliceInt16)
		assert.True(t, result.SetValueSliceInt32)
		assert.True(t, result.SetValueSliceInt64)
		assert.True(t, result.SetValueSliceInterface)
		assert.True(t, result.SetValueMapStringKey)
		assert.True(t, result.SetValueStruct)
		assert.True(t, result.SetValueSliceStruct)
		assert.True(t, result.SetValueSlicePtrStruct)
	})
}
