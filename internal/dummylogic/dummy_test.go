package dummylogic_test

import (
	"demo/internal/dummylogic"
	"testing"
)

func TestDummyLogic(t *testing.T) {
	type testcase struct {
		Name           string
		Input          int
		ExpectedOutput string
	}

	testcases := []testcase{
		{
			Name:           "less than 10",
			Input:          5,
			ExpectedOutput: "less than 10",
		},
		// {
		// 	Name:           "less than 20",
		// 	Input:          15,
		// 	ExpectedOutput: "less than 20",
		// },
		{
			Name:           "more than 20",
			Input:          25,
			ExpectedOutput: "more than 20",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			output := dummylogic.DummyLogic(tc.Input)
			if output != tc.ExpectedOutput {
				t.Errorf("Expected %s, got %s", tc.ExpectedOutput, output)
			}
		})
	}
}
