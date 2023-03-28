package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string
	Age  uint16
	Id   int16
}

/* func (p Person) getInformation() string {
	return fmt.Sprintf("Person Name is: %s, He is %d age, His Id: %d", p.Name, p.Age, p.Id)
} */

func mainPage(w http.ResponseWriter, r *http.Request) {
	Tom := Person{Name: "Tom", Age: 28, Id: 1}
	//fmt.Fprintf(w, Tom.getInformation())

	tmpl, _ := template.ParseFiles("HTML/mainPage.html")
	tmpl.Execute(w, Tom)
}

func profilesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profiles")
}

func HandleFunc() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/Profiles", profilesPage)
	http.ListenAndServe(":8888", nil)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)

	fileServer := http.FileServer(http.Dir("CSS"))

	router.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))

	HandleFunc()
}
