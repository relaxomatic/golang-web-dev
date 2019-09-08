package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("body.gohtml"))
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request := request(conn)

	// write response
	respond(conn, request)
}

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	body    string
}

func request(conn net.Conn) *Request {
	i := 0
	request := new(Request)
	request.Headers = map[string]string{}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			request.Method = strings.Fields(ln)[0]
			request.URL = strings.Fields(ln)[1]
		}
		if ln == "" {
			break
		} else {
			indexOfColon := strings.IndexByte(ln, ':')
			if indexOfColon > 0 {
				headerFieldType := ln[:indexOfColon]
				headerFieldValue := strings.Trim(ln[indexOfColon:], " ")
				request.Headers[headerFieldType] = headerFieldValue
			}
		}
		i++
	}
	return request
}

func respond(conn net.Conn, r *Request) {
	bodyBuffer := &bytes.Buffer{}
	var responseCode string
	var responseMessage string
	switch r.URL {
	case "/":
		responseCode = "200"
		responseMessage = "OK"
		break
	default:
		responseCode = "404"
		responseMessage = "Not Found"
		break
	}
	err := tpl.Execute(bodyBuffer, r)
	if err != nil {
		panic(err)
	}
	body := bodyBuffer.String()

	fmt.Fprint(conn, strings.Join([]string{"HTTP/1.1", responseCode, responseMessage, "\r\n"}, " "))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
