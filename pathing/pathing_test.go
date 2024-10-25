package pathing_test

import (
	"fmt"
	"testing"

	"lemin/pathing"
	"lemin/static"
)

func TestFindAllPath(t *testing.T) {
	c := static.Colony{
		Rooms: map[string][]string{
			"0": {"2", "6"},
			"1": {"3", "4"},
			"2": {"0", "3", "4"},
			"3": {"2", "1"},
			"4": {"6", "1", "2"},
			"6": {"0", "4"},
		},
		Start:  "0",
		Finish: "1",
	}
	res := pathing.FindAllPaths(&c)
	if fmt.Sprint(res) != fmt.Sprint([][]string{{"0", "2", "3", "1"}, {"0", "6", "4", "1"}}) {
		t.Fail()
	}
}

func TestRemoveRepetition(t *testing.T) {
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
		res := pathing.RemoveRepetition(test.d)
		if fmt.Sprint(res) != fmt.Sprint(test.g) {
			t.Logf("failed test nbr:%d with output :%v", i, res)
			t.Fail()
		}
	}
}

func TestBacktrack(t *testing.T) {
	c := static.Colony{
		Rooms: map[string][]string{
			"0": {"2", "6"},
			"1": {"3", "4"},
			"2": {"0", "3", "4"},
			"3": {"2", "1"},
			"4": {"6", "1", "2"},
			"6": {"0", "4"},
		},
		Start:  "0",
		Finish: "1",
	}
	p := pathing.BFS(&c)
	good := [][]string{
		{"0", "2", "3", "1"},
		{"0", "2", "4", "1"},
		{"0", "6", "4", "1"},
		{"0", "6", "4", "2", "3", "1"},
	}
	if len(p) != len(good) {
		t.Errorf("wrong number of paths\npaths : %v\nwant  :%v\n", p, good)
		t.Fail()
		return
	}
	for i := range p {
		if fmt.Sprint(p[i]) != fmt.Sprint(good[i]) {
			t.Error("wrong path", p[i], good[i])
			t.Fail()
		}
	}
}
