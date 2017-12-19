package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/nuuls/speedtest/util"
)

func main() {
	var (
		tcpAddr = flag.String("addr", getenv("ADDR", ":3254"), "address to listen on")
	)
	flag.Parse()
	ln, err := net.Listen("tcp", *tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("cannot accept conn:", err)
			continue
		}
		go handleConn(conn)
	}

}

const (
	testLength = time.Second * 60
	bufLength  = 1024 * 512 // 512 KB
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Println("new client connected:", conn.RemoteAddr())
	buf := make([]byte, bufLength) // 256 KB
	for i := range buf {
		buf[i] = byte(i) // fill with random shit
	}
	start := time.Now()
	bytesWritten := 0
	for {
		now := time.Now()
		if now.Sub(start) > testLength {
			log.Printf("test completed for %s: %d bytes in %s (%s/s)",
				conn.RemoteAddr(), bytesWritten, testLength,
				util.PrettyPrintBytes(util.PerSecond(bytesWritten, testLength)))
			return
		}
		conn.SetWriteDeadline(now.Add(testLength))
		n, err := conn.Write(buf)
		if err != nil {
			log.Println("client died:", err)
			return
		}
		bytesWritten += n
	}
}

func getenv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
