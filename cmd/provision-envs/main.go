package main

import (
	"apivalues"
	logs "loggers"
)

func main() {
	result:= apivalues.GetApiValues()
	logs.Info(result)
}
