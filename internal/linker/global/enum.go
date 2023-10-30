package linker

type BinaryFileFlag uint16

const (
	BFF_STRIPPED_RELOCATION       BinaryFileFlag = 0x0001
	BFF_IS_RELOCATABLE            BinaryFileFlag = 0x0002
	BFF_STRIPPED_LINE_NUMBERS     BinaryFileFlag = 0x0004
	BFF_STRIPPED_SYMBOLS          BinaryFileFlag = 0x0008
	BFF_AGGRESSIVE_WS_TRIM        BinaryFileFlag = 0x0010
	BFF_LARGE_ADDRESS_AWARE       BinaryFileFlag = 0x0020
	BFF_DUPLICATE_SYMBOLS_REMOVED BinaryFileFlag = 0x0040
	BFF_BYTES_REVERSED_LO         BinaryFileFlag = 0x0080
	BFF_LITTLE_ENDIAN             BinaryFileFlag = 0x0100
	BFF_BIG_ENDIAN_MARKED         BinaryFileFlag = 0x0200
	BFF_REMOVABLE_RUN_FROM_SWAP   BinaryFileFlag = 0x0400
	BFF_NET_RUN_FROM_SWAP         BinaryFileFlag = 0x0800
	BFF_SYSTEM                    BinaryFileFlag = 0x1000
	BFF_DLL                       BinaryFileFlag = 0x2000
	BFF_UP_SYSTEM_ONLY            BinaryFileFlag = 0x4000
	BFF_BYTES_REVERSED_HI         BinaryFileFlag = 0x8000
)

var BinaryFileFlagsMap = map[BinaryFileFlag]bool{
	BFF_STRIPPED_RELOCATION:       true,
	BFF_IS_RELOCATABLE:            true,
	BFF_STRIPPED_LINE_NUMBERS:     true,
	BFF_STRIPPED_SYMBOLS:          true,
	BFF_AGGRESSIVE_WS_TRIM:        true,
	BFF_LARGE_ADDRESS_AWARE:       true,
	BFF_DUPLICATE_SYMBOLS_REMOVED: true,
	BFF_BYTES_REVERSED_LO:         true,
	BFF_LITTLE_ENDIAN:             true,
	BFF_BIG_ENDIAN_MARKED:         true,
	BFF_REMOVABLE_RUN_FROM_SWAP:   true,
	BFF_NET_RUN_FROM_SWAP:         true,
	BFF_SYSTEM:                    true,
	BFF_DLL:                       true,
	BFF_UP_SYSTEM_ONLY:            true,
	BFF_BYTES_REVERSED_HI:         true,
}

type BinaryFileTarget uint16

const (
	BFT_UNKNOWN     BinaryFileTarget = 0x0
	BFT_ALPHA       BinaryFileTarget = 0x184
	BFT_ALPHA64     BinaryFileTarget = 0x284
	BFT_AM33        BinaryFileTarget = 0x1d3
	BFT_AMD64       BinaryFileTarget = 0x8664
	BFT_ARM         BinaryFileTarget = 0x1c0
	BFT_ARM64       BinaryFileTarget = 0xaa64
	BFT_ARMNT       BinaryFileTarget = 0x1c4
	BFT_EBC         BinaryFileTarget = 0xebc
	BFT_I386        BinaryFileTarget = 0x14c
	BFT_IA64        BinaryFileTarget = 0x200
	BFT_LOONGARCH32 BinaryFileTarget = 0x6232
	BFT_LOONGARCH64 BinaryFileTarget = 0x6264
	BFT_M32R        BinaryFileTarget = 0x9041
	BFT_MIPS16      BinaryFileTarget = 0x266
	BFT_MIPSFPU     BinaryFileTarget = 0x366
	BFT_MIPSFPU16   BinaryFileTarget = 0x466
	BFT_POWERPC     BinaryFileTarget = 0x1f0
	BFT_POWERPCFP   BinaryFileTarget = 0x1f1
	BFT_R4000       BinaryFileTarget = 0x166
	BFT_RISCV32     BinaryFileTarget = 0x5032
	BFT_RISCV64     BinaryFileTarget = 0x5064
	BFT_RISCV128    BinaryFileTarget = 0x5128
	BFT_SH3         BinaryFileTarget = 0x1a2
	BFT_SH3DSP      BinaryFileTarget = 0x1a3
	BFT_SH4         BinaryFileTarget = 0x1a6
	BFT_SH5         BinaryFileTarget = 0x1a8
	BFT_THUMB       BinaryFileTarget = 0x1c2
	BFT_WCEMIPSV2   BinaryFileTarget = 0x169
)

