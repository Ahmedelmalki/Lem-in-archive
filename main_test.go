package main_test

import (
	"lemin/parse"
	"lemin/pathing"
	"os"
	"testing"
)

var cpuProfile *os.File

// func TestMain(m *testing.M) {
// 	var err error
// 	cpuProfile, err = os.Create("cpu.prof")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cpuProfile.Close()

// 	pprof.StartCPUProfile(cpuProfile)
// 	defer pprof.StopCPUProfile()

// 	os.Exit(m.Run())
// }

func BenchmarkFindAllPath(b *testing.B) {

	colony := parse.Parse("./parse/test.txt") // Setup
	b.ResetTimer()                            // Start timing after setup

	for i := 0; i < b.N; i++ {
		pathing.FindAllPath(colony, "0", "1") // Benchmark the function
	}
}
