package langutils

import "testing"

func TestGetFuncInfo(t *testing.T) {
	funcInfo := GetFuncInfo()

	if funcInfo != "github.com/rykugur/goutils/lang.TestGetFuncInfo" {
		t.Error("expected=github.com/rykugur/langutils.TestGetFuncInfo, actual=", funcInfo)
	}
}

func TestGetPackageName(t *testing.T) {
	packageName := GetPackageName()

	if packageName != "github.com/rykugur/goutils/lang" {
		t.Error("expected=github.com/rykugur/langutils, actual=", packageName)
	}
}

func TestGetFuncName(t *testing.T) {
	funcName := GetFuncName()

	if funcName != "TestGetFuncName" {
		t.Error("expected=TestGetFuncName, actual=", funcName)
	}
}
