package main

import (
	"fmt"
	l "github.com/retgits/creditcard"
)
func main(){
	c := l.Card{"Visa", "969696969696", 12, 2023, "564"}
	validation := c.Validate()
    fmt.Printf("%+v\n", validation)
    fmt.Printf("%+v\n", validation.Card)

}