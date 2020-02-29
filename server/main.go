package main

import (
	"fmt"
	"github.com/burhon94/http-demo/pkg"
	"log"
	"net"
	"os"
)

func main() {
	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("start server on: %s", addr)

	err := startServer(addr)
	if err != nil {
		log.Fatalf("can't start server on: %s, error: %v", addr, err)
	}
}

func startServer(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("can't listen on: %s, error: %v", addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("can't close server listener: %v", err)
		}
	}()

	for {
		log.Println("try accept connection")
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can't accept connection: %v", err)
			continue
		}

		log.Printf("connect accept success")
		err = pkg.HandleConn(conn)
		if err != nil {
			log.Printf("can't handle connect: %v", err)
		}
	}
}
