package _10_pyramid_

import "testing"

func TestPyramid(t *testing.T) {
	testcases := []struct {
		num int
	}{
		{2},
		{3},
		{5},
	}

	getResult := func(do func(int)) {
		for _, testCase := range testcases {
			do(testCase.num)
		}
	}

	getResult(pyramid)
}
