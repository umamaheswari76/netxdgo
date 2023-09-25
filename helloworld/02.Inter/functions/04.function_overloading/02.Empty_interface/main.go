package main
import (
	"fmt"
	"reflect"
)

//Different number of parameters and of different types
//Empty Interface – It is an interface without any methods.

// func handle(params ...interface{}){
// 	for _,v := range params{

// 		fmt.Println(v)
// 	}
// }

// func main(){
// 	handle(1,"First")
// }


//practiced
func handle(params ...interface{})([]int, []string){
	var a []int
	var s []string
	for _,v := range params{
		if reflect.TypeOf(v)==reflect.TypeOf(int(0)){
			a = append(a, v.(int)) //refer type assertion
			//changes the interface type(v) into int type
		}
		if reflect.TypeOf(v)==reflect.TypeOf(string("")){
			s = append(s, v.(string))
		}
		//fmt.Println(v)
	}
	return a,s
}
func main(){
	i,s := handle(1,"First","second",2)
	fmt.Println(i)
	fmt.Println(s)

}