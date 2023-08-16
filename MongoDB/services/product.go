package services

import (
	"MongoDb/config"
	"MongoDb/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
//	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func ProductContext()*mongo.Collection{
	client,_ := config.ConnectDataBase()
	return config.GetCollection(client,"inventory","products")
}


// func InsertProduct(){
// 	var product models.Product
// 	product.Name = "iphone"
// 	product.Price = 115000
// 	product.Description = "It is awesome"
// 	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
// 	// client,_ := config.ConnectDataBase()
// 	// var productCollection *mongo.Collection = config.GetCollection(client,"inventory","products")
// 	//result,err := productCollection.InsertOne(ctx,product)
// 	result,err := ProductContext.InsertOne(ctx,product)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }
//}

func InsertProduct(product models.Product)(*mongo.InsertOneResult,error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,err := ProductContext().InsertOne(ctx,product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	return result,nil
}

func InsertProductList(products []interface{})(*mongo.InsertManyResult,error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,err := ProductContext().InsertMany(ctx,products)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	return result,nil
}


// func FindProducts() ( []*models.Product, error){
// 	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
// 	//filter := bson.D{} //bson.D -> converts the filter cond to go understanble format
// 	filter := bson.D{{"name","samsung"}} 
// 	//products []models.Product
// 	//filter -> fetch all records
// 	result,err := ProductContext().Find(ctx, filter)
// 	if err!=nil{
// 		fmt.Println(err.Error())
// 		return nil,err
// 	}else {
// 		//do something
// 		fmt.Println(result)
// 		//build the array of prodcts for the cursor that we received
// 		var products []*models.Product

// 		for result.Next(ctx){
// 			product := &models.Product{}
// 			err := result.Decode(product)

// 			if err!=nil{
// 				return nil,err
// 			}
// 			//fmt.Println(product)
// 			products = append(products, product)
// 		}
// 		if err := result.Err();err!=nil{
// 			return nil,err
// 		}
// 		if len(products) == 0{
// 			return []*models.Product{},nil
// 		}
// 		return products, nil
// 	}

// }

func FindProducts() ( []*models.Product, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	//filter := bson.D{} //bson.D -> converts the filter cond to go understanble format
	filter := bson.D{{"name","samsung"}} 
	//products []models.Product
	//filter -> fetch all records
	result,err := ProductContext().Find(ctx, filter)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else {
		//do something
		fmt.Println(result)
		//build the array of prodcts for the cursor that we received
		var products []*models.Product

		for result.Next(ctx){
			product := &models.Product{}
			err := result.Decode(product)

			if err!=nil{
				return nil,err
			}
			//fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err();err!=nil{
			return nil,err
		}
		if len(products) == 0{
			return []*models.Product{},nil
		}
		return products, nil
	}

}