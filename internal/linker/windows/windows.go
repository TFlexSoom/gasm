package linker

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
