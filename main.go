package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	// Render the project html page
	http.ServeFile(w, r, "static/projects.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/projects", projectsPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
