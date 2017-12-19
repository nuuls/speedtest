package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/nuuls/speedtest/util"
)

func main() {
	addr := "localhost:3254"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	n, err := drain(conn)
	if err != nil {
		log.Println("error reading from connection:", err)
	}
	timePassed := time.Since(start)
	log.Printf("read %d bytes in %s ( %s/s )",
		n, timePassed,
		util.PrettyPrintBytes(util.PerSecond(int(n), timePassed)))
}

func drain(reader io.Reader) (int, error) {
	buf := make([]byte, 1000*1000*10)
	bytesRead := 0
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			return bytesRead, nil
		}
		bytesRead += n
	}
}
