package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Функция для проверки наличия полей из struct в JSON и присвоения значений
func PopulateFieldsAndSets(jsonData []byte, result interface{}) error {
	// Разбираем JSON в map[string]interface{}
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	// Получаем отражение структуры
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	// Перебираем все поля структуры
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		// Пропускаем поля, которые не имеют тег json
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		// Проверяем наличие поля в JSON
		if jsonValue, exists := data[jsonTag]; exists {
			// Присваиваем значение полю
			switch field.Kind() {
			case reflect.String:
				field.SetString(jsonValue.(string))
			case reflect.Bool:
				field.SetBool(jsonValue.(bool))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				// Обрабатываем типы int с приведение типов
				// jsonValue ok всегда float64, так как это делает
				if val, ok := jsonValue.(float64); ok {
					// Преобразуем число в нужный тип
					switch field.Type().Kind() {
					case reflect.Int:
						intVal := int(val) // Преобразуем в int
						field.Set(reflect.ValueOf(intVal))
					case reflect.Int8:
						int8Val := int8(val)
						field.Set(reflect.ValueOf(int8Val))
					case reflect.Int16:
						int16Val := int16(val)
						field.Set(reflect.ValueOf(int16Val))
					case reflect.Int32:
						int32Val := int32(val)
						field.Set(reflect.ValueOf(int32Val))
					case reflect.Int64:
						int64Val := int64(val)
						field.Set(reflect.ValueOf(int64Val))
					}
				}
			case reflect.Float32:
				if val, ok := jsonValue.(float64); ok {
					float32Val := float32(val)
					field.Set(reflect.ValueOf(float32Val))
				}
			case reflect.Float64:
				if val, ok := jsonValue.(float64); ok {
					field.Set(reflect.ValueOf(val))
				}
			case reflect.Ptr:
				// Обрабатываем указатели на другие типы
				switch field.Type().Elem().Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					// Преобразуем число в нужный тип
					if val, ok := jsonValue.(float64); ok {
						// Создаем новый экземпляр для указателя на int, int8, int16, int32, int64
						switch field.Type().Elem().Kind() {
						case reflect.Int:
							intVal := int(val)
							field.Set(reflect.ValueOf(&intVal))
						case reflect.Int8:
							int8Val := int8(val)
							field.Set(reflect.ValueOf(&int8Val))
						case reflect.Int16:
							int16Val := int16(val)
							field.Set(reflect.ValueOf(&int16Val))
						case reflect.Int32:
							int32Val := int32(val)
							field.Set(reflect.ValueOf(&int32Val))
						case reflect.Int64:
							int64Val := int64(val)
							field.Set(reflect.ValueOf(&int64Val))
						}
					}
				case reflect.Float32:
					if val, ok := jsonValue.(float64); ok {
						float32Val := float32(val)
						field.Set(reflect.ValueOf(&float32Val))
					}
				case reflect.Float64:
					if val, ok := jsonValue.(float64); ok {
						float64Val := float64(val)
						field.Set(reflect.ValueOf(&float64Val))
					}
				case reflect.String:
					// Преобразуем строку
					if val, ok := jsonValue.(string); ok {
						field.Set(reflect.ValueOf(&val))
					}
				case reflect.Bool:
					// Преобразуем булево значение
					if val, ok := jsonValue.(bool); ok {
						field.Set(reflect.ValueOf(&val))
					}
				case reflect.Map:
					if val, ok := jsonValue.(map[string]interface{}); ok {
						field.Set(reflect.ValueOf(&val))
					}
				case reflect.Struct:
					if val, ok := jsonValue.(map[string]interface{}); ok {
						byteValue, _ := json.Marshal(val)
						newStruct := reflect.New(field.Type().Elem()).Interface()
						err := PopulateFieldsAndSets(byteValue, newStruct)
						if err != nil {
							return err
						}
						field.Set(reflect.ValueOf(newStruct))
					}
				case reflect.Slice:
					var interfaceValue interface{} = data[jsonTag]
					switch t.Field(i).Type.Elem().Elem().Kind() {
					case reflect.Int:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []int
						var intSlice []int
						// Преобразуем каждый элемент из []interface{} в []int
						for _, v := range values {
							// Преобразуем каждый элемент к типу int8 и добавляем в срез
							if val, ok := v.(float64); ok {
								intSlice = append(intSlice, int(val))
							} else {
								fmt.Println("Impossible convert slice value in int")
							}
						}
						field.Set(reflect.ValueOf(&intSlice))
					case reflect.Int8:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []int8
						var intSlice []int8
						// Преобразуем каждый элемент из []interface{} в []int8
						for _, v := range values {
							// Преобразуем каждый элемент к типу int8 и добавляем в срез
							if val, ok := v.(float64); ok {
								intSlice = append(intSlice, int8(val))
							} else {
								fmt.Println("Impossible convert slice value in int8")
							}
						}
						field.Set(reflect.ValueOf(&intSlice))
					case reflect.Int16:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []int16
						var intSlice []int16
						// Преобразуем каждый элемент из []interface{} в []int16
						for _, v := range values {
							// Преобразуем каждый элемент к типу int16 и добавляем в срез
							if val, ok := v.(float64); ok {
								intSlice = append(intSlice, int16(val))
							} else {
								fmt.Println("Impossible convert slice value in int16")
							}
						}
						field.Set(reflect.ValueOf(&intSlice))
					case reflect.Int32:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []int32
						var intSlice []int32
						// Преобразуем каждый элемент из []interface{} в []int32
						for _, v := range values {
							// Преобразуем каждый элемент к типу int32 и добавляем в срез
							if val, ok := v.(float64); ok {
								intSlice = append(intSlice, int32(val))
							} else {
								fmt.Println("Impossible convert slice value in int32")
							}
						}
						field.Set(reflect.ValueOf(&intSlice))
					case reflect.Int64:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []int64
						var intSlice []int64
						// Преобразуем каждый элемент из []interface{} в []int64
						for _, v := range values {
							// Преобразуем каждый элемент к типу int64 и добавляем в срез
							if val, ok := v.(float64); ok {
								intSlice = append(intSlice, int64(val))
							} else {
								fmt.Println("Impossible convert slice value in int64")
							}
						}
						field.Set(reflect.ValueOf(&intSlice))
					case reflect.Float64:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []float64
						var float64Slice []float64
						// Преобразуем каждый элемент из []interface{} в []float64
						for _, v := range values {
							// Преобразуем каждый элемент к типу float64 и добавляем в срез
							if val, ok := v.(float64); ok {
								float64Slice = append(float64Slice, float64(val))
							} else {
								fmt.Println("Impossible convert slice value in float64")
							}
						}
						field.Set(reflect.ValueOf(&float64Slice))
					case reflect.Float32:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []float32
						var float32Slice []float32
						// Преобразуем каждый элемент из []interface{} в []float32
						for _, v := range values {
							// Преобразуем каждый элемент к типу float64 и добавляем в срез
							if val, ok := v.(float64); ok {
								float32Slice = append(float32Slice, float32(val))
							} else {
								fmt.Println("Impossible convert slice value in float32")
							}
						}
						field.Set(reflect.ValueOf(&float32Slice))
					case reflect.String:
						// Преобразование интерфейса в срез
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Imposible convert interface to slice")
						}
						// Создаем срез типа []string
						var stringSlice []string
						// Преобразуем каждый элемент из []interface{} в []string
						for _, v := range values {
							// Преобразуем каждый элемент к типу int и добавляем в срез
							if val, ok := v.(string); ok {
								stringSlice = append(stringSlice, string(val))
							} else {
								fmt.Println("Impossible convert slice value in string")
							}
						}
						field.Set(reflect.ValueOf(&stringSlice))
					case reflect.Struct:
						// Обрабатываем структуру
						values, ok := interfaceValue.([]interface{})
						if !ok {
							fmt.Println("Impossible to convert interface to slice")
						}

						// Создаем срез нужного типа
						structSlice := reflect.MakeSlice(field.Type().Elem(), len(values), len(values))

						for k, v := range values {
							// Преобразуем каждый элемент к типу нужной структуры
							newStruct := reflect.New(field.Type().Elem().Elem()).Elem()
							byteValue, _ := json.Marshal(v)
							// Пример заполнения структуры (если нужно)
							err := PopulateFieldsAndSets(byteValue, newStruct.Addr().Interface())
							if err != nil {
								return err
							}
							structSlice.Index(k).Set(newStruct)
						}

						// Получаем указатель на срез, передаем его в field.Set
						structSlicePtr := reflect.New(structSlice.Type())
						structSlicePtr.Elem().Set(structSlice) // Копируем срез в указатель на срез

						// Присваиваем указатель на срез в поле
						field.Set(structSlicePtr)
					default:
						// Преобразование интерфейса в срез
						interfaceData, ok := interfaceValue.([]interface{})
						if ok {
							if field.Type().String() == "*json.RawMessage" {
								rawJsonConverted, _ := json.Marshal(interfaceData)
								rawMessage := json.RawMessage(rawJsonConverted)
								field.Set(reflect.ValueOf(&rawMessage))
							} else {
								field.Set(reflect.ValueOf(&interfaceData))
							}
						} else {
							if field.Type().String() == "*json.RawMessage" {
								rawJsonConverted, _ := json.Marshal(interfaceData)
								rawMessage := json.RawMessage(rawJsonConverted)
								field.Set(reflect.ValueOf(&rawMessage))
							} else {
								field.Set(reflect.ValueOf(&interfaceData))
							}
						}
					}
				}
			case reflect.Map:
				// Обрабатываем карты
				if val, ok := jsonValue.(map[string]interface{}); ok {
					field.Set(reflect.ValueOf(val))
				}
			case reflect.Struct:
				// Обрабатываем cтруктуру
				if val, ok := jsonValue.(map[string]interface{}); ok {
					byteValue, _ := json.Marshal(val)
					err := PopulateFieldsAndSets(byteValue, field.Addr().Interface())
					if err != nil {
						return err
					}
				}
			case reflect.Slice:
				var interfaceValue interface{} = data[jsonTag]
				switch t.Field(i).Type.Elem().Kind() {
				case reflect.Int:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []int
					var intSlice []int
					// Преобразуем каждый элемент из []interface{} в []int
					for _, v := range values {
						// Преобразуем каждый элемент к типу int8 и добавляем в срез
						if val, ok := v.(float64); ok {
							intSlice = append(intSlice, int(val))
						} else {
							fmt.Println("Impossible convert slice value in int")
						}
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.Int8:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []int8
					var intSlice []int8
					// Преобразуем каждый элемент из []interface{} в []int8
					for _, v := range values {
						// Преобразуем каждый элемент к типу int8 и добавляем в срез
						if val, ok := v.(float64); ok {
							intSlice = append(intSlice, int8(val))
						} else {
							fmt.Println("Impossible convert slice value in int8")
						}
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.Int16:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []int16
					var intSlice []int16
					// Преобразуем каждый элемент из []interface{} в []int16
					for _, v := range values {
						// Преобразуем каждый элемент к типу int16 и добавляем в срез
						if val, ok := v.(float64); ok {
							intSlice = append(intSlice, int16(val))
						} else {
							fmt.Println("Impossible convert slice value in int16")
						}
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.Int32:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []int32
					var intSlice []int32
					// Преобразуем каждый элемент из []interface{} в []int32
					for _, v := range values {
						// Преобразуем каждый элемент к типу int32 и добавляем в срез
						if val, ok := v.(float64); ok {
							intSlice = append(intSlice, int32(val))
						} else {
							fmt.Println("Impossible convert slice value in int32")
						}
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.Int64:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []int64
					var intSlice []int64
					// Преобразуем каждый элемент из []interface{} в []int64
					for _, v := range values {
						// Преобразуем каждый элемент к типу int64 и добавляем в срез
						if val, ok := v.(float64); ok {
							intSlice = append(intSlice, int64(val))
						} else {
							fmt.Println("Impossible convert slice value in int64")
						}
					}
					field.Set(reflect.ValueOf(intSlice))
				case reflect.Float64:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []float64
					var float64Slice []float64
					// Преобразуем каждый элемент из []interface{} в []float64
					for _, v := range values {
						// Преобразуем каждый элемент к типу float64 и добавляем в срез
						if val, ok := v.(float64); ok {
							float64Slice = append(float64Slice, float64(val))
						} else {
							fmt.Println("Impossible convert slice value in float64")
						}
					}
					field.Set(reflect.ValueOf(float64Slice))
				case reflect.Float32:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []float32
					var float32Slice []float32
					// Преобразуем каждый элемент из []interface{} в []float32
					for _, v := range values {
						// Преобразуем каждый элемент к типу float64 и добавляем в срез
						if val, ok := v.(float64); ok {
							float32Slice = append(float32Slice, float32(val))
						} else {
							fmt.Println("Impossible convert slice value in float32")
						}
					}
					field.Set(reflect.ValueOf(float32Slice))
				case reflect.String:
					// Преобразование интерфейса в срез
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []string
					var stringSlice []string
					// Преобразуем каждый элемент из []interface{} в []string
					for _, v := range values {
						// Преобразуем каждый элемент к типу int и добавляем в срез
						if val, ok := v.(string); ok {
							stringSlice = append(stringSlice, string(val))
						} else {
							fmt.Println("Impossible convert slice value in string")
						}
					}
					field.Set(reflect.ValueOf(stringSlice))
				case reflect.Struct:
					// Обрабатываем cтруктуру
					values, ok := interfaceValue.([]interface{})
					if !ok {
						fmt.Println("Imposible convert interface to slice")
					}
					// Создаем срез типа []string
					//var stringSlice []interface{}
					// Преобразуем каждый элемент из []interface{} в []string
					structSlice := reflect.MakeSlice(field.Type(), len(values), len(values))
					for k, v := range values {
						// Преобразуем каждый элемент к типу int и добавляем в срез
						byteValue, _ := json.Marshal(v)
						newStruct := reflect.New(field.Type().Elem()).Elem()
						err := PopulateFieldsAndSets(byteValue, newStruct.Addr().Interface())
						if err != nil {
							return err
						}
						structSlice.Index(k).Set(newStruct)
						field.Set(structSlice)
					}
				default:
					// Преобразование интерфейса в срез
					interfaceData, ok := interfaceValue.([]interface{})
					if ok {
						if field.Type().String() == "json.RawMessage" {
							rawJsonConverted, _ := json.Marshal(interfaceData)
							rawMessage := json.RawMessage(rawJsonConverted)
							field.Set(reflect.ValueOf(rawMessage))
						} else {
							field.Set(reflect.ValueOf(interfaceData))
						}
					}
				}
			}

			// Устанавливаем значение Set<название поля> в true
			setField := v.FieldByName("Set" + fieldName)
			if setField.IsValid() {
				setField.SetBool(true)
			}
		} else {
			// Если поля нет в JSON, устанавливаем Set<название поля> в false
			setField := v.FieldByName("Set" + fieldName)
			if setField.IsValid() {
				setField.SetBool(false)
			}
		}
	}

	return nil
}
