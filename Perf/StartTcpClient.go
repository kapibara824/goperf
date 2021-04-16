package Perf

import (
	"bytes"
	"log"
	"net"
)

func StartTcpClient(ipaddr string) {
	sum := 0
	ec := 0
	count := 0
	data := 0
	repeat := bytes.Repeat([]byte("a"), 2048*10)
	add := ipaddr + ":8240"
	addr, err := net.ResolveTCPAddr("tcp", add)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	for {

		n, err := conn.Write(repeat)
		if err == nil {
			data = int(n)
		}
		count++
		sum += int(n)
		if err != nil {
			ec++
			if ec > 5 {
				log.Println(count, data, count*data, ec)
				log.Println("send", sum, "bytes")
				log.Fatal("Error Count is 5")
			}
		}
		/*
			r, err := conn.Read(res)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(r)
		*/
	}

}
