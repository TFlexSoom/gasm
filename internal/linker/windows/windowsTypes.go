package linker

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
