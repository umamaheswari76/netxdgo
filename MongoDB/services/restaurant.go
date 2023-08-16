package services

import (
	"MongoDb/config"
	"MongoDb/models"
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

func RestaurantContext()*mongo.Collection{
	client,_ := config.ConnectDataBase()
	return config.GetCollection(client,"sample_restaurants","restaurants")
}

func FindRestaurant() ( []*models.Restaurant, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)	
	filter := bson.D{} 	
	limit:=options.Find().SetLimit(10)
	result,err := RestaurantContext().Find(ctx, filter,limit)
	
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else {
		fmt.Println(result)
		var restaurants []*models.Restaurant

		for result.Next(ctx){
			restaurant := &models.Restaurant{}
			err := result.Decode(restaurant)

			if err!=nil{
				return nil,err
			}
			//fmt.Println(product)
			restaurants = append(restaurants, restaurant)
		}
		if err := result.Err();err!=nil{
			return nil,err
		}
		if len(restaurants) == 0{
			return []*models.Restaurant{},nil
		}
		return restaurants, nil
	}

}