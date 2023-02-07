package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func Ticker(response http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)   //  for {category}/{itemnumber}
	fmt.Print(vars)   // gives a map

	var body map[string]string 
	json.NewDecoder(request.Body).Decode(&body)

	fmt.Print(body)
    
    response.Write([]byte("Hello World"))
	fmt.Fprintf(response, "Hahaha")
	
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/Hello/{category}", Ticker).Methods("POST")
	http.ListenAndServe(":8000",router)
}
