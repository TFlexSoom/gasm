package linker

import (
	"fmt"
	"io"

	global "github.com/tflexsoom/gasm/internal/linker/global"
)

type OptionalHeaderPE struct {
	MagicNumber           uint16
	MajorLinkerVersion    uint8
	MinorLinkerVersion    uint8
	MajorOSVersion        uint16
	MinorOSVersion        uint16
	MajorImageVersion     uint16
	MinorImageVersion     uint16
	MajorSubsystemVersion uint16
	MinorSubsystemVersion uint16
	Win32VersionValue     uint64
	Checksum              uint64
	Subsystem             WindowsSubsystem
	DllCharacteristics    []DllCharacteristic
	SizeOfStackReserve    uint64
	SizeOfHeapReserve     uint64
	SizeOfHeapCommit      uint64
	LoaderFlags           uint64
	NumberOfRvaAndSizes   uint64
}

type BinaryFileResultPE struct {
	peHeader OptionalHeaderPE
	coffFile global.BinaryFile
}

func (bfrPE BinaryFileResultPE) Result() *global.BinaryFile {
	return &bfrPE.coffFile
}

type PeFsmState struct {
	bFileOptionalHeader     OptionalHeaderPE
	bFile                   global.BinaryFile
	numSectionHeaders       uint16
	symbolTableStartingAddr uintptr
	numSymbols              uint32
	hasOptionalHeader       bool
	def                     PeFsmDefinition
}

type PeFsmDefinition struct {
	sizeBytes  int
	transform  func([]byte) (interface{}, error)
	consume    func(*PeFsmState, interface{})
	nextState  func(*PeFsmState) int
	isEndState bool
}

var fsmDefinitions = map[int]PeFsmDefinition{
	0: { // ver no
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.bFile.VersionNo = (data).(uint16) },
		nextState:  func(s *PeFsmState) int { return 1 },
		isEndState: false,
	},
	1: { // num of section headers
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.numSectionHeaders = (data).(uint16) },
		nextState:  func(s *PeFsmState) int { return 2 },
		isEndState: false,
	},
	2: { // timestamp
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.bFile.Timestamp = (data).(uint32) },
		nextState:  func(s *PeFsmState) int { return 3 },
		isEndState: false,
	},
	3: { // symbol table starting addr
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.symbolTableStartingAddr = uintptr((data).(uint32)) },
		nextState:  func(s *PeFsmState) int { return 4 },
		isEndState: false,
	},
	4: { // num symbols
		sizeBytes:  4,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.numSymbols = (data).(uint32) },
		nextState:  func(s *PeFsmState) int { return 5 },
		isEndState: false,
	},
	5: { // optional header size
		sizeBytes:  2,
		transform:  func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume:    func(s *PeFsmState, data interface{}) { s.hasOptionalHeader = (data).(uint16) == 28 },
		nextState:  func(s *PeFsmState) int { return 6 },
		isEndState: false,
	},
	6: { // flags
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			flags := (data).(uint16)
			s.bFile.Flags = make(map[global.BinaryFileFlag]bool)
			global.VariableToFlagValues(flags, func(flag uint16) { s.bFile.Flags[global.BinaryFileFlag(flag)] = true })
		},
		nextState:  func(s *PeFsmState) int { return 7 },
		isEndState: false,
	},
	7: { // targetId
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.TargetMagicNumber = global.BinaryFileTarget((data).(uint16))
		},
		nextState: func(s *PeFsmState) int {
			if s.hasOptionalHeader {
				return 8
			}

			if s.numSectionHeaders > 0 {
				return 20
			}

			return 100
		},
		isEndState: false,
	},
	8: { // optionalFileHeader :: Optional Magic Number
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.FileMagicNumber = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 9 },
		isEndState: false,
	},
	9: { // optionalFileHeader :: Major Linker Version
		sizeBytes: 1,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MajorLinkerVersion = (data).(uint8)
		},
		nextState:  func(s *PeFsmState) int { return 10 },
		isEndState: false,
	},
	10: { // optionalFileHeader :: Minor Linker Version
		sizeBytes: 1,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MinorLinkerVersion = (data).(uint8)
		},
		nextState:  func(s *PeFsmState) int { return 11 },
		isEndState: false,
	},
	11: { // optionalFileHeader :: Minor Linker Version
		sizeBytes: 1,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header. = (data).(uint8)
		},
		nextState:  func(s *PeFsmState) int { return 11 },
		isEndState: false,
	},
	20: { // section Headers
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.TargetMagicNumber = global.BinaryFileTarget((data).(uint16))
		},
		nextState:  func(s *PeFsmState) int { return 9 },
		isEndState: false,
	},
	100: { // End State
		sizeBytes:  0,
		transform:  nil,
		consume:    nil,
		nextState:  nil,
		isEndState: true,
	},
	/* ----------------------------- END TODO -------------------------------- */
}

func (fsmState *PeFsmState) NumBytes() int {
	return fsmState.def.sizeBytes
}

func (fsmState *PeFsmState) Transform(bs []byte) (interface{}, error) {
	return fsmState.def.transform(bs)
}

func (fsmState *PeFsmState) Consume(data interface{}) {
	fsmState.def.consume(fsmState, data)
}

func (fsmState *PeFsmState) NextState() {
	fsmState.def = fsmDefinitions[fsmState.def.nextState(fsmState)]
}

func (fsmState *PeFsmState) IsEndState() bool {
	return fsmState.def.isEndState
}

func GetPEReader() global.BinaryFileReader {
	return (func(reader io.Reader) (global.BinaryFileResult, error) {
		fsmState := PeFsmState{}

		err := global.ReadIntoFSM(reader, (&fsmState), global.Fsm)
		if err != nil {
			return nil, err
		}

		if !fsmState.def.isEndState {
			return nil, fmt.Errorf("not enough bytes to fulfill coff file! state: %v", fsmState.def)
		}

		result := BinaryFileResultPE{
			coffFile: fsmState.bFile,
		}

		return result, nil
	})
}
