package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/agalue/onms-tcp-receiver/protobuf/perf"
	"github.com/golang/protobuf/proto"
)

func main() {
	var tcpPort int
	log.SetOutput(os.Stdout)
	flag.IntVar(&tcpPort, "port", 8999, "Port to listen for Performance Metrics")
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		l.Close()
		log.Println("Good bye!")
		os.Exit(0)
	}()

	log.Printf("Listening on port %d", tcpPort)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting: %v", err)
			os.Exit(1)
		}
		if payload, err := ioutil.ReadAll(conn); err == nil {
			log.Printf("Received payload of %d bytes from %s", len(payload), conn.RemoteAddr().String())
			perfdata := new(perf.PerformanceDataReadings)
			if err := proto.Unmarshal(payload, perfdata); err == nil {
				log.Printf("Parsed %d messsages", len(perfdata.Message))
				for _, msg := range perfdata.Message {
					bytes, _ := json.Marshal(msg)
					log.Printf("Message: %s", string(bytes))
				}
			} else {
				log.Printf("Error parsing: %v. Payload: %s", err, string(payload))
			}
		} else {
			log.Printf("Error reading: %v", err)
		}
		conn.Close()
	}
}
