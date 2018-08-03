//go test -v github.com\sineatos\deag\tools\support -run ^TestMultiChapters$
package support

import (
	"testing"
)

var (
	logbook Logbook
)

func initFunc() {
	logbook = NewDefaultLogbook(10, 10)
}

func TestMultiChapters(t *testing.T) {
	initFunc()
	data1 := Dict{
		"gen":   0,
		"evals": 100,
		"fitness": Dict{
			"obj 1": Dict{
				"avg": float64(1.0),
				"max": float64(10),
			},
			"avg": float64(1.0),
			"max": 10,
		},
		"length": Dict{
			"avg": float64(1.0),
			"max": 30,
		},
		"test": Dict{
			"avg": float64(1.0),
			"max": 20,
		},
	}
	logbook.Record(data1)

	data2 := Dict{
		"gen":   0,
		"evals": 100,
		"fitness": Dict{
			"obj 1": Dict{
				"avg": float64(1.0),
				"max": 10,
			},
		},
		"length": Dict{
			"avg": float64(1.0),
			"max": 30,
		},
		"test": Dict{
			"avg": 1.0,
			"max": 20,
		},
	}
	logbook.Record(data2)
	t.Log("\n" + logbook.String())
}

func TestNoChapters(t *testing.T) {
	initFunc()
	logbook.Record(Dict{
		"gen":   0,
		"evals": 100,
		"avg":   float64(1.0),
		"max":   float64(10),
	})
	t.Log("\n" + logbook.String())
}

func TestOneChapter(t *testing.T) {
	initFunc()
	logbook.Record(Dict{
		"gen":   0,
		"evals": 100,
		"fitness": Dict{
			"avg": 1.0,
			"max": 10,
		},
	})

	logbook.Record(Dict{
		"gen":   1,
		"evals": 100,
		"fitness": Dict{
			"avg": 2.0,
			"max": 20,
		},
	})

	t.Log("\n" + logbook.String())
}

func TestOneBigChapter(t *testing.T) {
	initFunc()
	logbook.Record(Dict{
		"gen":   0,
		"evals": 100,
		"fitness": Dict{
			"obj 1": Dict{
				"avg": 1.0,
				"max": 10,
			},
			"obj 2": Dict{
				"avg": 2.0,
				"max": 20,
			},
		},
	})

	logbook.Record(Dict{
		"gen":   1,
		"evals": 200,
		"fitness": Dict{
			"obj 1": Dict{
				"avg": 3.0,
				"max": 30,
			},
			"obj 2": Dict{
				"avg": 4.0,
				"max": 40,
			},
		},
	})

	t.Log("\n" + logbook.String())
}

func TestLogbookAndStatistics(t *testing.T) {
	initFunc()

	negative := func(val interface{}) interface{} {
		v, ok := val.(int)
		if !ok {
			t.Errorf("the value isn't int, %v", val)
		} else {
			return -v
		}
		return nil
	}

	power2 := func(val interface{}) interface{} {
		v, ok := val.(int)
		if !ok {
			t.Errorf("the value isn't int, %v", val)
		} else {
			return v * v
		}
		return nil
	}

	// test NewDefaultStatistics
	stat1 := NewDefaultMultiStatistics("stat1")
	stat2 := NewDefaultStatistics("stat2", nil)
	stat3 := NewDefaultMultiStatistics("stat3")
	stat4 := NewDefaultStatistics("stat4", nil)
	stat2Name := stat1.AddStats(stat2)
	stat3.AddStats(stat4)
	stat1.AddStats(stat3)
	if stat2Name != "stat2" {
		t.Errorf("target %s, actual %s", "stat2", stat2Name)
	}
	if stat1.GetStats(stat2Name) != stat2 {
		t.Errorf("target %v, actual %v", stat1.GetStats(stat2Name), stat2)
	}
	stat1.Register("power2_1", power2)
	stat2.Register("negative_2", negative)
	stat3.Register("power2_3", power2)
	stat3.Register("negative_3", negative)
	stat4.Register("power2_4", power2)
	stat4.Register("negative_4", negative)
	logbook.Record(stat1.Compile(10))
	t.Log(logbook)
}
