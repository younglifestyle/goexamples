package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"123abc", "cba321"},
	}

	getResult := func(do func(s1 string) string) {
		for _, testCase := range testCases {
			if testCase[0] != do(testCase[1]) {
				t.Error("expect", testCase[0], "equals", do(testCase[1]))
			}
		}
	}

	getResult(reverseByByte)
}
