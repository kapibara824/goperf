package Perf

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var sum int64 = 0
var tsum int64 = 0

func StartTcpServer() {
	CloseHandler()

	addr, _ := net.ResolveTCPAddr("tcp", ":8240")
	conn, _ := net.ListenTCP("tcp", addr)

	fmt.Println("Listening...")
	defer conn.Close()

	n, _ := conn.AcceptTCP()
	defer n.Close()

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("%.3fGbps\n", float64(sum-tsum)/3/1024/1024/1024*8)
			tsum = sum
		default:
			n, _ := io.CopyN(io.Discard, n, 1024*1024)
			sum += n
		}
	}
}

func CloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("interrupted")
		log.Println("received", sum, "bytes")
		os.Exit(0)
	}()
}
