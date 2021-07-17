package utils

import (
	"reflect"
)

// DeepEQLString cheks m1 and m2 type string map are deeply qual or not
func DeepEQLString(m1 map[int]string, m2 map[int]string) bool {

	res := reflect.DeepEqual(m1, m2)

	return res

}

// DeepEQLInt cheks m1 and m2 type int map are deeply qual or not
func DeepEQLInt(m1 map[int]int, m2 map[int]int) bool {
	res := reflect.DeepEqual(m1, m2)

	return res

}
