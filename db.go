package main

import (
	"database/sql"
	"fmt"
)

type Vehicle struct {
	Brand   string
	Country string
	Price   string
	Year    string
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "pg"
	DB_NAME     = "ailab1"
)

func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password =%s dbname=%s sslmode = disable",
		DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS vehicles (brand text, country text, price text, build_year text)"); err != nil {
		return err
	}
	return nil
}

func dbAddVehicles(brand, country, price, year string) error {
	sqlstmt := "INSERT INTO vehicles VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, brand, country, price, year)
	if err != nil {
		return err
	}
	return nil
}

func dbGetVehicles() ([]Vehicle, error) {
	var vehicles []Vehicle
	stmt, err := db.Prepare("SELECT brand, country, price, build_year FROM vehicles")
	if err != nil {
		return vehicles, err
	}
	res, err := stmt.Query()
	if err != nil {
		return vehicles, err
	}
	var tempVehicle Vehicle
	for res.Next() {
		err = res.Scan(&tempVehicle.Brand, &tempVehicle.Country, &tempVehicle.Price, &tempVehicle.Year)
		if err != nil {
			return vehicles, err
		}
		vehicles = append(vehicles, tempVehicle)
	}
	return vehicles, err
}

func dbGetVehiclesByBrand(brand string) ([]Vehicle, error) {
	var vehicles []Vehicle
	sqlstmt := "SELECT brand, country, price, build_year FROM vehicles WHERE brand LIKE ($1)"
	res, err := db.Query(sqlstmt, brand)
	if err != nil {
		return vehicles, err
	}
	var tempVehicle Vehicle
	for res.Next() {
		err = res.Scan(&tempVehicle.Brand, &tempVehicle.Country, &tempVehicle.Price, &tempVehicle.Year)
		if err != nil {
			return vehicles, err
		}
		vehicles = append(vehicles, tempVehicle)
	}
	return vehicles, err
}
