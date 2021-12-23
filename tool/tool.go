package tool

import "reflect"

func Struct2TypeTuple(sender interface{}) [][2]string {
	var typeTuple [][2]string

	typeOpt := reflect.TypeOf(sender)

	for i := 0; i < typeOpt.NumField(); i++ {
		field := typeOpt.Field(i)

		typeTuple = append(typeTuple, [2]string{
			field.Name, field.Type.Name(),
		})

	}

	return typeTuple
}
