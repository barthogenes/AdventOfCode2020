package day9

import "container/list"

type XMAS struct {
	Preamble     *list.List
	PreambleSize int
}

func (xmas *XMAS) ReadData(value int) int {
	xmas.Preamble.PushBack(value)
	if xmas.Preamble.Len() <= xmas.PreambleSize {
		return 0
	}
	newElement := xmas.Preamble.Back()
	for firstElement := xmas.Preamble.Front(); firstElement != nil; firstElement = firstElement.Next() {
		for secondElement := xmas.Preamble.Front(); secondElement != nil; secondElement = secondElement.Next() {
			if firstElement.Value.(int) == secondElement.Value.(int) {
				continue
			}

			sum := firstElement.Value.(int) + secondElement.Value.(int)
			if sum == newElement.Value.(int) {
				xmas.Preamble.Remove(xmas.Preamble.Front())
				return 0
			}
		}
	}

	return newElement.Value.(int)
}
