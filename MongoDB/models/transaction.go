package models

import (
	//"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Transactions struct{
	Date primitive.DateTime `json: date bson:"date",required`
	Amount int `json:"amount" bson:"amount",required`
	Transaction_code string `json:"transaction_code" bson:"transaction_code",required`
	Symbol string `json:"symbol" bson:"symbol",required`
	Price string `json:"price" bson:"price",required`
	Total string `json:"total" bson:"total",required`
}

type Transaction struct{
	ID primitive.ObjectID `json:"_id bson:"_id"`
	Account_id int `json: "account_id" bson: "account_id", required`
	Transaction_count int `json:"transaction_count" bson: "transaction_count",required`
	Bucket_start_date primitive.DateTime `json: "bucket_start_date"  bson: "bucket_start_date" ,required`
	Bucket_end_date primitive.DateTime`json: "bucket_end_date" bsn "bucket_end_date", required `
	//Trans []Transactions `json:"transactions" bson:"transactions",required`
}
