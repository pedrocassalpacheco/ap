package main

import (
	. "./services"
	"./trace"
	"./util"
	"fmt"
	"log"
	"net/http"
)

func init() {
	util.Configure()
}
func main() {

	// Arguments
	port := util.Configure().Port
	log.Println("Server started on: http://localhost:", port)
	log.Printf("Using tracer %T", trace.GoSensor)
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/employees", trace.GoSensor.TracingHandler("employees", Employees))
	http.HandleFunc("/step2", trace.GoSensor.TracingHandler("employees", Step2))
	http.HandleFunc("/error", trace.GoSensor.TracingHandler("error", Error))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
