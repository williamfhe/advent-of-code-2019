package main

type parameterMode int
type opcodeType int

const (
	positionMode  parameterMode = 0
	immediateMode parameterMode = 1

	addCode         opcodeType = 1
	mulCode         opcodeType = 2
	inputCode       opcodeType = 3
	outputCode      opcodeType = 4
	jumpIfTrueCode  opcodeType = 5
	jumpIfFalseCode opcodeType = 6
	lessThanCode    opcodeType = 7
	equalsCode      opcodeType = 8
	stopCode        opcodeType = 99
)
