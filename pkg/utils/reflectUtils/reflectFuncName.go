package reflectUtils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameEnd := filepath.Ext(runtime.FuncForPC(pc).Name()) // errrorCode/errorCode.E1000100
	name := strings.TrimPrefix(nameEnd, ".")
	return name
}