var BinaryFileTargetMap = map[BinaryFileTarget]bool{
	BFT_UNKNOWN:     true,
	BFT_ALPHA:       true,
	BFT_ALPHA64:     true,
	BFT_AM33:        true,
	BFT_AMD64:       true,
	BFT_ARM:         true,
	BFT_ARM64:       true,
	BFT_ARMNT:       true,
	BFT_EBC:         true,
	BFT_I386:        true,
	BFT_IA64:        true,
	BFT_LOONGARCH32: true,
	BFT_LOONGARCH64: true,
	BFT_M32R:        true,
	BFT_MIPS16:      true,
	BFT_MIPSFPU:     true,
	BFT_MIPSFPU16:   true,
	BFT_POWERPC:     true,
	BFT_POWERPCFP:   true,
	BFT_R4000:       true,
	BFT_RISCV32:     true,
	BFT_RISCV64:     true,
	BFT_RISCV128:    true,
	BFT_SH3:         true,
	BFT_SH3DSP:      true,
	BFT_SH4:         true,
	BFT_SH5:         true,
	BFT_THUMB:       true,
	BFT_WCEMIPSV2:   true,
}

type BinarySectionFlag int

const (
	BSF_REGULAR BinarySectionFlag = iota
	BSF_DUMMY
	BSF_NO_LOAD
	BSF_GROUP
	BSF_PADDING
	BSF_COPY
	BSF_TEXT
	BSF_DATA
	BSF_BSS
	BSF_BLOCK
	BSF_PASS
	BSF_CLINK
	BSF_VECTOR
	BSF_PADDED
)

var BinarySectionFlagMap = map[BinarySectionFlag]bool{
	BSF_REGULAR: true,
	BSF_DUMMY:   true,
	BSF_NO_LOAD: true,
	BSF_GROUP:   true,
	BSF_PADDING: true,
	BSF_COPY:    true,
	BSF_TEXT:    true,
	BSF_DATA:    true,
	BSF_BSS:     true,
	BSF_BLOCK:   true,
	BSF_PASS:    true,
	BSF_CLINK:   true,
	BSF_VECTOR:  true,
	BSF_PADDED:  true,
}

type RelocationType uint16

const (
	RT_ADD RelocationType = iota
	RT_SUB
	RT_NEG
	RT_MPY
	RT_DIV
	RT_MOD
	RT_SR
	RT_ASR
	RT_SL
	RT_AND
	RT_OR
	RT_XOR
	RT_NOTB
	RT_ULDFLD
	RT_SLDFLD
	RT_USTFLD
	RT_SSTFLD
	RT_PUSH
	RT_PUSHSK
	RT_PUSHUK
	RT_PUSHPC
	RT_DUP
	RT_XSTFLD
	RT_PUSHSV
	RT_ABS
	RT_RELBYTE
	RT_RELWORD
	RT_REL24
	RT_RELLONG
)

var RelocationTypeMap = map[RelocationType]bool{
	RT_ADD:     true,
	RT_SUB:     true,
	RT_NEG:     true,
	RT_MPY:     true,
	RT_DIV:     true,
	RT_MOD:     true,
	RT_SR:      true,
	RT_ASR:     true,
	RT_SL:      true,
	RT_AND:     true,
	RT_OR:      true,
	RT_XOR:     true,
	RT_NOTB:    true,
	RT_ULDFLD:  true,
	RT_SLDFLD:  true,
	RT_USTFLD:  true,
	RT_SSTFLD:  true,
	RT_PUSH:    true,
	RT_PUSHSK:  true,
	RT_PUSHUK:  true,
	RT_PUSHPC:  true,
	RT_DUP:     true,
	RT_XSTFLD:  true,
	RT_PUSHSV:  true,
	RT_ABS:     true,
	RT_RELBYTE: true,
	RT_RELWORD: true,
	RT_REL24:   true,
	RT_RELLONG: true,
}

type StorageClass uint8

const (
	SC_NULL StorageClass = iota
	SC_AUTO
	SC_EXT
	SC_STAT
	SC_REG
	SC_EXTREF
	SC_LABEL
	SC_ULABEL
	SC_MOS
	SC_ARG
	SC_STRTAG
	SC_MOU
	SC_UNTAG
	SC_TPDEF
	SC_USTATIC
	SC_ENTAG
	SC_MOE
	SC_REGPARAM
	SC_FIELD
	SC_UEXT
	SC_STATLAB
	SC_EXTLAB

	SC_VARARG StorageClass = 27

	SC_BLOCK StorageClass = 100
	SC_FCN   StorageClass = 101
	SC_EOS   StorageClass = 102
	SC_FILE  StorageClass = 103
	SC_LINE  StorageClass = 104
)

var StorageClassMap = map[StorageClass]bool{
	SC_NULL:     true,
	SC_AUTO:     true,
	SC_EXT:      true,
	SC_STAT:     true,
	SC_REG:      true,
	SC_EXTREF:   true,
	SC_LABEL:    true,
	SC_ULABEL:   true,
	SC_MOS:      true,
	SC_ARG:      true,
	SC_STRTAG:   true,
	SC_MOU:      true,
	SC_UNTAG:    true,
	SC_TPDEF:    true,
	SC_USTATIC:  true,
	SC_ENTAG:    true,
	SC_MOE:      true,
	SC_REGPARAM: true,
	SC_FIELD:    true,
	SC_UEXT:     true,
	SC_STATLAB:  true,
	SC_EXTLAB:   true,
	SC_VARARG:   true,
	SC_BLOCK:    true,
	SC_FCN:      true,
	SC_EOS:      true,
	SC_FILE:     true,
	SC_LINE:     true,
}
