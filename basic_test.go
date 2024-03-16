package main

import (
	"os"
	"testing"
)

var testfiles = [12]string{
	"measurements-10000-unique-keys",
	"measurements-10",
	"measurements-1",
	"measurements-20",
	"measurements-2",
	"measurements-3",
	"measurements-boundaries",
	"measurements-complex-utf8",
	"measurements-dot",
	"measurements-rounding",
	"measurements-shortest",
	"measurements-short",
}

var approaches = map[string](func(*os.File) string){
	"naive":   Naive,
	"builder": StringBuilder,
}

func TestApproaches(t *testing.T) {
	for _, filename := range testfiles {
		expectedBytes, err := os.ReadFile("samples/" + filename + ".out")
		if err != nil {
			panic(err)
		}
		expected := string(expectedBytes)

		for name, function := range approaches {
			t.Run(filename+" ["+name+"]", func(t *testing.T) {
				file, err := os.Open("samples/" + filename + ".txt")
				if err != nil {
					panic(err)
				}
				defer file.Close()

				result := function(file)
				if result != expected {
					t.Errorf("Expected %s, got %s", expected, result)
				}
			})
		}
	}
}

func TestLarge(t *testing.T) {
	file, err := os.Open("bench/1_000_000.txt")
	if err != nil {
		panic(err)
	}
	expected := Naive(file)
	file.Close()

	for name, function := range approaches {
		if name == "naive" {
			continue
		}

		t.Run("["+name+"]", func(t *testing.T) {
			file, err := os.Open("bench/1_000_000.txt")
			if err != nil {
				panic(err)
			}
			result := function(file)
			file.Close()

			if result != expected {
				t.Errorf("Expected %s, got %s", expected, result)
			}
		})
	}
}

var benchfiles = [4]string{
	"100_000",
	"1_000_000",
	"10_000_000",
	"100_000_000",
}

func BenchmarkApproaches(b *testing.B) {
	for _, filename := range benchfiles {
		for name, function := range approaches {
			b.Run(filename+" ["+name+"]", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					file, err := os.Open("bench/" + filename + ".txt")
					if err != nil {
						panic(err)
					}

					_ = function(file)

					file.Close()
				}
			})
		}
	}
}
