package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type employeeDetail struct {
	Name        string
	Designation string
	Age         int
}

var empData []employeeDetail

func employeeHandler(res http.ResponseWriter, req *http.Request) {
	//var newEmployee employeeDetail
	switch req.Method {
	case http.MethodGet:
		res.Header().Set("Content-Type", "application/json")
		empDataJSON, _ := json.Marshal(empData)
		res.Write(empDataJSON)
	case http.MethodPost:
		var newEmp employeeDetail
		body, _ := ioutil.ReadAll(req.Body)
		_ = json.Unmarshal([]byte(body), &newEmp)
		empData = append(empData, newEmp)
		empDataJSON, _ := json.Marshal(empData)
		ioutil.WriteFile("employee.data", empDataJSON, os.ModePerm)
	}

}

// func employeeUpdateHandler(res http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case http.MethodPut:
// 		requestBody := req.Header.Get("id")

// 		//var name string
// 		//_ = json.Unmarshal(requestBody, &name)
// 		fmt.Println(requestBody)
// 		// empData = append(empData, newEmp)
// 		// empDataJSON, _ := json.Marshal(empData)
// 		// ioutil.WriteFile("employee.data", empDataJSON, os.ModePerm)
// 	}
// }

func main() {
	file, _ := ioutil.ReadFile("employee.data")
	_ = json.Unmarshal([]byte(file), &empData)
	http.HandleFunc("/employee", employeeHandler)
	//http.HandleFunc("/employee/{id}", employeeUpdateHandler)
	http.ListenAndServe(":5000", nil)
}
