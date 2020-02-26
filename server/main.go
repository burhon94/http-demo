package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
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
		err = handleConn(conn)
		if err != nil {
			log.Printf("can't handle connect: %v", err)
		}
	}
}

func handleConn(conn net.Conn) (err error) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("can't close handle connect: %v", err)
		}
	}()
	log.Println("success handle")
	log.Println("try read client request")
	reader := bufio.NewReaderSize(conn, 4096)
	writer := bufio.NewWriter(conn)
	counter := 0
	buffer := [4096]byte{}

	for {
		if counter == 4096 {
			log.Printf("request too long header")
			log.Printf("too long request header")
			_, _ = writer.WriteString("HTTP/1.1 413 Payload Too Large\r\n")
			_, _ = writer.WriteString("Content-Length: 0\r\n")
			_, _ = writer.WriteString("Connection: close\r\n")
			_, _ = writer.WriteString("\r\n")
			err := writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
			return err
		}

		read, err := reader.ReadByte()
		if err != nil {
			log.Printf("can't read request line: %v", err)
			_, _ = writer.WriteString("HTTP/1.1 400 Bad Request\r\n")
			_, _ = writer.WriteString("Content-Length: 0\r\n")
			_, _ = writer.WriteString("Connection: close\r\n")
			_, _ = writer.WriteString("\r\n")
			err := writer.Flush()
			if err != nil {
				log.Printf("can't sent response: %v", err)
			}
			return err
		}

		buffer[counter] = read
		counter++
		if counter < 4 {
			continue
		}
		if string(buffer[counter-4:counter]) == "\r\n\r\n" {
			break
		}
	}

	headerString := string(buffer[:counter-4])
	requestHeaderParts := strings.Split(headerString, "\r\n")
	log.Println("parse request line")
	requestLine := requestHeaderParts[0]
	requestParts := strings.Split(strings.TrimSpace(requestLine), " ")

	if len(requestParts) != 3 {
		return err
	}

	method, request, protocol := requestParts[0], requestParts[1], requestParts[2]
	typeOfContent := ""
	nameOfFile := ""

	log.Printf("request: %s", request)

	if method == "GET" && protocol == "HTTP/1.1" {

		switch request {
		// HTML Requests
		case "/":
			{
				file, err := os.Open("./server/pages/index.html")
				if err != nil {
					log.Printf("can't open index.html: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "text/html"
			}

		//Image Requests
		case "/favicon.ico":
			{
				file, err := os.Open("./server/pages/img/icon.png")
				if err != nil {
					log.Printf("can't open icon.png: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "image/x-icon"
			}
		case "/img/img1.jpg":
			{
				file, err := os.Open("./server/pages/img/img1.jpg")
				if err != nil {
					log.Printf("can't open img1.jpg: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "image/jpg"
			}
			case "/img/img2.jpg":
			{
				file, err := os.Open("./server/pages/img/img2.jpg")
				if err != nil {
					log.Printf("can't open img2.jpg: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "image/jpg"
			}
			case "/img/2.png":
			{
				file, err := os.Open("./server/pages/img/2.png")
				if err != nil {
					log.Printf("can't open 2.jpg: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "image/png"
			}
			case "/img/bg3-dots.png":
			{
				file, err := os.Open("./server/pages/img/bg3-dots.png")
				if err != nil {
					log.Printf("can't open bg3-dots.png: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "image/png"
			}
			
		//CSS Requests 	
		case "/css/styles.css":
			{
				file, err := os.Open("./server/pages/css/styles.css")
				if err != nil {
					log.Printf("can't open styles.css: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "text/css"
			}
		
		//JS Requests
		case "/js/script.js":
			{
				file, err := os.Open("./server/pages/js/script.js")
				if err != nil {
					log.Printf("can't open script.js: %v", err)
					return err
				}
				nameOfFile += file.Name()
				typeOfContent += "text/javascript"
			}
		
		}
	} else {
		log.Printf("Wrong Method: %s, or Protocol: %s", method, protocol)
		return err
	}

	err = writeHeader(conn, nameOfFile, typeOfContent, request)
	if err != nil {
		log.Printf("can't response to reqeest: %s, error: %v", request, err)
		return err
	}
	return nil
}

func writeHeader(conn net.Conn, fileName, contentType, request string) (err error) {
	writer := bufio.NewWriter(conn)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("can't read file: %v", err)
		return err
	}
	_, err = writer.WriteString("HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
		return err
	}
	_, err = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
	if err != nil {
		log.Printf("can't write: %v", err)
	}
	_, err = writer.WriteString("Content-Type:" + " " + contentType + "\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
		return err
	}
	_, err = writer.WriteString("Connection: Close\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
		return err
	}
	_, err = writer.WriteString("\r\n")
	if err != nil {
		log.Printf("can't write: %v", err)
		return err
	}
	_, err = writer.Write(bytes)
	if err != nil {
		log.Printf("can't write: %v", err)
		return err
	}
	err = writer.Flush()
	if err != nil {
		log.Printf("can't sent response: %v", err)
		return err
	}
	log.Printf("response on: %s", request)
	return nil
}
