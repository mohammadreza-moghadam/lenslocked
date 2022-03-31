package main

import (
	"html/template"
	"os"
)

type Profile struct {
	Email    string
	Username string
}

type User struct {
	Name    string
	Dog     string
	Age     int
	Float   float64
	Slice   []string
	Names   map[string]string
	Profile Profile
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "John Doe",
		Dog:   "Bark Dog",
		Age:   12,
		Float: 3.14,
		Slice: []string{"a", "b", "c"},
		Names: map[string]string{
			"first":  "Mohammad",
			"middle": "Reza",
			"last":   "Moghadam",
		},
		Profile: Profile{
			Email:    "mohammadreza.che.eng@gmail.com",
			Username: "M_R_M_1997",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
