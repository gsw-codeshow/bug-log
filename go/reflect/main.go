package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Demo struct {
	Id   int32
	Name int32
}

func getValue(v interface{}, data map[string]int32) {
	t := reflect.TypeOf(v)
	o := reflect.ValueOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() //获取类型指针中的类型
	}
	if o.Kind() == reflect.Ptr {
		o = o.Elem() //获取值地址中的值
	}
	num := t.NumField() //获取字段个数
	for i := 0; i < num; i++ {
		f := o.Field(i)              //获取字段的值
		fieldName := t.Field(i).Name //获取字段名称
		switch f.Kind() {
		case reflect.Int:
		case reflect.String:
			fmt.Println(fieldName, ": ", f.String())
		case reflect.Int32:
			fmt.Println("---int32---")
			fmt.Println(fieldName, ": ", f.Int())
			f.SetInt(int64(data[fieldName]))
		default:
			fmt.Println("类型不不支持")
		}
	}
}
func main() {
	demo := &Demo{
		Id:   10,
		Name: 11,
	}
	demoByte, _ := json.Marshal(demo)
	var data map[string]int32
	data = make(map[string]int32)
	data["Id"] = 100
	demo2 := &Demo{}
	getValue(demo2, data)
	fmt.Println(demo2)
	_ = json.Unmarshal(demoByte, &data)
	fmt.Println(data)

	key := "WBossDamage_100"
	id, _ := strconv.Atoi(strings.Trim(key, "WBossDamage_"))
	fmt.Println(id)
	return
}
