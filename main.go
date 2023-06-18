package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
)

type Person struct {
	name string
}

func main(){
	
	helloHandler := func( w http.ResponseWriter, req *http.Request ){
		io.WriteString( w, "Hello World!! \n" )
	}

	helloNamePrintForm := func( w http.ResponseWriter, req *http.Request ){
		io.WriteString( w, "Recieved body \n" )

		if err := req.ParseForm(); err != nil {
			log.Println("Error in parsing")
		}

		fmt.Fprintf(w, "req.PostForm = %v\n", req.PostForm)

		name := req.FormValue("name")

		fmt.Fprintf(w, "name := %v", name)
	}

	helloNamePrintJson := func( w http.ResponseWriter, req *http.Request ){
		var p Person
		
		io.WriteString( w, "Recieved body \n" )

		err := json.NewDecoder(req.Body).Decode(&p)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Fprintf(w, "req.PostForm = %v", p)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/name-form", helloNamePrintForm)
	http.HandleFunc("/name-json", helloNamePrintJson)

	log.Println("Listing for requests at http://localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}