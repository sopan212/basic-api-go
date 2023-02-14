package main

import (
	"basicAPI/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/employees", routes.GetEmployees)
	http.HandleFunc("/employees/submit", routes.PostEemployees)
	http.HandleFunc("/employees/update", routes.PutEemployees)
	http.HandleFunc("/employees/delete", routes.DeleteEemployees)

	http.HandleFunc("/articles", routes.GetArticles)
	http.HandleFunc("/articles/submit", routes.PostArticles)
	http.HandleFunc("/articles/update", routes.PutArticles)
	http.HandleFunc("/articles/delete", routes.DeleteArticles)

	log.Printf("server running on :3000")
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing home"))
}
