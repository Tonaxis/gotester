package gotester

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func TestUnit(fn interface{}, data TestRange) (bool, UnitResult) {
	if !checkFunction(fn) || !checkArgs(fn, data.Args...) {
		return false, UnitResult{}
	}

	reflectArgs := convertArgs(data.Args...)

	result, timer := run(fn, reflectArgs)

	results := convertResult(result)

	testResult := test(data, results)

	printResult(fn, data, results, timer, testResult)

	unitResult := UnitResult{
		Result:   results,
		Duration: timer,
	}

	return testResult, unitResult
}

// func TestMethod(fn interface{}, data TestRanges) {
// 	fmt.Println()
// 	fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;4;96m%v\u001b[0m", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()), "\u001b[0;97m=\u001b[0m", 100))

// 	count := 0
// 	for _, d := range data {
// 		ok, _ := TestUnit(fn, d)
// 		if ok {
// 			count++
// 		}
// 	}

// 	fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;96m%v/%v\u001b[0m", count, len(data)), "\u001b[0;97m=\u001b[0m", 100))
// 	fmt.Println()
// }

func TestMethodCases(fn interface{}, data []TestCase) CompleteResult {
	methodName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	fmt.Println()
	fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;4;96m%v\u001b[0m", methodName), "\u001b[0;97m=\u001b[0m", 102))

	completeResult := CompleteResult{
		Method: methodName,
	}

	total := 0
	totalCount := 0
	for i, r := range data {
		fmt.Println()
		fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;4;32m[%v]\u001b[0m", r.Case), "\u001b[0;90m=\u001b[0m", 102))

		completeResult.Cases = append(completeResult.Cases, struct {
			Name  string
			Tests []struct {
				State           bool
				Args            []interface{}
				ExpectedReturns []interface{}
				Result          Result
				Duration        time.Duration
			}
		}{
			Name: r.Case,
		})

		count := 0
		for _, d := range r.Tests {
			ok, res := TestUnit(fn, d)
			if ok {
				count++
			}

			completeResult.Cases[i].Tests = append(completeResult.Cases[i].Tests, struct {
				State           bool
				Args            []interface{}
				ExpectedReturns []interface{}
				Result          Result
				Duration        time.Duration
			}{
				State:           ok,
				Args:            d.Args,
				ExpectedReturns: d.ExpectedReturns,
				Result:          res.Result,
				Duration:        res.Duration,
			})
		}

		fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;32m[%v/%v]\u001b[0m", count, len(r.Tests)), "\u001b[0;90m=\u001b[0m", 100))
		fmt.Println()

		totalCount += count
		total += len(r.Tests)
	}

	fmt.Println(centerText(fmt.Sprintf("\u001b[0;1;96m%v/%v\u001b[0m", totalCount, total), "\u001b[0;97m=\u001b[0m", 100))
	fmt.Println()

	return completeResult
}
