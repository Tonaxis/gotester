package gotester

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

var typesByName = map[string]reflect.Type{
	"int":        reflect.TypeOf(0),
	"int8":       reflect.TypeOf(int8(0)),
	"int16":      reflect.TypeOf(int16(0)),
	"int32":      reflect.TypeOf(int32(0)),
	"int64":      reflect.TypeOf(int64(0)),
	"uint":       reflect.TypeOf(uint(0)),
	"uint8":      reflect.TypeOf(uint8(0)),
	"uint16":     reflect.TypeOf(uint16(0)),
	"uint32":     reflect.TypeOf(uint32(0)),
	"uint64":     reflect.TypeOf(uint64(0)),
	"float32":    reflect.TypeOf(float32(0.0)),
	"float64":    reflect.TypeOf(float64(0.0)),
	"complex64":  reflect.TypeOf(complex64(0 + 0i)),
	"complex128": reflect.TypeOf(complex128(0 + 0i)),
}

func JSONBind(filename string) ([]TestCase, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	var rawData []map[string]interface{}
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&rawData); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du JSON: %v", err)
	}

	var testCases []TestCase

	for _, raw := range rawData {
		var testRanges TestRanges
		for _, test := range raw["tests"].([]interface{}) {
			testMap := test.(map[string]interface{})
			args := testMap["args"].([]interface{})
			expectedReturns := testMap["expectedReturns"].([]interface{})

			var testRange TestRange
			for _, rawArg := range args {
				argMap := rawArg.(map[string]interface{})
				value := Value{
					Value: argMap["value"],
					Type:  argMap["type"].(string),
				}
				testRange.Args = append(testRange.Args, convertToType(value.Value, value.Type))
			}

			for _, rawReturn := range expectedReturns {
				returnMap := rawReturn.(map[string]interface{})
				value := Value{
					Value: returnMap["value"],
					Type:  returnMap["type"].(string),
				}
				testRange.ExpectedReturns = append(testRange.ExpectedReturns, convertToType(value.Value, value.Type))
			}

			testRanges = append(testRanges, testRange)
		}
		testCases = append(testCases, TestCase{
			Case:  raw["case"].(string),
			Tests: testRanges,
		})
	}

	return testCases, nil
}

func convertToType(input interface{}, targetType string) interface{} {
	inputType := reflect.TypeOf(input)
	targetTypeType := typesByName[targetType]

	if inputType.ConvertibleTo(targetTypeType) {
		value := reflect.ValueOf(input).Convert(targetTypeType).Interface()
		return value
	}

	return input
}
