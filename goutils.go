package goutils

import "runtime"
import "strings"

// Private worker method.
func _getFuncInfo(skip int) string {
	pc := make([]uintptr, 10)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])

	return strings.Trim(f.Name(), " ")
}

// Gets the FuncInfo for the current function.
func GetFuncInfo() string {
	return _getFuncInfo(3)
}

// Gets the Package name the current function lives in.
func GetPackageName() string {
	funcInfo := _getFuncInfo(3)

	dotIndex := strings.LastIndex(funcInfo, ".")

	return funcInfo[:dotIndex]
}

// Gets the name of the current function.
func GetFuncName() string {
	funcInfo := _getFuncInfo(3)

	dotIndex := strings.LastIndex(funcInfo, ".")

	return funcInfo[dotIndex+1:] // ignore the dot
}
