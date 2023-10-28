package linker

import (
	"io"

	global "github.com/tflexsoom/gasm/internal/linker/global"
)

type OptionalHeaderPE struct {
	Size                  uint16
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

func GetPEReader() global.BinaryFileReader {
	return (func(io.Reader) (global.BinaryFileResult, error) {
		return BinaryFileResultPE{}, nil
	})
}
