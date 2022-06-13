package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	
	fmt.Println("Initializing the handlers...")
	initializeHandlers(fileServer)

	fmt.Println("Starting to listen at port 9000...")

	if err:=http.ListenAndServe(":9000", nil); err!=nil{
		log.Fatal(err)
	}
}


func userDetailHandler(w http.ResponseWriter, r *http.Request){
	if r.Method!="POST"{
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err!=nil{
		fmt.Fprintf(w, "Error Parsing form %v", err)
		return
	}

	name:=r.FormValue("name")
	email:=r.FormValue("email")
	address:=r.FormValue("address")
	password:=r.FormValue("password")

	fmt.Fprintf(w, "Name is %v\n", name)
	fmt.Fprintf(w, "Email is %v\n", email)
	fmt.Fprintf(w, "Address is %v\n", address)
	fmt.Fprintf(w, "Password is %v\n", password)
}

func contactUsPageHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/contactUs" {
		http.Error(w, "404 URL not found", http.StatusNotFound)
		return
	}

	if r.Method!="GET"{
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Contact Us page reached")
}

func initializeHandlers(fileServer http.Handler){
	http.Handle("/", fileServer)
	http.HandleFunc("/form", userDetailHandler)
	http.HandleFunc("/contactUs", contactUsPageHandler)
}