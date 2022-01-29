package main

import "strings"

func Repeat(str string, count int) string {
	// output := ""
	// for i := 0; i < count; i++ {
	// 	output += str
	// }
	// return output
	return strings.Repeat(str, count)
}
