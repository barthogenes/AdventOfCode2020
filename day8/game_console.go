package day8

type GameConsole struct {
	Accumulator  int
	Instructions []Instruction // I could combine the LineHitCount map with the Instructions slice but then we would lose the order as Go does not have ordered maps.
	LineHitCount map[int]int
}

func (gc *GameConsole) Start() int {
	gc.LineHitCount = make(map[int]int)

	i := 0
	for i < len(gc.Instructions) {
		_, exists := gc.LineHitCount[i+1]
		if exists {
			return gc.Accumulator
		}
		gc.LineHitCount[i+1] = 1

		instruction := gc.Instructions[i]
		if instruction.Type == Jump {
			i += instruction.Argument
			continue
		}

		if instruction.Type == Accumulate {
			gc.Accumulator += instruction.Argument
		}
		i++
	}

	return gc.Accumulator
}
