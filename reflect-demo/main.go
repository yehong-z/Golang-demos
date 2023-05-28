package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10

	// 获取变量的类型信息
	fmt.Println(reflect.TypeOf(x))

	// 获取变量的值信息
	fmt.Println(reflect.ValueOf(x)) // 输出: 10

	// 修改变量的值信息
	v := reflect.ValueOf(&x).Elem()
	v.SetInt(20)
	fmt.Println(x) // 输出: 20

	// 判断类型是否相同
	var y int64 = 20
	fmt.Println(reflect.TypeOf(x).AssignableTo(reflect.TypeOf(y)))  // 输出: false
	fmt.Println(reflect.TypeOf(x).ConvertibleTo(reflect.TypeOf(y))) // 输出: true

	// 判断变量是否为 Nil
	var z *int = nil
	fmt.Println(reflect.ValueOf(z).IsNil()) // 输出: true
}
