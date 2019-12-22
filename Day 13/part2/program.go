package main

import (
	"log"
)

type message struct {
	msgType messageType
	data    int
}

type program struct {
	instructions map[int]int
	inputChan    <-chan int
	outputChan   chan message
	relativeBase int
	instrPointer int // points to an instruction
	stopped      bool
}

func newProgram(instructions map[int]int) (*program, chan int, <-chan message) {
	inputChan := make(chan int)
	outputChan := make(chan message)
	prog := &program{
		instructions: instructions,
		inputChan:    inputChan,
		outputChan:   outputChan,
	}

	return prog, inputChan, outputChan
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
	readIndex := -1
	switch mode {
	case positionMode:
		readIndex = p.instructions[index]
	case immediateMode:
		readIndex = index
	case relativeMode:
		readIndex = p.instructions[index] + p.relativeBase
	default:
		log.Fatalf("Unknown parameter mode: %d for value at index: %d\n", mode, index)
	}

	if readIndex < 0 {
		log.Fatalf("Trying to write at %d - negative memory address!", index)
	}

	return p.instructions[readIndex]
}

func (p *program) writeValue(index, value int, mode parameterMode) {
	writeIndex := -1
	if mode == positionMode {
		writeIndex = p.instructions[index]
	} else if mode == relativeMode {
		writeIndex = p.instructions[index] + p.relativeBase
	} else {
		log.Fatalf("Unknown mode: %d for value: %d at %d\n", mode, value, index)
	}

	if writeIndex < 0 {
		log.Fatalf("Trying to write at %d - negative memory address!", index)
	}

	p.instructions[writeIndex] = value
}

func (p *program) run() {
	for !p.stopped {
		instruction := p.instructions[p.instrPointer]
		opcode, paramsMode := p.parseInstruction(instruction)

		switch opcode {
		case addCode:
			p.addInstr(paramsMode)
		case mulCode:
			p.mulInstr(paramsMode)
		case inputCode:
			p.inputInstr(paramsMode)
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
		case adjustRelativeBaseCode:
			p.adjustRelativeBaseInst(paramsMode)
		case stopCode:
			p.stopInstr()

		default:
			log.Fatalf("Unknown opcode: %d from instruction at index: %d\n", opcode, p.instrPointer)
		}
	}

	close(p.outputChan)
}
