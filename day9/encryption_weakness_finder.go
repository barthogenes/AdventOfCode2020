package day9

import (
	"container/list"

	"github.com/barthogenes/adventofcode2020/util"
)

type EncryptionWeaknessFinder struct {
	InvalidNumber int
}

func (ewf *EncryptionWeaknessFinder) FindWeakness(data []int) int {
	weaknessList := list.New()
	for _, value := range data {
		weaknessList.PushBack(value)
		sum := util.GetSum(weaknessList)
		for sum > ewf.InvalidNumber {
			weaknessList.Remove(weaknessList.Front())
			sum = util.GetSum(weaknessList)
		}
		if sum == ewf.InvalidNumber {
			max := util.GetMax(weaknessList)
			min := util.GetMin(weaknessList)
			return max + min
		}
	}

	return 0
}
