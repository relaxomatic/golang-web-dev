package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type parsedLine []string

type Table struct {
	Header parsedLine
	Rows   []parsedLine
}

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	csvFile, _ := os.Open("table.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var rows []parsedLine
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		rows = append(rows, line)
	}

	err := tpl.Execute(res, Table{
		Header: rows[0],
		Rows:   rows[1:],
	})
	if err != nil {
		log.Fatalln(err)
	}
}
