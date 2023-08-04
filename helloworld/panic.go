//graceful system shutdown statements --> defer, panic, recover


package main

import "fmt"

func handlepanic(){
	if a := recover(); a!=nil{
		fmt.Println("3. RECOVER from - ",a)
	}
	fmt.Println("4. I am stopping here by closing all connections")
}

func entry(language *string, name *string){
	//2

	fmt.Println("1-Entry Regular")//1
	defer handlepanic()
	defer docleanup()
	if language == nil{
		panic("Error: Language Cannot be nill")
	}

	if name == nil{
		panic("panic - because of name") //4
	}
	fmt.Println("language",language)
	fmt.Println("name",name)
}

func docleanup(){
	fmt.Println("2. defer-entry-cleanup")
}

func main(){
	defer fmt.Println("I am the first defer")
	defer fmt.Println("6.defer- in main")
	lang := "go lang"
	entry(&lang,nil)
	fmt.Println("5.main-normal",lang)
}

