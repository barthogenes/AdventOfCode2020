package day8

type GameConsole struct {
	Accumulator  int
	Instructions []Instruction
	LineHitCount map[int]int
}

func (gc *GameConsole) Start() {

}
