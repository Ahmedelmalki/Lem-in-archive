package pathing_test

import (
	"testing"

	"lemin/parse"
	"lemin/pathing"
)

func TestFindAllPath(t *testing.T) {
	colony := parse.Parse("../test.txt")
	paths := pathing.FindAllPath(colony)
	t.Log(paths)
	t.SkipNow()
}
