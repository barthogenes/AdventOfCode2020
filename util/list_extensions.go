package util

import (
	"container/list"
	"math"
)

func GetMax(l *list.List) int {
	return get(l, math.Max)
}

func GetMin(l *list.List) int {
	return get(l, math.Min)
}

func GetSum(l *list.List) int {
	sum := float64(0)

	get(l, func(f1, f2 float64) float64 {
		sum = f1 + f2
		return sum
	})

	return int(sum)
}

func get(l *list.List, f func(float64, float64) float64) int {
	val := 0
	firstValOk := false
	for element := l.Front(); element != nil; element = element.Next() {
		for !firstValOk {
			val, firstValOk = element.Value.(int)
			element = element.Next()
		}

		elVal, ok := element.Value.(int)
		if !ok {
			continue
		}
		val = int(f(float64(val), float64(elVal)))
	}

	return val
}
