package main

import (
	"flag"
	"goperf/Perf"
	"log"
)

var (
	servermode = flag.Bool("s", false, "Server mode")
	ipaddr     = flag.String("c", "localhost", "Destination Address")
	time       = flag.Int("t", 10, "time")
)

func main() {
	flag.Parse()
	if *servermode {
		log.Println("Running Server mode.")
		Perf.StartTcpServer(*time)
	}
	if !*servermode {
		log.Println("Running Client mode.")
		log.Println("Send to", *ipaddr)
		Perf.StartTcpClient(*ipaddr)
	}

}
