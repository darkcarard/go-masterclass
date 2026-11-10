package main

import "fmt"

type LogLevel int

const (
	LogWarn LogLevel = iota
	LogInfo
	LogDebug
	LogError
	LogFatal
)

func main() {
	fmt.Println("Hello world!")
}
