package main

import "runtime"

func detectOs() string {
	var detectedOS string
	if runtime.GOOS == "windows" {
		detectedOS = "windows"
	} else if runtime.GOOS == "linux" {
		detectedOS = "linux"
	}
	return detectedOS
}
