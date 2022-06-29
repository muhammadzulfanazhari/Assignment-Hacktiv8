package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alok87/goutils/pkg/random"
)

type Data struct {
	WindSpeed   int
	WaterHeight int
	WindStatus  string
	WaterStatus string
}

type Weather struct {
	Status Data
}

func main() {
	dataWeather := &Weather{}
	http.HandleFunc("/", dataWeather.display)
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	http.ListenAndServe(":5000", nil)
}

func (dataWeather *Weather) display(w http.ResponseWriter, r *http.Request) {
	open, err := os.Open("weatherDataInfo.json")

	if err != nil {
		fmt.Println("Cannot open the file!")
	}

	jsonData, err := ioutil.ReadAll(open)
	if err != nil {
		fmt.Println("Cannot parse the file!")
	}

	if err := json.Unmarshal(jsonData, &dataWeather); err != nil {
		fmt.Println("Error unmarshal", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	waterNow := random.RangeInt(1, 10, 1)
	dataWeather.Status.WaterHeight = waterNow[0]

	windNow := random.RangeInt(1, 20, 1)
	dataWeather.Status.WindSpeed = windNow[0]

	jsonMarshal, _ := json.Marshal(dataWeather)
	err = ioutil.WriteFile("weatherDataInfo.json", jsonMarshal, 0644)

	// check the height of water
	if dataWeather.Status.WaterHeight <= 5 {
		dataWeather.Status.WaterStatus = "Aman"
	} else if dataWeather.Status.WaterHeight >= 6 && dataWeather.Status.WaterHeight <= 8 {
		dataWeather.Status.WaterStatus = "Siaga"
	} else {
		dataWeather.Status.WaterStatus = "Bahaya"
	}

	// check the speed of wind
	if dataWeather.Status.WindSpeed <= 6 {
		dataWeather.Status.WindStatus = "Aman"
	} else if dataWeather.Status.WindSpeed >= 7 && dataWeather.Status.WindSpeed <= 15 {
		dataWeather.Status.WindStatus = "Siaga"
	} else {
		dataWeather.Status.WindStatus = "Bahaya"
	}

	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dataWeather.Status)
}
