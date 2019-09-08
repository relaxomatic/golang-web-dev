package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(conn)
	}
}

type Request struct {
	Method   string
	URI      string
	Protocol string
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	lineNum := 0
	var req Request
	for scanner.Scan() {
		text := scanner.Text()
		if lineNum == 0 {
			//We have the request line
			requestParts := strings.Split(text, " ")
			req = Request{
				Method:   requestParts[0],
				URI:      requestParts[1],
				Protocol: requestParts[2],
			}
		}
		if text == "" {
			break
		}
		// fmt.Println(text)
		lineNum++
	}

	var body string
	responseCode := 404
	responseMessage := "Not Found"
	if req.URI == "/" && req.Method == "GET" {
		body = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Code Gangsta</title>
			</head>
			<body>
				GET /
			</body>
			</html>
		`
	}
	if req.URI == "/apply" {
		if req.Method == "GET" {
			body = `
				<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Code Gangsta</title>
				</head>
				<body>
					GET /apply
				</body>
				</html>
			`
		} else if req.Method == "Post" {
			body = `
				<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Code Gangsta</title>
				</head>
				<body>
					POST /apply
				</body>
				</html>
			`
		}
	}
	if body != "" {
		responseCode = 200
		responseMessage = "OK"
	} else {
		body = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Code Gangsta</title>
			</head>
			<body>
				Not Found
			</body>
			</html>
		`
	}
	responseLine := "HTTP/1.1 "
	responseLine += string(responseCode)
	responseLine += " "
	responseLine += responseMessage
	responseLine += "\r\n"
	fmt.Println("Method:", req.Method)
	fmt.Println("URI:", req.URI)
	io.WriteString(conn, responseLine)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	fmt.Println("Code got here.")
}
