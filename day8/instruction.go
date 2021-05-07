package day8

import (
	"regexp"

	"github.com/barthogenes/adventofcode2020/util"
)

type InstructionType int

const (
	NoOPeration InstructionType = iota
	Jump
	Accumulate
)

func (it InstructionType) String() string {
	return [...]string{"NoOPeration"}[it]
}

type Instruction struct {
	Type     InstructionType
	Argument int
}

func (i *Instruction) SetFromString(instruction string) {
	instructionRegex := regexp.MustCompile(`(\w*) ([+|-]\d*)`)
	match := instructionRegex.FindStringSubmatch(instruction)
	i.SetTypeFromString(match[1])
	i.SetArgumentFromString(match[2])
}

func (i *Instruction) SetTypeFromString(instructionType string) {
	if instructionType == "nop" {
		i.Type = NoOPeration
	}

	if instructionType == "jmp" {
		i.Type = Jump
	}

	if instructionType == "acc" {
		i.Type = Accumulate
	}
}

func (i *Instruction) SetArgumentFromString(argument string) {
	i.Argument = util.ToNumber(argument)
}
