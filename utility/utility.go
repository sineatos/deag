package utility

import (
	"fmt"
)

// If equals to (condition ? trueVal : falseVal)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// Interface2Float64Slice returns a float64 slice full with value, and name is used to explain the slice's meanings
func Interface2Float64Slice(name string, value interface{}, size int) []float64 {
	var values []float64
	switch value.(type) {
	case float64:
		values = make([]float64, size)
		num := value.(float64)
		for i := 0; i < size; i++ {
			values[i] = num
		}
	case []float64:
		values = value.([]float64)
		if len(values) < size {
			panic(fmt.Sprintf("%s must be at least the size of individual: %d < %d", name, len(values), size))
		}
	default:
		panic(fmt.Sprintf("%s must be float64 or []float64", name))
	}
	return values
}

// Interface2IntSlice returns a int slice full with value, and name is used to explain the slice's meanings
func Interface2IntSlice(name string, value interface{}, size int) []int {
	var values []int
	switch value.(type) {
	case int:
		values = make([]int, size)
		num := value.(int)
		for i := 0; i < size; i++ {
			values[i] = num
		}
	case []int:
		values = value.([]int)
		if len(values) < size {
			panic(fmt.Sprintf("%s must be at least the size of individual: %d < %d", name, len(values), size))
		}
	default:
		panic(fmt.Sprintf("%s must be int or []int", name))
	}
	return values
}

// CollectFloat64 collects float64 values and returns as slice
func CollectFloat64(values ...float64) []float64 {
	ans := make([]float64, len(values))
	for i, value := range values {
		ans[i] = value
	}
	return ans
}
