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
		consume:    func(s *PeFsmState, data interface{}) { s.hasOptionalHeader = (data).(uint16) != 0 /* 0, 24, or 28 */ },
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
				return 50
			}

			return 100
		},
		isEndState: false,
	},
	8: { // optionalFileHeader :: Optional Magic Number
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MagicNumber = (data).(uint16)
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
	11: { // optionalFileHeader :: Size of Code
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header.CodeSize = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 12 },
		isEndState: false,
	},
	12: { // optionalFileHeader :: Size of Initialized Data
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 13 },
		isEndState: false,
	},
	13: { // optionalFileHeader :: Size of Uninitialized Data
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 14 },
		isEndState: false,
	},
	14: { // optionalFileHeader :: Address of Entry Point
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.Entrypoint = (data).(uintptr)
		},
		nextState:  func(s *PeFsmState) int { return 15 },
		isEndState: false,
	},
	15: { // optionalFileHeader :: Base of Code
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header.CodeBase = (data).(uintptr)
		},
		nextState:  func(s *PeFsmState) int { return 16 },
		isEndState: false,
	},
	16: { // optionalFileHeader :: Base of Data -- PE32
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header.DataBase = (data).(uintptr)
		},
		nextState:  func(s *PeFsmState) int { return 17 },
		isEndState: false,
	},
	17: { // optionalFileHeader :: Section Alignment
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.SectionAlignment = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 18 },
		isEndState: false,
	},
	18: { // optionalFileHeader :: File Alignment
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.FileAlignment = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 19 },
		isEndState: false,
	},
	19: { // optionalFileHeader :: Major Operating System Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MajorOSVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 20 },
		isEndState: false,
	},
	20: { // optionalFileHeader :: Minor Operating System Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MinorOSVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 21 },
		isEndState: false,
	},
	21: { // optionalFileHeader :: Major Image Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MajorImageVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 22 },
		isEndState: false,
	},
	22: { // optionalFileHeader :: Minor Image Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MinorImageVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 23 },
		isEndState: false,
	},
	23: { // optionalFileHeader :: Major Subsystem Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MajorSubsystemVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 24 },
		isEndState: false,
	},
	24: { // optionalFileHeader :: Minor Subsystem Version
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFileOptionalHeader.MinorSubsystemVersion = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 25 },
		isEndState: false,
	},
	25: { // optionalFileHeader :: Win32VersionValue
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFileOptionalHeader.MinorSubsystemVersion = (data).(uint32)
		},
		nextState:  func(s *PeFsmState) int { return 26 },
		isEndState: false,
	},
	26: { // optionalFileHeader :: Size Of Image
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.ImageSize = (data).(uint64)
		},
		nextState:  func(s *PeFsmState) int { return 27 },
		isEndState: false,
	},
	27: { // optionalFileHeader :: Size Of Headers
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			s.bFile.Header.HeaderSize = (data).(uint64)
		},
		nextState:  func(s *PeFsmState) int { return 28 },
		isEndState: false,
	},
	28: { // optionalFileHeader :: Checksum
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint64)
		},
		nextState:  func(s *PeFsmState) int { return 29 },
		isEndState: false,
	},
	29: { // optionalFileHeader :: Subsystem
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 30 },
		isEndState: false,
	},
	30: { // optionalFileHeader :: DllCharacteristics
		sizeBytes: 2,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState: func(s *PeFsmState) int {
			if s.bFileOptionalHeader.MagicNumber == 0x20b {
				return 41
			}

			return 31
		},
		isEndState: false,
	},
	31: { // optionalFileHeader :: SizeOfStackReserve
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 32 },
		isEndState: false,
	},
	32: { // optionalFileHeader :: SizeOfStackCommit
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 33 },
		isEndState: false,
	},
	33: { // optionalFileHeader :: SizeOfHeapReserve
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 34 },
		isEndState: false,
	},
	34: { // optionalFileHeader :: SizeOfHeapCommit
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 47 },
		isEndState: false,
	},
	41: { // optionalFileHeader :: SizeOfStackReserve
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 42 },
		isEndState: false,
	},
	42: { // optionalFileHeader :: SizeOfStackCommit
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 43 },
		isEndState: false,
	},
	43: { // optionalFileHeader :: SizeOfHeapReserve
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 44 },
		isEndState: false,
	},
	44: { // optionalFileHeader :: SizeOfHeapCommit
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 47 },
		isEndState: false,
	},
	47: { // optionalFileHeader :: LoaderFlags
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 48 },
		isEndState: false,
	},
	48: { // optionalFileHeader :: NumberOfRvaAndSizes
		sizeBytes: 4,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.Header. = (data).(uint16)
		},
		nextState:  func(s *PeFsmState) int { return 49 },
		isEndState: false,
	},
	49: { // optionalHeaderDataDirectories
		sizeBytes: 128,
		transform: func(bs []byte) (interface{}, error) { return global.BytesToVarBigEndian(bs), nil },
		consume: func(s *PeFsmState, data interface{}) {
			// s.bFile.TargetMagicNumber = global.BinaryFileTarget((data).(uint16))
		},
		nextState:  func(s *PeFsmState) int { return 50 },
		isEndState: false,
	},
	50: { // section Headers
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
