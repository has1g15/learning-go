// +build integration

package Testing

import "testing"

type AddResult struct {
	x int
	y int
	expected int
}

var addResults = []AddResult{
	{1,2,3},
	{2,3,5},
	{3,4,7},
	{4,5,9},
}

func TestAddNums(t *testing.T) {
	for _, test := range addResults {
		result := AddNums(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}

//go test -tags=integration