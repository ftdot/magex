package platform

import "runtime"

type PlatformSpec uint8
type PlatformGenericSpec uint8

const (
	PlatformWindows PlatformSpec = iota
	PlatformDarwin
	PlatformAndroid
	PlatformIOS
	PlatformOther
)

var platform PlatformSpec = GetPlatformRuntime()

func GetPlatformRuntime() PlatformSpec {
	switch runtime.GOOS {
	case "windows": return PlatformWindows
	case "darwin": return PlatformDarwin
	case "android": return PlatformAndroid
	case "ios": return PlatformIOS
	default: return PlatformOther
	}
}

func Platform() PlatformSpec {
	return platform
}

func IsMobile() bool {
	return platform == PlatformAndroid || platform == PlatformIOS
}

func IsDesktop() bool {
	return !IsMobile()
}
