package main

func (p *program) addInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 3)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])
	p.writeValue(p.instrPointer+3, v1+v2, paramsMode[2])

	p.instrPointer += 4
}

func (p *program) mulInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 3)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])
	p.writeValue(p.instrPointer+3, v1*v2, paramsMode[2])

	p.instrPointer += 4
}

func (p *program) inputInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 1)

	msg := message{
		msgType: needInputMessage,
		data:    -1,
	}

	p.outputChan <- msg

	val := <-p.inputChan
	p.writeValue(p.instrPointer+1, val, paramsMode[0])

	p.instrPointer += 2
}

func (p *program) outputInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 1)
	val := p.readValue(p.instrPointer+1, paramsMode[0])

	msg := message{
		msgType: dataOutputMessage,
		data:    val,
	}

	p.outputChan <- msg
	p.instrPointer += 2
}

func (p *program) jumpIfTrueInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 2)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])

	if v1 != 0 {
		p.instrPointer = v2
	} else {
		p.instrPointer += 3
	}
}

func (p *program) jumpIfFalseInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 2)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])

	if v1 == 0 {
		p.instrPointer = v2
	} else {
		p.instrPointer += 3
	}
}

func (p *program) lessThanInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 3)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])

	valueToWrite := 0
	if v1 < v2 {
		valueToWrite = 1
	}

	p.writeValue(p.instrPointer+3, valueToWrite, paramsMode[2])
	p.instrPointer += 4
}

func (p *program) equalsInstr(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 3)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	v2 := p.readValue(p.instrPointer+2, paramsMode[1])

	valueToWrite := 0
	if v1 == v2 {
		valueToWrite = 1
	}

	p.writeValue(p.instrPointer+3, valueToWrite, paramsMode[2])
	p.instrPointer += 4
}

func (p *program) adjustRelativeBaseInst(paramsMode []parameterMode) {
	paramsMode = p.fillParametesrMode(paramsMode, 1)
	v1 := p.readValue(p.instrPointer+1, paramsMode[0])
	p.relativeBase += v1
	p.instrPointer += 2
}

func (p *program) stopInstr() {
	p.stopped = true
}
