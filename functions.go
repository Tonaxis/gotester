package gotester

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func checkFunction(fn interface{}) bool {
	fnType := reflect.TypeOf(fn)

	if fnType.Kind() != reflect.Func {
		panic("La méthode passée doit être une fonction")
	}

	return true
}

func checkArgs(fn interface{}, args ...interface{}) bool {
	fnType := reflect.TypeOf(fn)

	if fnType.NumIn() != len(args) {
		panic("Le nombre d'arguments ne correspond pas à la méthode")
	}

	return true
}

func convertArgs(args ...interface{}) []reflect.Value {
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	return reflectArgs
}

func convertResult(result []reflect.Value) []interface{} {
	results := make([]interface{}, len(result))
	for i, v := range result {
		results[i] = v.Interface()
	}

	return results
}

func run(fn interface{}, args []reflect.Value) ([]reflect.Value, time.Duration) {
	start := time.Now()
	result := reflect.ValueOf(fn).Call(args)
	end := time.Now()

	elapsed := end.Sub(start)

	return result, elapsed
}

func printResult(fn interface{}, data TestRange, results Result, duration time.Duration, testCheck bool) {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	testResult := "\u001b[0;1;90;102m[ OK  ]\u001b[0m"
	if !testCheck {
		testResult = "\u001b[0;1;97;101m[ERROR]\u001b[0m"
	}
	fmt.Printf("%s \u001b[1;4m%s\u001b[0m(%v) => \u001b[1;93m%v\u001b[0m > Expected : \u001b[1;93m%v\u001b[0m        | in \u001b[1;95m%v\u001b[0mms\n", testResult, fnName, interfaceToString(data.Args...), interfaceToString(results...), interfaceToString(data.ExpectedReturns...), duration.Milliseconds())
}
