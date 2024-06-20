package linker

import (
	"fmt"
	"io"
)

type BinaryFileResultCOFF struct {
	coffFile BinaryFile
}

func (bfrPE BinaryFileResultCOFF) Result() *BinaryFile {
	return &bfrPE.coffFile
}

type CoffFsmState struct {
	bFile                   BinaryFile
	numSectionHeaders       uint16
	symbolTableStartingAddr uintptr
	numSymbols              uint32
	hasOptionalHeader       bool
	def                     CoffFsmDefinition
}

type CoffFsmDefinition struct {
	sizeBytes  int
	transform  func([]byte) (interface{}, error)
	consume    func(*CoffFsmState, interface{})
	nextState  func(*CoffFsmState) int
	isEndState bool
}

var fsmDefinitions = map[int]CoffFsmDefinition{
	0: { // TODO
		// No Definition since these files are hard to comeby without specific compilers from TI-84
		sizeBytes:  0,
		transform:  nil,
		consume:    nil,
		nextState:  nil,
		isEndState: true,
	},
}

func (fsmState *CoffFsmState) NumBytes() int {
	return fsmState.def.sizeBytes
}

func (fsmState *CoffFsmState) Transform(bs []byte) (interface{}, error) {
	return fsmState.def.transform(bs)
}

func (fsmState *CoffFsmState) Consume(data interface{}) {
	fsmState.def.consume(fsmState, data)
}

func (fsmState *CoffFsmState) NextState() {
	fsmState.def = fsmDefinitions[fsmState.def.nextState(fsmState)]
}

func (fsmState *CoffFsmState) IsEndState() bool {
	return fsmState.def.isEndState
}

func GetCOFFReader() BinaryFileReader {
	return (func(reader io.Reader) (BinaryFileResult, error) {
		fsmState := CoffFsmState{}

		err := ReadIntoFSM(reader, (&fsmState), Fsm)
		if err != nil {
			return nil, err
		}

		if !fsmState.def.isEndState {
			return nil, fmt.Errorf("not enough bytes to fulfill coff file! state: %v", fsmState.def)
		}

		result := BinaryFileResultCOFF{
			coffFile: fsmState.bFile,
		}

		return result, nil
	})
}
