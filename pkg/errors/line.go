package errs

import (
	"fmt"
	"runtime"
)

func line() string {
	_, file, line, _ := runtime.Caller(2)

	callerLine := splitLine(file, 3)

	return fmt.Sprintf("[%s:%d]", callerLine, line)
}

func splitLine(file string, piece int) string {
	short := file
	var cnt int
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			cnt++
			if cnt == piece {
				short = file[i:]
				break
			}

		}
	}
	file = short
	return file
}
