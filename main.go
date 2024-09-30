package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	addr    = "127.0.0.1:12345"
	network = "tcp"
)

var sayings = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	defer listener.Close()

	log.Printf("Server is listening on %s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConn(conn) // Запуск обработки соединения в отдельной горутине
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	for {
		saying := sayings[rand.Intn(len(sayings))]
		_, err := conn.Write([]byte(saying + "\r\n"))
		if err != nil {
			log.Printf("Failed to write to connection: %v", err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}
