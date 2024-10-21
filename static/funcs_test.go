package static_test

import (
	"lemin/static"
	"testing"
)

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
		{g: false, d: [2][]string{{"0", "azer", "1"}, {"0", "azer", "1"}}},
	}
	for i, test := range tests {
		if static.Diffrent(&test.d[0], &test.d[1]) != test.g {
			t.Logf("failed test nbr:%d with data :%v", i, test)
			t.Fail()
		}
	}
}

func TestIsIn(t *testing.T) {
	type test struct {
		g bool
		i string
		s []string
	}
	tests := []test{
		{g: true, i: "2", s: []string{"0", "2", "3", "1"}},
		{g: false, i: "4", s: []string{"0", "2", "3", "1"}},
		{g: true, i: "azer", s: []string{"0", "azer", "1"}},
		{g: false, i: "123", s: []string{"0", "2", "3", "1"}},
		{g: false, i: "0", s: []string{"azer", "1"}},
		{g: true, i: "0", s: []string{"0", "2", "3", "1"}},
		{g: false, i: "6", s: []string{"0", "2", "3", "1"}},
	}
	for i, test := range tests {
		if static.IsIn(test.i, test.s) != test.g {
			t.Logf("failed test nbr:%d with data :%v", i, test)
			t.Fail()
		}
	}
}
