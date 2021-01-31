package models

import (
	"Uber/config"
	"database/sql"
	"fmt"
	"math"
	"strconv"

	"github.com/twinj/uuid"
)

type SearchCab struct {
	latitude  string `json:"latitude"`
	longitude string `json:"longitude"`
}

type CabList struct {
	DriverId  uuid.UUID
	Name      string
	Mobile    string
	Address   string
	CarNo     string
	Latitude  string
	Longitude string
}

func SearchCabWithInTwoKM(cab SearchCab) ([]CabList, error) {
	var NearAllCabList []CabList
	db := config.GetDB()
	rows, err := db.Query("SELECT COUNT(*) as count FROM  drivers")
	fmt.Println("Total count:", checkCount(rows))
	checkErr(err)

	rows, err = db.Query("SELECT driver_id, name, mobile,address,car_no, c_latitude,c_longitude FROM drivers LIMIT $1", checkCount(rows))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cablist CabList
		err = rows.Scan(&cablist.DriverId, &cablist.Name, &cablist.Mobile, &cablist.Address, &cablist.CarNo, &cablist.Latitude, &cablist.Longitude)
		if err != nil {
			// handle this error
			panic(err)
		}
		res := haversine(cab.latitude, cab.longitude, cablist.Latitude, cablist.Longitude)
		if res > 20.0 || res < 100 {
			NearAllCabList = append(NearAllCabList, cablist)
		}
	}
	checkErr(err)

	return NearAllCabList, nil
}

//haversign method
func haversine(lat1 string, lon1 string, lat2 string, lon2 string) float64 {
	lat11, _ := strconv.ParseFloat(lat1, 64)
	lat22, _ := strconv.ParseFloat(lat2, 64)

	lon11, _ := strconv.ParseFloat(lat1, 64)
	lon22, _ := strconv.ParseFloat(lat2, 64)

	dLat := (lat22 - lat11) * float64(math.Pi) / 180.0
	dLon := (lon22 - lon11) * float64(math.Pi) / 180.0

	lat111 := (lat11) * float64(math.Pi) / 180.0
	lat222 := (lat22) * float64(math.Pi) / 180.0

	a := (math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat111)*math.Cos(lat222))
	rad := float64(6371)
	c := 2 * math.Asin(math.Sqrt(a))
	return rad * c
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BookYorRide(cab CabList) (*CabList, error) {
	return &cab, nil

	//your ride car data will be entere in booking history details
}

//https://www.movable-type.co.uk/scripts/latlong.html
