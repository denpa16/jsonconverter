package main

import (
	"fmt"
)

type ExampleChildStruct struct {
	ID int `json:"id"`
}

// Структура, которая описывает результат
type ResultData struct {
	Name                  *string                 `json:"name"`
	Age                   int                     `json:"age"`
	Age8                  *int8                   `json:"age8"`
	Age16                 *int16                  `json:"age16"`
	Age32                 *int32                  `json:"age32"`
	Age64                 *int64                  `json:"age64"`
	Height                float32                 `json:"height"`
	Active                *bool                   `json:"active"`
	Tags                  *[]string               `json:"tags"`     // указатель на срез строк
	Numbers               *[]int32                `json:"numbers"`  // указатель на срез целых чисел
	Objects               *[]interface{}          `json:"objects"`  // указатель на срез интерфейсов
	Settings              *map[string]interface{} `json:"settings"` // указатель на карту
	ExampleStruct         ExampleChildStruct      `json:"examplestruct"`
	ExampleStructSlice    []ExampleChildStruct    `json:"examplestructslice"`
	ExampleStructPtrSlice *[]ExampleChildStruct   `json:"examplestructptrslice"`

	// Дополнительные поля для хранения Set
	SetName    bool `json:"-"`
	SetAge     bool `json:"-"`
	SetAge8    bool `json:"-"`
	SetAge16   bool `json:"-"`
	SetAge32   bool `json:"-"`
	SetAge64   bool `json:"-"`
	SetHeight  bool `json:"-"`
	SetActive  bool `json:"-"`
	SetTags    bool `json:"-"`
	SetNumbers bool `json:"-"`
	SetObjects bool `json:"-"`
	//SetSettings bool `json:"-"`
}

func main() {
	// Пример JSON с различными типами данных
	data := `{
		"age": 30,
		"age8": 10,
		"age16": 100,
		"age32": 1000,
		"age64": 10000,
		"height": 175.5,
		"active": true,
		"tags": ["golang", "magic", "reflection"],
		"numbers": [1,2,3],
		"objects": [false, 1, {"title": "Vasco", "chars": [1,2,3], "nas": {"r": "222"}}],
		"settings": {"1": "2", "2": 2},
		"examplestruct": {"id": 1},
		"examplestructslice": [{"id": 1}],
		"examplestructptrslice": [{"id": 1}]
	}`

	// Инициализируем структуру ResultData
	result := &ResultData{}

	// Заполняем структуру данными из JSON

	testData := []byte(data)

	err := PopulateFieldsAndSets(testData, result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Выводим результат
	//fmt.Println(result.ExampleStructSlice)
	fmt.Println(result.ExampleStructPtrSlice)
}
