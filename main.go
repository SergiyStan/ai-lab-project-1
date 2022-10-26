package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		vehicles, err := dbGetVehicles()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, vehicles)
	}
}

func rollHandlerByBrand(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("search_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		brand := r.Form.Get("brand")
		vehicles, err := dbGetVehiclesByBrand(brand)
		if err != nil {
			log.Fatal(err)
		}

		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, vehicles)

	}

}

func addVehicleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		brand := r.Form.Get("brand")
		country := r.Form.Get("country")
		price := r.Form.Get("price")
		year := r.Form.Get("year")
		err := dbAddVehicles(brand, country, price, year)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "6942"
		fmt.Println(port)
	}
	return ":" + port
}

func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/filter", rollHandlerByBrand)
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addVehicleHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
