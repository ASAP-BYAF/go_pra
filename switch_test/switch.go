package switch_test

import "runtime"

func SwitchTest() string{
    switch os := runtime.GOOS; os{
	case "darwin":
		return "darwin"
	case "linux":
		return "linux"
	default:
		return string(os)
	}
}