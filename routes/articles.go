package routes

import (
	"basicAPI/config"
	"basicAPI/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()
		rows, err := db.Query("SELECT id,title,author from articles")

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		defer rows.Close()

		var resp models.Articles
		for rows.Next() {
			var data = models.Article{}
			var err = rows.Scan(&data.Id, &data.Title, &data.Author)
			if err != nil {
				w.Write([]byte(err.Error()))
				log.Printf("errors")
				return
			}

			resp = append(resp, data)
		}
		json.NewEncoder(w).Encode(resp)

		fmt.Println("get succes")
	}

}
func PostArticles(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var article models.Article

		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("INSERT into articles(title,author)values(?,?)", article.Title, article.Author)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("insert success")

		json.NewEncoder(w).Encode(article)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}
func PutArticles(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var article models.Article

		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = db.Exec("UPDATE articles SET title=?,author=? WHERE id= ?", article.Title, article.Author, article.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("updated success")

		json.NewEncoder(w).Encode(article)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}

func DeleteArticles(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var article models.Article

		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("DELETE from articles WHERE id= ?", article.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("delete success")

		json.NewEncoder(w).Encode(article)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}
