package models

import (
	"Uber/config"
	"math"
	"strconv"

	"github.com/twinj/uuid"
)

type SearchCab struct {
	latitude  string `json:"latitude"`
	longitude string `json:"longitude"`
}

type CabList struct {
	DriverId  uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Address   string    `json:"address"`
	CarNo     string    `json:"carno"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
}

type BookedRide struct {
	Cab        CabList
	RiderName  string
	UserName   string
	UserMobile string
}

func SearchCabWithInTwoKM(cab SearchCab) ([]CabList, error) {
	var NearAllCabList []CabList
	db := config.GetDB()

	rows, err := db.Query("SELECT driver_id, name, mobile,address,car_no, c_latitude,c_longitude FROM drivers order by driver_id desc") //LIMIT $1", checkCount(rows))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cablist CabList
		err = rows.Scan(&cablist.DriverId, &cablist.Name, &cablist.Mobile, &cablist.Address, &cablist.CarNo, &cablist.Latitude, &cablist.Longitude)
		if err != nil {
			panic(err)
		}
		//fmt.Println("*****", cablist.DriverId)
		res := haversine(cab.latitude, cab.longitude, cablist.Latitude, cablist.Longitude)
		//fmt.Println("Result ===", res)
		if res >= 0.0 && res < 4000.9 {
			NearAllCabList = append(NearAllCabList, cablist)
		}
	}
	checkErr(err)
	//fmt.Println(NearAllCabList)
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BookYorRide(cab CabList, userId uint64) (BookedRide, error) {
	db := config.GetDB()
	sqlStatement := `SELECT name, username, mobile FROM users where user_id=$1`
	row := db.QueryRow(sqlStatement, userId)

	var user Signup
	err := row.Scan(&user.Name, &user.Username, &user.Mobile)
	if err != nil {
		panic(err)
	}
	CabBooked := BookedRide{
		Cab:        cab,
		RiderName:  user.Name,
		UserName:   user.Username,
		UserMobile: user.Mobile,
	}

	return CabBooked, nil
}

//https://www.movable-type.co.uk/scripts/latlong.html
