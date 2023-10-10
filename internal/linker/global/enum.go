package linker_global

type BinaryFileFlag int

const (
	BFF_STRIPPED_RELOCATION BinaryFileFlag = iota
	BFF_IS_RELOCATABLE
	BFF_STRIPPED_LINE_NUMBERS
	BFF_STRIPPED_SYMBOLS
	BFF_LITTLE_ENDIAN
	BFF_BIG_ENDIAN_MARKED
	BFF_DUPLICATE_SYMBOLS_REMOVED
	BFF_LARGE_ADDRESS_AWARE
	BFF_32BIT_MACHINE
	BFF_SYSTEM
	BFF_DLL
	BFF_UP_SYSTEM_ONLY
	BFF_BYTES_REVERSED_HI
	BFF_BYTES_REVERSED_LO
)

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
	BFT_AXP64       BinaryFileTarget = 0x284
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

type StorageClass uint8

const (
	SC_NULL uint8 = iota
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

	SC_VARARG = 27

	SC_BLOCK = 100
	SC_FCN   = 101
	SC_EOS   = 102
	SC_FILE  = 103
	SC_LINE  = 104
)
