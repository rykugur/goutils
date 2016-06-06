package goutils

import "testing"

func TestGetFuncInfo(t *testing.T) {
	funcInfo := GetFuncInfo()

	if funcInfo != "github.com/rykugur/utils/goutils.TestGetFuncInfo" {
		t.Error("expected=github.com/rykugur/goutils.TestGetFuncInfo, actual=", funcInfo)
	}
}

func TestGetPackageName(t *testing.T) {
	packageName := GetPackageName()

	if packageName != "github.com/rykugur/utils/goutils" {
		t.Error("expected=github.com/rykugur/goutils, actual=", packageName)
	}
}

func TestGetFuncName(t *testing.T) {
	funcName := GetFuncName()

	if funcName != "TestGetFuncName" {
		t.Error("expected=TestGetFuncName, actual=", funcName)
	}
}
