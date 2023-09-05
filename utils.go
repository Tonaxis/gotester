package gotester

import (
	"fmt"
	"reflect"
	"strings"
)

func test(data TestRange, results Result) bool {
	if len(data.ExpectedReturns) != len(results) {
		return false
	}

	for i, expected := range data.ExpectedReturns {
		if !reflect.DeepEqual(expected, results[i]) {
			return false
		}
	}

	return true
}

func interfaceToString(args ...interface{}) string {
	var arguments []string
	for _, arg := range args {
		arguments = append(arguments, fmt.Sprintf("%v", arg))
	}
	return strings.Join(arguments, ", ")
}

func centerText(text string, spaceChar string, totalLength int) string {
	space := ""
	for i := 0; i < ((totalLength-2)-len(text))/2; i++ {
		space += spaceChar
	}

	return fmt.Sprintf("%s %s %s", space, text, space)
}
