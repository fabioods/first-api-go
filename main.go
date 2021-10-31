package main

import (
	"encoding/json"
	"net/http"
)

type Car struct {
	Name string
	Year int
}

func main() {
	http.HandleFunc("/", Init)
	http.HandleFunc("/car", ShowCar)
	http.ListenAndServe(":5000", nil)
}

func Init(write http.ResponseWriter, req *http.Request) {
	write.Write([]byte("{\"hello\": \"world\"}"))
}

func ShowCar(write http.ResponseWriter, req *http.Request) {
	car := Car{
		Name: "Ford",
		Year: 2020,
	}
	jsonCar, _ := json.Marshal(car)
	write.Write(jsonCar)
}
