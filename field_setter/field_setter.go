package field_setter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// setFields рекурсивно устанавливает Set-поля в true, если соответствующие поля присутствуют в JSON

func FillFields(result interface{}, data []byte) error {

	if err := json.Unmarshal(data, result); err != nil {
		fmt.Println("Error unmarshalling JSON into struct:", err)
		return fmt.Errorf("error: %v", err)
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return fmt.Errorf("error: %v", err)
	}
	SetFields(result, jsonMap)
	return nil
}

func SetFields(data interface{}, jsonData map[string]interface{}) error {
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// Пропускаем поля, которые не являются Set-полями
		if strings.HasPrefix(fieldName, "Set") {
			continue
		}
		jsonTag := typ.Field(i).Tag.Get("json")

		if nestedField, exists := jsonData[jsonTag]; exists {
			setField := val.FieldByName("Set" + fieldName)
			if setField.IsValid() {
				setField.SetBool(true)
			}
			switch field.Kind() {
			case reflect.Struct:
				// Если поле — структура, рекурсивно вызываем SetFields
				if nestedMap, ok := nestedField.(map[string]interface{}); ok {
					SetFields(field.Addr().Interface(), nestedMap)
				}
			case reflect.Ptr:
				// Если поле — указатель, обрабатываем его в зависимости от типа
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				switch field.Elem().Kind() {
				case reflect.Struct:
					// Если указатель на структуру, рекурсивно вызываем SetFields
					if nestedMap, ok := nestedField.(map[string]interface{}); ok {
						SetFields(field.Interface(), nestedMap)
					}
				case reflect.Slice:
					// Если указатель на срез, обрабатываем каждый элемент
					if nestedSlice, ok := nestedField.([]interface{}); ok {
						slice := reflect.MakeSlice(field.Elem().Type(), len(nestedSlice), len(nestedSlice))
						field.Elem().Set(slice)
						for j := 0; j < slice.Len(); j++ {
							elem := slice.Index(j)
							if elem.Kind() == reflect.Ptr {
								elem.Set(reflect.New(elem.Type().Elem()))
							}
							if nestedMap, ok := nestedSlice[j].(map[string]interface{}); ok {
								// TODO: надо добавить другие типы на стоп-рекурсию
								if field.Type().Elem().Elem().Kind().String() == "uint8" {
									continue
								} else {
									SetFields(elem.Addr().Interface(), nestedMap)
								}

							} else {
								// Если элемент не является картой, просто устанавливаем значение
								elem.Set(reflect.ValueOf(nestedSlice[j]).Convert(elem.Type()))
							}
						}
					}
				case reflect.Map:
					// Если указатель на карту, создаём новую карту и заполняем её
					if nestedMap, ok := nestedField.(map[string]interface{}); ok {
						field.Elem().Set(reflect.MakeMap(field.Elem().Type()))
						for key, value := range nestedMap {
							mapValue := reflect.New(field.Elem().Type().Elem()).Elem()
							if mapValue.Kind() == reflect.Ptr {
								mapValue.Set(reflect.New(mapValue.Type().Elem()))
							}
							if nestedValueMap, ok := value.(map[string]interface{}); ok {
								SetFields(mapValue.Addr().Interface(), nestedValueMap)
							} else {
								// Если значение не является картой, просто устанавливаем значение
								mapValue.Set(reflect.ValueOf(value).Convert(mapValue.Type()))
							}
							field.Elem().SetMapIndex(reflect.ValueOf(key), mapValue)
						}
					}
				default:
					// Если указатель на примитивный тип, просто устанавливаем значение
					if nestedField != nil {
						field.Elem().Set(reflect.ValueOf(nestedField).Convert(field.Elem().Type()))
					}
				}
			case reflect.Slice:
				// Если поле — срез, обрабатываем каждый элемент
				if nestedSlice, ok := nestedField.([]interface{}); ok {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						if elem.Kind() == reflect.Ptr {
							if elem.IsNil() {
								elem.Set(reflect.New(elem.Type().Elem()))
							}
							if nestedMap, ok := nestedSlice[j].(map[string]interface{}); ok {
								SetFields(elem.Interface(), nestedMap)
							} else {
								// Если элемент не является картой, просто устанавливаем значение
								elem.Set(reflect.ValueOf(nestedSlice[j]).Convert(elem.Type()))
							}
						} else if elem.Kind() == reflect.Struct {
							if nestedMap, ok := nestedSlice[j].(map[string]interface{}); ok {
								SetFields(elem.Addr().Interface(), nestedMap)
							} else {
								// Если элемент не является картой, просто устанавливаем значение
								elem.Set(reflect.ValueOf(nestedSlice[j]).Convert(elem.Type()))
							}
						}
					}
				}
			case reflect.Map:
				// Если поле — карта, создаём новую карту и заполняем её
				if nestedMap, ok := nestedField.(map[string]interface{}); ok {
					field.Set(reflect.MakeMap(field.Type()))
					for key, value := range nestedMap {
						mapValue := reflect.New(field.Type().Elem()).Elem()
						if mapValue.Kind() == reflect.Ptr {
							mapValue.Set(reflect.New(mapValue.Type().Elem()))
						}
						if nestedValueMap, ok := value.(map[string]interface{}); ok {
							// Если значение — карта, рекурсивно вызываем SetFields
							SetFields(mapValue.Addr().Interface(), nestedValueMap)
						} else {
							// Если значение не является картой, просто устанавливаем значение
							mapValue.Set(reflect.ValueOf(value).Convert(mapValue.Type()))
						}
						field.SetMapIndex(reflect.ValueOf(key), mapValue)
					}
				}
			default:
				// Если поле — примитивный тип, просто устанавливаем значение
				if nestedField != nil {
					field.Set(reflect.ValueOf(nestedField).Convert(field.Type()))
				}
			}
		}
	}
	return nil
}
