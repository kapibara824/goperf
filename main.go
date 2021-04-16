package main

import (
	"flag"
	"goperf/Perf"
	"log"
)

var (
	servermode = flag.Bool("s", false, "Server mode")
	ipaddr     = flag.String("c", "localhost", "Destination Address")
	duration   = flag.String("t", "10", "time")
)

func main() {
	flag.Parse()
	if *servermode {
		log.Println("Running Server mode.")
		Perf.StartTcpServer(*duration)
	}
	if !*servermode {
		log.Println("Running Client mode.")
		log.Println("Send to", *ipaddr)
		Perf.StartTcpClient(*ipaddr)
	}

}
