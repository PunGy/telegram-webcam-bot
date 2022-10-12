package helpers

import (
	"path"
	"runtime"
)

func GetDir(pathOut string) string {
	_, filePath, _, _ := runtime.Caller(0)

	return path.Dir(path.Join(filePath, pathOut))
}
