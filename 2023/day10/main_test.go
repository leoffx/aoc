package main

import (
	"reflect"
	"testing"
)

func TestCalculateFarthestPointDistance(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected int
	}{
		{
			input: [][]string{
				{".", ".", ".", ".", "."},
				{".", "S", "-", "7", "."},
				{".", "|", ".", "|", "."},
				{".", "L", "-", "J", "."},
				{".", ".", ".", ".", "."},
			},
			expected: 4,
		},
		{
			input: [][]string{
				{".", ".", "F", "7", "."},
				{".", "F", "J", "|", "."},
				{"S", "J", ".", "L", "7"},
				{"|", "F", "-", "-", "J"},
				{"L", "J", ".", ".", "."},
			},
			expected: 8,
		},
	}

	for i, tt := range tests {
		got, _ := calculateFarthestPointDistance(tt.input)
		if tt.expected != got {
			t.Fatalf("tests[%d] calculateFarthestPointDistance fail. expected=%d got=%d", i, tt.expected, got)
		}
	}
}
func TestFindConnectedPipes(t *testing.T) {
	tests := []struct {
		pipes    [][]string
		i        int
		j        int
		expected coords
	}{
		{
			pipes: [][]string{
				{".", ".", ".", ".", "."},
				{".", "S", "-", "7", "."},
				{".", "|", ".", "|", "."},
				{".", "L", "-", "J", "."},
				{".", ".", ".", ".", "."},
			},
			i:        1,
			j:        1,
			expected: coords{x: 2, y: 1},
		},
		{
			pipes: [][]string{
				{".", ".", "F", "7", "."},
				{".", "F", "J", "|", "."},
				{"S", "J", ".", "L", "7"},
				{"|", "F", "-", "-", "J"},
				{"L", "J", ".", ".", "."},
			},
			i:        0,
			j:        2,
			expected: coords{x: 1, y: 2},
		},
	}

	for i, tt := range tests {
		got := findConnectedPipes(tt.pipes, tt.i, tt.j)
		if !reflect.DeepEqual(tt.expected, got) {
			t.Fatalf("tests[%d] calculateFarthestPointDistance fail. expected=%d got=%d", i, tt.expected, got)
		}
	}
}
