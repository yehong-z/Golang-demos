package convert

import (
	"fmt"
	"reflect"
)

func PrintNameOfField(any interface{}) {
	stype := reflect.TypeOf(any).Elem()
	for i := 0; i < stype.NumField(); i++ {
		field := stype.Field(i)
		fmt.Println(field.Name)
	}
}

// 通过反射来复制结构体字段，注意只能复制公有方法
func mapStruct(source interface{}, target interface{}) error {
	sourceValue := reflect.ValueOf(source).Elem()
	targetType := reflect.TypeOf(target).Elem()

	targetValue := reflect.New(targetType).Elem()

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		sourceField := sourceValue.FieldByName(field.Name)

		if sourceField.IsValid() && sourceField.Type().AssignableTo(field.Type) {
			targetValue.FieldByName(field.Name).Set(sourceField)
		}
	}

	reflect.ValueOf(target).Elem().Set(targetValue)

	return nil
}
