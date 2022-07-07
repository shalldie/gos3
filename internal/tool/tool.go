package tool

import (
	"reflect"

	"github.com/shalldie/gog/hashmap"
)

func Struct2TypeTuples(sender any) *hashmap.HashMap[string, string] {
	hm := hashmap.New[string, string]()

	typeOpt := reflect.TypeOf(sender)

	for i := 0; i < typeOpt.NumField(); i++ {
		field := typeOpt.Field(i)

		hm.Set(field.Name, field.Type.Name())

	}

	return hm
}
