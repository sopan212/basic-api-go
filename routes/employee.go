package routes

import (
	"basicAPI/config"
	"basicAPI/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	rows, err := db.Query("SELECT id,name,phone,address from employees")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var resp []models.Employe
	for rows.Next() {
		var data = models.Employe{}
		err := rows.Scan(&data.Id, &data.Name, &data.Phone, &data.Address)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Printf("errors")
			return
		}

		resp = append(resp, data)
		json.NewEncoder(w).Encode(resp)

	}
	fmt.Println(resp)
}

func PostEemployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var employee models.Employe

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("INSERT into employees(name,phone,address)values(?,?,?)", employee.Name, employee.Phone, employee.Address)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("insert success")

		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}
func PutEemployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var employee models.Employe

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("UPDATE employees SET name=?,phone=?,address=? WHERE id= ?", employee.Name, employee.Phone, employee.Address, employee.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("updated success")

		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}

func DeleteEemployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		db, err := config.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return

		}
		defer db.Close()

		var employee models.Employe

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = db.Exec("DELETE from employees WHERE id= ?", employee.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("delete success")

		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
	}
}
