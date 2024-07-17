package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template  

type user struct {
	Name string
	Age  int
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		u := user{
			Name: "John Doe",
			Age:  25,
		}

		templates.ExecuteTemplate(w, "index.html", u)
	})

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
