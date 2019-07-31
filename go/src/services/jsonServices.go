package services

import (
	"../trace"
	"../types"
	"../util"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"time"

	//"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
)


func Employees(w http.ResponseWriter, r *http.Request) {

	// Fetch all employees
	//
	connectionSpan := trace.TraceDBConnection(r, "Database connection")
	db := util.DBConn()
	connectionSpan.Finish()

	trace.TraceFunctionExecution(r, func () {
		time.Sleep(1*time.Second)
	}, "Going to sleep")

	//time.Sleep(1 * time.Second)
	//
	//
	sql := "SELECT * FROM Employee ORDER BY id DESC"
	sqlSpan := trace.TraceSQLExecution(r, sql, "SQL execution")
	selDB, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	emp := types.Employee{}
	var res []types.Employee
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	time.Sleep(5*time.Second)
	sqlSpan.Finish()

	// Marshall
	bytes, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	} else {
		w.Write(bytes)
		log.Println(string(bytes))
		http.Redirect(w, r, "/", 200)
	}

	defer db.Close()
}