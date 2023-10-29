package linker

import (
	"errors"
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
	0: { // ver no
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.bFile.VersionNo = (data).(uint16) },
		nextState:  func(cfs *CoffFsmState) int { return 1 },
		isEndState: false,
	},
	1: { // num of section headers
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.numSectionHeaders = (data).(uint16) },
		nextState:  func(cfs *CoffFsmState) int { return 2 },
		isEndState: false,
	},
	2: { // timestamp
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.bFile.Timestamp = (data).(uint32) },
		nextState:  func(cfs *CoffFsmState) int { return 3 },
		isEndState: false,
	},
	3: { // symbol table starting addr
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.symbolTableStartingAddr = uintptr((data).(uint32)) },
		nextState:  func(cfs *CoffFsmState) int { return 4 },
		isEndState: false,
	},
	4: { // num symbols
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.numSymbols = (data).(uint32) },
		nextState:  func(cfs *CoffFsmState) int { return 5 },
		isEndState: false,
	},
	5: { // optional header size
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume:    func(s *CoffFsmState, data interface{}) { s.hasOptionalHeader = (data).(uint16) == 28 },
		nextState:  func(cfs *CoffFsmState) int { return 6 },
		isEndState: false,
	},
	6: { // flags
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return BytesToVarBigEndian(bs), nil },
		consume: func(s *CoffFsmState, data interface{}) {
			// TODO FLAGS translator
			// flags := (data).(uint16)
			s.bFile.Flags = map[BinaryFileFlag]bool{}
		},
		nextState:  func(cfs *CoffFsmState) int { return 7 },
		isEndState: false,
	},
}

func (fsmState *CoffFsmState) numBytes() int {
	return fsmState.def.sizeBytes
}

func (fsmState *CoffFsmState) transform(bs []byte) (interface{}, error) {
	return fsmState.def.transform(bs)
}

func (fsmState *CoffFsmState) consume(data interface{}) {
	fsmState.def.consume(fsmState, data)
}

func (fsmState *CoffFsmState) nextState() {
	fsmState.def = fsmDefinitions[fsmState.def.nextState(fsmState)]
}

func (fsmState *CoffFsmState) isEndState() bool {
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
			return nil, errors.New(fmt.Sprintf("Not Enough Bytes To Fulfill COFF File! State: %v", fsmState.def))
		}

		result := BinaryFileResultCOFF{
			coffFile: fsmState.bFile,
		}

		return result, nil
	})
}
