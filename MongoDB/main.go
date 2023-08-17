package main

import (
	"MongoDb/services"
	"fmt"
)


func main(){

    fmt.Println("MongoDB successfully connected..")
    //storing the documents in transactions var
    // transactions,_ := services.FindTransaction() //retrievng documents using find query implemented in services package
    // for ind,transaction := range(transactions){
    //     //retrieving each documents and print their account_id
    //     fmt.Println(ind," ",transaction)
    // }
    //services.FindTransaction()
    services.UpdateTransaction()
}