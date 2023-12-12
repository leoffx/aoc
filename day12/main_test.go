package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateAllPossibilities(t *testing.T) {
	tests := []struct {
		springs []string
		want    [][]string
	}{
		{
			springs: []string{"?"},
			want:    [][]string{{"#"}, {"."}},
		},
		{
			springs: []string{"?", "#"},
			want:    [][]string{{"#", "#"}, {".", "#"}},
		},
		{
			springs: []string{"?", "?"},
			want:    [][]string{{"#", "#"}, {"#", "."}, {".", "#"}, {".", "."}},
		},
		{
			springs: []string{"?", "?", "?"},
			want: [][]string{
				{"#", "#", "#"},
				{"#", "#", "."},
				{"#", ".", "#"},
				{".", "#", "#"},
				{"#", ".", "."},
				{".", "#", "."},
				{".", ".", "#"},
				{".", ".", "."},
			},
		},
	}
	sortSlices := func(slices [][]string) {
		sort.Slice(slices, func(i, j int) bool {
			for x := range slices[i] {
				if slices[i][x] == slices[j][x] {
					continue
				}
				return slices[i][x] < slices[j][x]
			}
			return false
		})
	}

	for i, tt := range tests {
		got := generateAllPossibilities(tt.springs)
		sortSlices(got)
		sortSlices(tt.want)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("tests[%d] generateAllPossibilities fail. got=%v want=%v", i, got, tt.want)
		}
	}

}
