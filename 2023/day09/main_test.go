package main

import "testing"

func TestCalculateNextSequenceValue(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}
	for i, tt := range tests {
		got := calculateNextSequenceValue(&tt.input)
		if tt.expected != got {
			t.Fatalf("tests[%d] calculateNextSequenceValue fail. expected=%d got=%d", i, tt.expected, got)
		}
	}
}
