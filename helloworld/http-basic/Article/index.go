package main

import (
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

//this struct can be modified to json
//golang to json mapping
type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author string `json:"Author"`
}

func main(){
	http.HandleFunc("/products",handleProducts)
	http.HandleFunc("/contactus",handleContacts)
	http.HandleFunc("/article",handleArticle)
	http.HandleFunc("/getArticles",getArticles)
	log.Fatal(http.ListenAndServe(":3001",nil))//wrapping to manage exceptions
}

func handleProducts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprint(w, "<h1>Welcome to  products</h1>")
}

func handleContacts(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	fmt.Println(req.URL.Path)
	fmt.Fprint(w,"<h1>xyz@gmail.com</h1>")
}


//json data in serialized format
//we have to provide mapping for that format 
//unmarshal - converting json obj/struct to golang obj/struct
//marshal - reverse of the above

func handleArticle(w http.ResponseWriter, req *http.Request){

	if req.Method == "POST"{
		reqBody, _ := io.ReadAll(req.Body)
			var post Article
			err:=json.Unmarshal(reqBody, &post)
			json.NewEncoder(w).Encode(post)
			// newData, err2 := json.Marshal(post)
			post.Author = "Jhon"
			if err!=nil{
				fmt.Fprint(w, err.Error())
			}else{
				json.NewEncoder(w).Encode(post)
				
			}
		}else{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Unable to process your request")

		}
	}


func getArticles(w http.ResponseWriter, req * http.Request){
	if req.Method == "POST"{
		reqBody, _ := io.ReadAll(req.Body)
	var articles []Article
	err:=json.Unmarshal(reqBody, &articles)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
	}else{
		// fmt.Fprintf(w, "Succes")
		fmt.Println(articles)
		articles = append(articles,Article{Id: "3",Title:"MIB"})
		json.NewEncoder(w).Encode(articles)
	}
}
}