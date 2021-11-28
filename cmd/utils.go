package cmd

import (
	"fmt"
	"reflect"
)

func PrintStructHeader(structValue interface{}) {
	fields := reflect.TypeOf(structValue)

	for j := 0; j < fields.NumField(); j++ {
		f := fields.Field(j)
		fmt.Print(f.Name, ",")
	}
	fmt.Println()

}

func PrintStructValues(structValue interface{}) {
	fields := reflect.TypeOf(structValue)
	values := reflect.ValueOf(structValue)

	for j := 0; j < fields.NumField(); j++ {
		f := fields.Field(j)
		value := values.FieldByName(f.Name).Interface()
		fmt.Print(value, ",")
	}
	fmt.Println()

}
