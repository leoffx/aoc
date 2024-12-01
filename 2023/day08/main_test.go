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
	expected := NodesToDirections{
		"AAA": {"BBB", "BBB"},
		"BBB": {"AAA", "ZZZ"},
		"ZZZ": {"ZZZ", "ZZZ"},
	}
	got, err := parseNodes(&input)
	if !reflect.DeepEqual(got, expected) || err != nil {
		t.Fatalf("test parse nodes fail. expected=%q got=%q", expected, got)
	}
}

func TestGetStartingNodes(t *testing.T) {
	input := NodesToDirections{
		"11A": {"11B", "XXX"},
		"11B": {"XXX", "11Z"},
		"11Z": {"11B", "XXX"},
		"22A": {"22B", "XXX"},
		"22B": {"22C", "22C"},
		"22C": {"22Z", "22Z"},
		"22Z": {"22B", "22B"},
		"XXX": {"XXX", "XXX"},
	}
	expected := []string{"11A", "22A"}
	got := getStartingNodes(&input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("test get starting nodes fail. expected=%q got=%q", expected, got)
	}
}
