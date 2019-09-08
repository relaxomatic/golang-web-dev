package main

import (
	"html/template"
	"log"
	"os"
)

type Item struct {
	Name        string
	Description string
	Price       float64
}

type Meal struct {
	Name  string
	Items []Item
}

type Resturant struct {
	Name string
	Menu []Meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	returants := []Resturant{
		Resturant{
			Name: "Bobs Diner",
			Menu: []Meal{
				Meal{
					Name: "Breakfast",
					Items: []Item{
						Item{
							Name:        "Bacon & Eggs",
							Description: "Crispy bacon with eggs how you would like with bread of your choice",
							Price:       5.99,
						},
						Item{
							Name:        "Baked beans on toast",
							Description: "Baked beans on bread of your choice",
							Price:       3.99,
						},
					},
				},
				Meal{
					Name: "Lunch",
					Items: []Item{
						Item{
							Name:        "BLT",
							Description: "Bacon Lettuce Tomato",
							Price:       7.99,
						},
						Item{
							Name:        "Steak samwitch",
							Description: "Steak how you want in bread that you want",
							Price:       8.99,
						},
					},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, returants)
	if err != nil {
		log.Fatalln(err)
	}
}
