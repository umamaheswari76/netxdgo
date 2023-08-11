package main

import (
	"fmt"
	"net/http"
	"log"
)

//net/tcp - intranet 
//net/http

func main(){
	// http.HandleFunc("/api/hello",handleHello) //route
	http.HandleFunc("/",handleHello) //route
	http.HandleFunc("/products",handleProducts) 
	http.HandleFunc("/contact",handleContact)
	//http.ListenAndServe(":3000",nil)
	log.Fatal(http.ListenAndServe(":3000",nil))//wrapping to manage exceptions

}

func handleProducts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to  products</h1>")

}

func handleHello(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound) here status not found is a status code..
		//so it is the coder responsibility to set proper response codes for the given request
		//status codes - 500 - internal server error(request is correct but logic failed internaly)
		// 404 - page not found
		//403 - forbidden (route is correct but to acces the patcular path bt there is no access for that page...it is forbidden)
		fmt.Fprintf(w, "<h1>Page Not Found</h1>")
	}else{
		fmt.Fprintf(w,"<h1>Welcome to HTTP</h1>")
	}

}

func handleContact(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprintf(w,"<h1>xyz@gmail.com</h1>")
}

//http://localhost:3000/api/hello ---> running in web browser will print the above Fprintf stmt