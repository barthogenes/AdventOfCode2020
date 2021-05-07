package day8

type InstructionType int

const (
	NoOPeration InstructionType = iota
	test
)

func (it InstructionType) String() string {
	return [...]string{"NoOPeration"}[it]
}

type Instruction struct {
	Type     InstructionType
	Argument int
}
