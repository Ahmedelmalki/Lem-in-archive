package pathing_test

import (
	"fmt"
	"testing"

	"lemin/parse"
	"lemin/pathing"
)

func TestFindAllPath(t *testing.T) {
	colony := parse.Parse("../parse/test.txt")
	paths := pathing.FindAllPath(colony)
	t.Log(paths)
	t.SkipNow()
}

func TestDiffrent(t *testing.T) {
	type test struct {
		g bool
		d [2][]string
	}
	tests := []test{
		{g: true, d: [2][]string{{"0", "6", "4", "1"}, {"0", "2", "3", "1"}}},
		{g: true, d: [2][]string{{"0", "6", "4", "1"}, {"0", "2", "3", "1"}}},
		{g: false, d: [2][]string{{"0", "6", "4", "1"}, {"0", "2", "4", "1"}}},
		{g: true, d: [2][]string{{"0", "6", "4", "1"}, {"0", "2", "456", "1"}}},
		{g: true, d: [2][]string{{"0", "4", "3", "1"}, {"0", "6", "6", "1"}}},
	}
	for i, test := range tests {
		if pathing.Diffrent(&test.d[0], &test.d[1]) != test.g {
			t.Logf("failed test nbr:%d with data :%v", i, test)
			t.Fail()
		}
	}
}

func TestWcov(t *testing.T) {
	type test struct {
		g [][]string
		d [][]string
	}
	tests := []test{
		{g: [][]string{{"0", "2", "3", "1"}, {"0", "6", "4", "1"}}, d: [][]string{{"0", "2", "4", "1"}, {"0", "2", "3", "1"}, {"0", "6", "4", "1"}, {"0", "6", "4", "2", "3", "1"}}},
		{
			[][]string{
				{"0", "2", "3", "1"},
				{"0", "6", "4", "1"},
				{"0", "7", "5", "1"},
			},
			[][]string{
				{"0", "2", "3", "1"},
				{"0", "6", "4", "1"},
				{"0", "7", "4", "1"},
				{"0", "7", "5", "1"},
				{"0", "8", "9", "3", "1"},
			},
		},
	}
	for i, test := range tests {
		res := pathing.Wcov(test.d)
		if fmt.Sprint(res) != fmt.Sprint(test.g) {
			t.Logf("failed test nbr:%d with output :%v", i, res)
			t.Fail()
		}
	}
}
