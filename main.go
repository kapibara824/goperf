package main

import (
	"flag"
	"goperf/Perf"
	"log"
)

var (
	servermode = flag.Bool("s", false, "Server mode")
	ipaddr     = flag.String("c", "localhost", "Destination Address")
)

func main() {
	flag.Parse()
	if *servermode {
		log.Println("Running Server mode.")
		Perf.StartTcpServer()
	}
	if !*servermode {
		log.Println("Running Client mode.")
		log.Println("Send to", *ipaddr)
		Perf.StartTcpClient(*ipaddr)
	}

}
