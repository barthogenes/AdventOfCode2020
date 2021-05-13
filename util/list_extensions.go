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
	sum := 0

	get(l, func(f1, f2 float64) float64 {
		sum += int(f2)
		return 0
	})

	return sum
}

func get(l *list.List, f func(float64, float64) float64) int {
	element := l.Front()
	val, ok := element.Value.(int)
	if !ok {
		return 0
	}
	for ; element != nil; element = element.Next() {
		elVal, ok := element.Value.(int)
		if !ok {
			return val
		}
		val = int(f(float64(val), float64(elVal)))
	}

	return val
}
