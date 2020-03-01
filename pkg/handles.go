package pkg

import (
	"bufio"
	"log"
	"net"
	"path"
	"strings"
)

func HandleConn(conn net.Conn) (err error) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("can't close handle connect: %v", err)
			return
		}
	}()
	log.Print("read client request")
	reader := bufio.NewReaderSize(conn, 4096)
	counter := 0
	buf := [4096]byte{}

	for {
		if counter == 4096 {
			log.Printf("too long request header")
			err := HTTP413(conn)
			if err != nil {
				log.Printf("can't sent response: %v ", err)
			}
			return err
		}
		read, err := reader.ReadByte()
		if err != nil {
			log.Printf("can't read request line: %v", err)
			err := HTTP400(conn)
			if err != nil {
				log.Printf("can't sent response: %v ", err)
			}
			return err
		}
		buf[counter] = read
		counter++
		if counter < 4 {
			continue
		}
		if string(buf[counter-4:counter]) == "\r\n\r\n" {
			break
		}
	}
	headersStr := string(buf[:counter-4])
	requestHeaderParts := strings.Split(headersStr, "\r\n")
	log.Print("parse request line")
	requestLine := requestHeaderParts[0]
	requestParts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(requestParts) != 3 {
		return
	}

	method, request, protocol := requestParts[0], requestParts[1], requestParts[2]
	for {

		if method != "GET" {
			return
		}
		if protocol != "HTTP/1.1" {
			return
		}

		if request == "/" {
			err := SendFile(conn, "index.html", request)
			if err != nil {
				log.Printf("can't process the request: %v", err)
			}
			return err
		}

		filesDir, err := FilesDir(ServerFilesPages)
		if err != nil {
			log.Printf("can't check server files: %s, error %v", ServerFilesPages, err)
		}

		ext := path.Ext(request)
		if ext == ".html" {
			for _, serverFile := range filesDir {
				if request == serverFile {
					err := SendFile(conn, request, request)
					if err != nil {
						log.Printf("can't process the request: %v", err)
					}
					return err
				}
			}
		}

		for _, serverFile := range filesDir {
			serverFile = serverFile[12:]
			if request == serverFile {
				err := SendFile(conn, request, request)
				if err != nil {
					log.Printf("can't process the request: %v", err)
				}
				return nil
			}
		}

		err = SendFile(conn, "html404.html", request)
		if err != nil {
			log.Printf("can't process the request: %v", err)
		}

		return nil
	}
}
