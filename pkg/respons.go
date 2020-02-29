package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path"
	"path/filepath"
)

func HTTP413(conn net.Conn) error {
	writer := bufio.NewWriter(conn)
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

func HTTP400(conn net.Conn) error {
	writer := bufio.NewWriter(conn)
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

func SendFile(conn net.Conn, fileName, request string) (err error) {
	contentType := ""
	fileFormat := filepath.Ext(fileName)

	if fileFormat == ".html" {
		contentType = ContentHTML
		err = headerWriter(conn, ServerFilesPages + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".jpg" {
		contentType = ContentJPG
		fileName = path.Base(fileName)
		err = headerWriter(conn, ServerFilesIMGs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".png" {
		contentType = ContentJPG
		fileName = path.Base(fileName)
		err = headerWriter(conn, ServerFilesIMGs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".gif" {
		contentType = ContentJPG
		fileName = path.Base(fileName)
		err = headerWriter(conn, ServerFilesIMGs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".ico" {
		contentType = ContentICO
		err = headerWriter(conn, ServerFilesIMGs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".css" {
		contentType = ContentCSS
		fileName = path.Base(fileName)
		err = headerWriter(conn, ServerFilesCSSs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	if fileFormat == ".js" {
		contentType = ContentJS
		fileName = path.Base(fileName)
		err = headerWriter(conn, ServerFilesJSs + fileName, contentType, request)
		if err != nil {
			log.Printf("can't sent data to header writer: %v", err)
		}
		return nil
	}

	return nil
}

func headerWriter(conn net.Conn, fileName, contentType, request string) (err error) {
	byteBuff := bytes.Buffer{}
	writer := bufio.NewWriter(conn)
	bytesFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("can't open file: %s, error: %v", fileName, err)
		return err
	}

	log.Printf("request: %s", request)
	_, err = byteBuff.WriteString("HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytesFile)))
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.WriteString("Content-Type: " + contentType + "\r\n")
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.WriteString("Connection: Close\r\n")
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.WriteString("\r\n")
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.Write(bytesFile)
	if err != nil {
		log.Printf("can't write to buffer: %v", err)
	}

	_, err = byteBuff.WriteTo(writer)
	if err != nil {
		log.Printf("error write from buffer to writer: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("error to response: %v", err)
	}

	log.Printf("response on: %s", request)
	return nil
}
