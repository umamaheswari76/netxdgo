package services

import (
	"MongoDb/config"
	//"MongoDb/models"
	"context"
	"fmt"
	"time"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

//used to connecting with the db and returns the collection
func TransactionContext()(*mongo.Collection){
	client,_:= config.ConnectDataBase()
	return config.GetCollection(client, "sample_analytics","transactions")
}

//returns array of transactions, err
func FindTransaction(){

	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	// filter := bson.M{"transaction_count":bson.D{{"$gte",85},{"$lt",90}}}
	// options := options.Find().SetSort(bson.D{{"transaction_count",1}}).SetSkip(30).SetLimit(10)
	//options := options.Find().
	//limit:=options.Find().SetLimit(10)

	matchStage := bson.D{{"$match", bson.D{{"transaction_count",100}}}}
	groupStage := bson.D{
		{"$group",bson.D{
			{"_id","$account_id"},
			{"total_count",bson.D{{"$sum","$transactions.amount"}}},
		}}}
	result,err := TransactionContext().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err!=nil{
		fmt.Println((err.Error()))
	}else{
	var showsWithInfo []bson.M
	if err=result.All(ctx, &showsWithInfo); err!=nil{
		panic(err)
	}
	//fmt.Println(showsWithInfo)
	formatted_data,err := json.MarshalIndent(showsWithInfo,""," ")
	if err!= nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println(string(formatted_data))
	}
	}
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// 	return nil,err
	// }else{
	// 	//fmt.Println(result)
	// 	var transactions []*models.Transaction

	// 	//.Next() -> used to move to next item
	// 	for result.Next(ctx){
	// 		transaction := &models.Transaction{}
	// 		//decode 
	// 		err := result.Decode(transaction)

	// 		if err!=nil{
	// 			return nil,err
	// 		}
	// 		transactions = append(transactions,transaction)
	// 	}
	// 	if err := result.Err();err!=nil{
	// 		return nil,err
	// 	}
	// 	if len(transactions) == 0{
	// 		return []*models.Transaction{},nil
	// 	}
	// 	return transactions, nil
	// }

}

func UpdateTransaction(initialValue int,newValue int)(*mongo.UpdateResult,error){
	ctx,_ := context.WithTimeout(context.Background(),100*time.Second)
	filter := bson.D{{"accont_id",initialValue}}
	update := bson.D{{"$set",bson.D{{"account_id",newValue}}}}
	result,err := TransactionContext().UpdateOne(ctx,filter,update)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return result, nil
}