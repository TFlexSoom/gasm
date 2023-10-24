package linker

const EI_MAG = string(rune(0x7f)) + "ELF"

type EI_CLASS rune

const (
	EI_CLASS_ELFCLASSNONE EI_CLASS = iota
	EI_CLASS_ELFCLASS32
	EI_CLASS_ELFCLASS64
)

type EI_DATA rune

const (
	EI_DATA_ELFDATANONE EI_DATA = iota
	EI_DATA_ELFDATALSB
	EI_DATA_ELFDATAMSB
)

type EI_VERSION rune

const (
	EI_VERSION_ELFOSABI_NONE       EI_VERSION = iota
	EI_VERSION_ELFOSABI_SYSV                  /* UNIX System V ABI */
	EI_VERSION_ELFOSABI_HPUX                  /* HP-UX ABI */
	EI_VERSION_ELFOSABI_NETBSD                /* NetBSD ABI */
	EI_VERSION_ELFOSABI_LINUX                 /* Linux ABI */
	EI_VERSION_ELFOSABI_SOLARIS               /* Solaris ABI */
	EI_VERSION_ELFOSABI_IRIX                  /* IRIX ABI */
	EI_VERSION_ELFOSABI_FREEBSD               /* FreeBSD ABI */
	EI_VERSION_ELFOSABI_TRU64                 /* TRU64 UNIX ABI */
	EI_VERSION_ELFOSABI_ARM                   /* ARM architecture ABI */
	EI_VERSION_ELFOSABI_STANDALONE            /* Stand-alone (embedded) ABI */
)

type EI_ABIVERSION rune

const (
	EI_ABIVERSION_DEFAULT EI_ABIVERSION = iota
)

type E_TYPE uint16

const (
	E_TYPE_ET_NONE E_TYPE = iota
	E_TYPE_ET_REL         // relocatable
	E_TYPE_ET_EXEC        // executable
	E_TYPE_ET_DYN         // shared object
	E_TYPE_ET_CORE        // core file
)

type E_MACHINE uint16

const (
	E_MACHINE_EM_NONE        uint16 = iota
	E_MACHINE_EM_M32                /** AT&T WE 32100 **/
	E_MACHINE_EM_SPARC              /** Sun Microsystems SPARC **/
	E_MACHINE_EM_386                /** Intel 80386 **/
	E_MACHINE_EM_68K                /** Motorola 68000 **/
	E_MACHINE_EM_88K                /** Motorola 88000 **/
	E_MACHINE_EM_860                /** Intel 80860 **/
	E_MACHINE_EM_MIPS               /** MIPS RS3000 (big-endian only) **/
	E_MACHINE_EM_PARISC             /** HP/PA **/
	E_MACHINE_EM_SPARC32PLUS        /** SPARC with enhanced instruction set **/
	E_MACHINE_EM_PPC                /** PowerPC **/
	E_MACHINE_EM_PPC64              /** PowerPC 64-bit **/
	E_MACHINE_EM_S390               /** IBM S/390 **/
	E_MACHINE_EM_ARM                /** Advanced RISC Machines **/
	E_MACHINE_EM_SH                 /** Renesas SuperH **/
	E_MACHINE_EM_SPARCV9            /** SPARC v9 64-bit **/
	E_MACHINE_EM_IA_64              /** Intel Itanium **/
	E_MACHINE_EM_X86_64             /** AMD x86-64 **/
	E_MACHINE_EM_VAX                /** DEC Vax **/
)

type E_VERSION uint32

const (
	E_VERSION_EV_NONE uint32 = iota
	E_VERSION_EV_CURRENT
)

// Maybe Programm Header FLags
// D Flags?
