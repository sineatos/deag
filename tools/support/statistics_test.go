package support

import (
	"testing"
)

var (
	stat    Statistics
	mulStat MultiStatistics
)

func TestDefaultStatistics(t *testing.T) {

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
	intStat := NewDefaultStatistics("default statistics", nil)
	// test GetName()
	if intStat.GetName() != "default statistics" {
		t.Errorf("target: %v, actual: %v", "default statistics", intStat.GetName())
	}
	// test Register
	intStat.Register("negative", negative)
	intStat.Register("power2", power2)
	// test Compile
	input1 := 10
	ans1 := intStat.Compile(input1)
	if val, ok := ans1["negative"]; !ok {
		t.Error("the compiled value of negative doesn't exist")
	} else if val.(int) != -input1 {
		t.Errorf("target of negative is %v, actual got %v", -input1, val)
	}

	if val, ok := ans1["power2"]; !ok {
		t.Error("the compiled value of power2 doesn't exist")
	} else if val.(int) != input1*input1 {
		t.Errorf("target of negative is %v, actual got %v", input1*input1, val)
	}
}

func TestDefaultMultiStatistics(t *testing.T) {

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
	ans1 := stat1.Compile(10)
	t.Logf("\n%v", ans1)
	t.Logf("\n%v", stat1.GetAllStats())
	stat1.RemoveStats("stat2")
	ans2 := stat1.Compile(10)
	t.Log("-----------------------------------------------------------")
	t.Log("after remove stat2")
	t.Logf("\n%v", ans2)
	t.Logf("\n%v", stat1.GetAllStats())

}
