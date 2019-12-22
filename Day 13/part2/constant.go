package main

type parameterMode int
type opcodeType int
type messageType int

const (
	positionMode  parameterMode = 0
	immediateMode parameterMode = 1
	relativeMode  parameterMode = 2

	addCode                opcodeType = 1
	mulCode                opcodeType = 2
	inputCode              opcodeType = 3
	outputCode             opcodeType = 4
	jumpIfTrueCode         opcodeType = 5
	jumpIfFalseCode        opcodeType = 6
	lessThanCode           opcodeType = 7
	equalsCode             opcodeType = 8
	adjustRelativeBaseCode opcodeType = 9
	stopCode               opcodeType = 99

	dataOutputMessage messageType = iota
	needInputMessage  messageType = iota
)
