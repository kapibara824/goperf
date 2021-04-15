package main

import (
	"goperf/Perf"
)

func main() {
	Perf.StartTcpServer()
	Perf.StartTcpClient()

}
