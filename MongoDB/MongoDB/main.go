package main

import (
	//"MongoDb/config"
//	"MongoDb/models"
	"MongoDb/services"

	//"MongoDb/constants"
	//"context"
	"fmt"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

//mongoclient as global var
var(
    mongoClient *mongo.Client
)

func main(){
    //    ctx:=context.TODO()

    //    mongoconn:=options.Client().ApplyURI(constants.ConnectionString)
    //    mongoclient,err:=mongo.Connect(ctx,mongoconn)
    //    if err!=nil{
    //     panic(err)
    //    }
    //    if err:=mongoclient.Ping(ctx,readpref.Primary()); err!=nil{
    //     panic(err)
    //    }
    // client,_ := config.ConnectDataBase()
    // collection := config.GetCollection(client,"sample_training","zips")
       fmt.Println("MongoDB successfully connected..")
    //    product := models.Product{ID: primitive.NewObjectID(),Name: "Samsung",Price :1}
    //    services.InsertProduct(product)
   
    // products := [] interface{}{
    //     models.Product{ID:primitive.NewObjectID(),Name:"oneplus",Price:1000000,Description: "this is oneplus"},
    //     models.Product{ID:primitive.NewObjectID(),Name:"oneplus",Price:1000000,Description: "this is oneplus"},
    // }
    // services.InsertProductList(products)

    // products,_ := services.FindProducts()
    // for _,product := range products{
    //     fmt.Println(product)

    // }
    restaurants,_ := services.FindRestaurant()
    for _,restaurant := range restaurants{
        fmt.Println(restaurant.Name)

    }

}