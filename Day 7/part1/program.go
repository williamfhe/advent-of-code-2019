package main

import (
	"log"
)

type program struct {
	instructions []int
	inputStack   []int
	outputStack  []int // where the output values are stored
	instrPointer int   // points to an instruction
	running      bool
}

func newProgram(instructions, inputStack []int) *program {
	return &program{
		instructions: instructions,
		inputStack:   inputStack,
	}

}

// parseInstruction parses the given instruction and returns the matching opcode
// and an array of parameters mode decoded from the instruction
func (p *program) parseInstruction(instruction int) (opcodeType, []parameterMode) {
	var paramsMode []parameterMode

	opcode := opcodeType(instruction % 100)
	instruction /= 100

	for instruction > 0 {
		paramMode := parameterMode(instruction % 10)
		paramsMode = append(paramsMode, paramMode)

		instruction /= 10
	}

	return opcode, paramsMode
}

func (p *program) fillParametesrMode(paramsMode []parameterMode, size int) []parameterMode {
	for i := len(paramsMode); i < size; i++ {
		paramsMode = append(paramsMode, positionMode)
	}

	return paramsMode
}

func (p *program) readValue(index int, mode parameterMode) int {
	var output int

	switch mode {
	case positionMode:
		readIndex := p.instructions[index]
		output = p.instructions[readIndex]
	case immediateMode:
		output = p.instructions[index]
	default:
		log.Fatalf("Unknown parameter mode: %d for value at index: %d\n", mode, index)
	}

	return output
}

func (p *program) writeValue(index, value int) {
	// write is always position mode
	readIndex := p.instructions[index]
	p.instructions[readIndex] = value
}

func (p *program) execute() []int {
	p.running = true

	for p.running {
		instruction := p.instructions[p.instrPointer]
		opcode, paramsMode := p.parseInstruction(instruction)

		switch opcode {
		case addCode:
			p.addInstr(paramsMode)
		case mulCode:
			p.mulInstr(paramsMode)
		case inputCode:
			p.inputInstr()
		case outputCode:
			p.outputInstr(paramsMode)
		case jumpIfTrueCode:
			p.jumpIfTrueInstr(paramsMode)
		case jumpIfFalseCode:
			p.jumpIfFalseInstr(paramsMode)
		case lessThanCode:
			p.lessThanInstr(paramsMode)
		case equalsCode:
			p.equalsInstr(paramsMode)
		case stopCode:
			p.stopInstr()

		default:
			log.Fatalf("Unknown opcode: %d from instruction at index: %d\n", opcode, p.instrPointer)
		}
	}

	return p.outputStack
}
