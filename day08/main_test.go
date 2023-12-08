package main

import (
	"reflect"
	"testing"
)

func TestParseNodes(t *testing.T) {
	input := []string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	expected := map[string][]string{
		"AAA": {"BBB", "BBB"},
		"BBB": {"AAA", "ZZZ"},
		"ZZZ": {"ZZZ", "ZZZ"},
	}
	got, err := parseNodes(&input)
	if !reflect.DeepEqual(got, expected) || err != nil {
		t.Fatalf("test parse nodes fail. expected=%q got=%q", expected, got)
	}
}
