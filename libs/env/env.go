package env

import (
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func GetPlatformBinExt() string {
	if ostype == "windows" {
		return ".exe"
	}
	return ""
}
