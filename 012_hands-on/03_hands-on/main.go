package main

import (
	"html/template"
	"log"
	"os"
)

type region int

const (
	Southern region = iota
	Central
	Northern
)

func (r region) String() string {
	names := [...]string{
		"Southern",
		"Central",
		"Northern"}
	return names[r]
}

type hotel struct {
	Name string
	Address,
	City string
	Zip    string
	Region region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "California",
			Address: "Dark Desert Highway",
			City:    "Unknown",
			Zip:     "0010",
			Region:  Southern,
		},
		hotel{
			Name:    "Grand",
			Address: "Great Road",
			City:    "Nowhere",
			Zip:     "0110",
			Region:  Northern,
		},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
