package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates/index.html
var f embed.FS

type PageVariables struct {
	Name string
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := PageVariables{
		Name: r.URL.Path[1:],
		Date: now.Format("Jan 2, 2006"),
		Time: now.Format("3:04:05 PM"),
	}

	t, err := template.ParseFS(f, "templates/index.html") //parse the html file homepage.html
	if err != nil {                                       // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
