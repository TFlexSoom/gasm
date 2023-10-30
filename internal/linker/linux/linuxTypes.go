package linker

const EI_MAG = string(rune(0x7f)) + "ELF"

type EI_CLASS rune

const (
	EI_CLASS_ELFCLASSNONE EI_CLASS = iota
	EI_CLASS_ELFCLASS32
	EI_CLASS_ELFCLASS64
)

var eiClassMap = map[EI_CLASS]bool{
	EI_CLASS_ELFCLASSNONE: true,
	EI_CLASS_ELFCLASS32:   true,
	EI_CLASS_ELFCLASS64:   true,
}

type EI_DATA rune

const (
	EI_DATA_ELFDATANONE EI_DATA = iota
	EI_DATA_ELFDATALSB
	EI_DATA_ELFDATAMSB
)

var eiDataMap = map[EI_DATA]bool{
	EI_DATA_ELFDATANONE: true,
	EI_DATA_ELFDATALSB:  true,
	EI_DATA_ELFDATAMSB:  true,
}

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

var eiVersionMap = map[EI_VERSION]bool{
	EI_VERSION_ELFOSABI_NONE:       true,
	EI_VERSION_ELFOSABI_SYSV:       true,
	EI_VERSION_ELFOSABI_HPUX:       true,
	EI_VERSION_ELFOSABI_NETBSD:     true,
	EI_VERSION_ELFOSABI_LINUX:      true,
	EI_VERSION_ELFOSABI_SOLARIS:    true,
	EI_VERSION_ELFOSABI_IRIX:       true,
	EI_VERSION_ELFOSABI_FREEBSD:    true,
	EI_VERSION_ELFOSABI_TRU64:      true,
	EI_VERSION_ELFOSABI_ARM:        true,
	EI_VERSION_ELFOSABI_STANDALONE: true,
}

type EI_ABIVERSION rune

const (
	EI_ABIVERSION_DEFAULT EI_ABIVERSION = iota
)

var eiAbiVersionMap = map[EI_ABIVERSION]bool{
	EI_ABIVERSION_DEFAULT: true,
}

type E_TYPE uint16

const (
	E_TYPE_ET_NONE E_TYPE = iota
	E_TYPE_ET_REL         // relocatable
	E_TYPE_ET_EXEC        // executable
	E_TYPE_ET_DYN         // shared object
	E_TYPE_ET_CORE        // core file
)

var eTypeMap = map[E_TYPE]bool{
	E_TYPE_ET_NONE: true,
	E_TYPE_ET_REL:  true,
	E_TYPE_ET_EXEC: true,
	E_TYPE_ET_DYN:  true,
	E_TYPE_ET_CORE: true,
}

type E_MACHINE uint16

const (
	E_MACHINE_EM_NONE        E_MACHINE = iota
	E_MACHINE_EM_M32                   /** AT&T WE 32100 **/
	E_MACHINE_EM_SPARC                 /** Sun Microsystems SPARC **/
	E_MACHINE_EM_386                   /** Intel 80386 **/
	E_MACHINE_EM_68K                   /** Motorola 68000 **/
	E_MACHINE_EM_88K                   /** Motorola 88000 **/
	E_MACHINE_EM_860                   /** Intel 80860 **/
	E_MACHINE_EM_MIPS                  /** MIPS RS3000 (big-endian only) **/
	E_MACHINE_EM_PARISC                /** HP/PA **/
	E_MACHINE_EM_SPARC32PLUS           /** SPARC with enhanced instruction set **/
	E_MACHINE_EM_PPC                   /** PowerPC **/
	E_MACHINE_EM_PPC64                 /** PowerPC 64-bit **/
	E_MACHINE_EM_S390                  /** IBM S/390 **/
	E_MACHINE_EM_ARM                   /** Advanced RISC Machines **/
	E_MACHINE_EM_SH                    /** Renesas SuperH **/
	E_MACHINE_EM_SPARCV9               /** SPARC v9 64-bit **/
	E_MACHINE_EM_IA_64                 /** Intel Itanium **/
	E_MACHINE_EM_X86_64                /** AMD x86-64 **/
	E_MACHINE_EM_VAX                   /** DEC Vax **/
)

var eMachineMap = map[E_MACHINE]bool{
	E_MACHINE_EM_NONE:        true,
	E_MACHINE_EM_M32:         true,
	E_MACHINE_EM_SPARC:       true,
	E_MACHINE_EM_386:         true,
	E_MACHINE_EM_68K:         true,
	E_MACHINE_EM_88K:         true,
	E_MACHINE_EM_860:         true,
	E_MACHINE_EM_MIPS:        true,
	E_MACHINE_EM_PARISC:      true,
	E_MACHINE_EM_SPARC32PLUS: true,
	E_MACHINE_EM_PPC:         true,
	E_MACHINE_EM_PPC64:       true,
	E_MACHINE_EM_S390:        true,
	E_MACHINE_EM_ARM:         true,
	E_MACHINE_EM_SH:          true,
	E_MACHINE_EM_SPARCV9:     true,
	E_MACHINE_EM_IA_64:       true,
	E_MACHINE_EM_X86_64:      true,
	E_MACHINE_EM_VAX:         true,
}

type E_VERSION uint32

const (
	E_VERSION_EV_NONE E_VERSION = iota
	E_VERSION_EV_CURRENT
)

var eVersionMap = map[E_VERSION]bool{
	E_VERSION_EV_NONE:    true,
	E_VERSION_EV_CURRENT: true,
}

// Maybe Programm Header FLags
// D Flags?
