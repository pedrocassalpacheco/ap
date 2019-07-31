package main

import (
	. "./services"
	"./trace"
	"log"
	"net/http"
)

func main() {

	log.Println("Server started on: http://localhost:8080")
	log.Printf("Tracing information %T", trace.GoSensor)
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/employees", trace.GoSensor.TracingHandler("employees", Employees))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

