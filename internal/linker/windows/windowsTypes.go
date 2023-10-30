package linker

type WindowsBinaryFileFlag uint16

const (
	WIN_BFF_STRIPPED_RELOCATION   WindowsBinaryFileFlag = 0x0001
	WIN_BFF_IS_RELOCATABLE        WindowsBinaryFileFlag = 0x0002
	WIN_BFF_STRIPPED_LINE_NUMBERS WindowsBinaryFileFlag = 0x0004
	WIN_BFF_STRIPPED_SYMBOLS      WindowsBinaryFileFlag = 0x0008
	WIN_BFF_AGGRESSIVE_WS_TRIM    WindowsBinaryFileFlag = 0x0010
	WIN_BFF_LARGE_ADDRESS_AWARE   WindowsBinaryFileFlag = 0x0020
	// WIN_BFF_DUPLICATE_SYMBOLS_REMOVED WindowsBinaryFileFlag = 0x0040
	WIN_BFF_BYTES_REVERSED_LO       WindowsBinaryFileFlag = 0x0080
	WIN_BFF_32BIT_MACHINE           WindowsBinaryFileFlag = 0x0100
	WIN_BFF_DEBUG_STRIPPED          WindowsBinaryFileFlag = 0x0200
	WIN_BFF_REMOVABLE_RUN_FROM_SWAP WindowsBinaryFileFlag = 0x0400
	WIN_BFF_NET_RUN_FROM_SWAP       WindowsBinaryFileFlag = 0x0800
	WIN_BFF_SYSTEM                  WindowsBinaryFileFlag = 0x1000
	WIN_BFF_DLL                     WindowsBinaryFileFlag = 0x2000
	WIN_BFF_UP_SYSTEM_ONLY          WindowsBinaryFileFlag = 0x4000
	WIN_BFF_BYTES_REVERSED_HI       WindowsBinaryFileFlag = 0x8000
)

var windowsBinaryFileFlagMap = map[WindowsBinaryFileFlag]bool{
	WIN_BFF_STRIPPED_RELOCATION:   true,
	WIN_BFF_IS_RELOCATABLE:        true,
	WIN_BFF_STRIPPED_LINE_NUMBERS: true,
	WIN_BFF_STRIPPED_SYMBOLS:      true,
	WIN_BFF_AGGRESSIVE_WS_TRIM:    true,
	WIN_BFF_LARGE_ADDRESS_AWARE:   true,
	// WIN_BFF_DUPLICATE_SYMBOLS_REMOVED: true,
	WIN_BFF_BYTES_REVERSED_LO:       true,
	WIN_BFF_32BIT_MACHINE:           true,
	WIN_BFF_DEBUG_STRIPPED:          true,
	WIN_BFF_REMOVABLE_RUN_FROM_SWAP: true,
	WIN_BFF_NET_RUN_FROM_SWAP:       true,
	WIN_BFF_SYSTEM:                  true,
	WIN_BFF_DLL:                     true,
	WIN_BFF_UP_SYSTEM_ONLY:          true,
	WIN_BFF_BYTES_REVERSED_HI:       true,
}

type WindowsSubsystem uint16

const (
	WS_UNKNOWN WindowsSubsystem = iota
	WS_NATIVE
	WS_WINDOWS_GUI
	WS_WINDOWS_CUI
	WS_OS2_CUI
	WS_POSIX_CUI
	WS_NATIVE_WINDOWS
	WS_WINDOWS_CE_GUI
	WS_EFI_APPLICATION
	WS_EFI_BOOT_SERVICE_DRIVER
	WS_EFI_RUNTIME_DRIVER
	WS_EFI_ROM
	WS_XBOX
	WS_WINDOWS_BOOT_APPLICATION = 16
)

var windowsSubsystemMap = map[WindowsSubsystem]bool{
	WS_UNKNOWN:                  true,
	WS_NATIVE:                   true,
	WS_WINDOWS_GUI:              true,
	WS_WINDOWS_CUI:              true,
	WS_OS2_CUI:                  true,
	WS_POSIX_CUI:                true,
	WS_NATIVE_WINDOWS:           true,
	WS_WINDOWS_CE_GUI:           true,
	WS_EFI_APPLICATION:          true,
	WS_EFI_BOOT_SERVICE_DRIVER:  true,
	WS_EFI_RUNTIME_DRIVER:       true,
	WS_EFI_ROM:                  true,
	WS_XBOX:                     true,
	WS_WINDOWS_BOOT_APPLICATION: true,
}

type DllCharacteristic uint16

const (
	DC_HIGH_ENTROPY_VA       DllCharacteristic = 0x0020
	DC_DYNAMIC_BASE          DllCharacteristic = 0x0040
	DC_FORCE_INTEGRITY       DllCharacteristic = 0x0080
	DC_NX_COMPAT             DllCharacteristic = 0x0100
	DC_NO_ISOLATION          DllCharacteristic = 0x0200
	DC_NO_SEH                DllCharacteristic = 0x0400
	DC_NO_BIND               DllCharacteristic = 0x0800
	DC_APP_CONTAINER         DllCharacteristic = 0x1000
	DC_WDM_DRIVER            DllCharacteristic = 0x2000
	DC_GUARD_CF              DllCharacteristic = 0x4000
	DC_TERMINAL_SERVER_AWARE DllCharacteristic = 0x8000
)

var dllCharacteristicMap = map[DllCharacteristic]bool{
	DC_HIGH_ENTROPY_VA:       true,
	DC_DYNAMIC_BASE:          true,
	DC_FORCE_INTEGRITY:       true,
	DC_NX_COMPAT:             true,
	DC_NO_ISOLATION:          true,
	DC_NO_SEH:                true,
	DC_NO_BIND:               true,
	DC_APP_CONTAINER:         true,
	DC_WDM_DRIVER:            true,
	DC_GUARD_CF:              true,
	DC_TERMINAL_SERVER_AWARE: true,
}
