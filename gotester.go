package gotester

import "time"

type TestCase struct {
	Case  string     `json:"case"`
	Tests TestRanges `json:"tests"`
}

type TestRange struct {
	Args            []interface{} `json:"args"`
	ExpectedReturns []interface{} `json:"expectedReturns"`
}

type TestRanges []TestRange

type Value struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type Result []interface{}

type CompleteResult struct {
	Method string
	Cases  []struct {
		Name  string
		Tests []struct {
			State           bool
			Args            []interface{}
			ExpectedReturns []interface{}
			Result          Result
			Duration        time.Duration
		}
	}
}

type UnitResult struct {
	Result   Result
	Duration time.Duration
}
