package main
  
import "fmt"
  
// Main function
// func main() {
// 	str := "hello"
// 	rne := []rune(str)
// 	str1 := string(rne)
// 	// fmt.Printf("%c",rne[0])
// 	//fmt.Println(str1)
// 	//fmt.Println(len(str1))
// 	//for i:=0;i<len(str1);i++{
// 		// fmt.Println(string(rne[i]))
// 		// if rne[i]=='e'{
// 		// 	// fmt.Printf("%c",rne[i])
// 		// 	fmt.Println(string(rne[i]))
// 		// 	break
// 		// }
// 		c:='e'
// 		d:=str[0]
// 		if 'c'<=d{
// 		//	fmt.Println( string(str[i])," ",i)
// 		}
// //	}

// }

func reverse(str string)(string){
	//hello --> olleh
	var rne []rune=
	var res string = ""
	for i:=len(str);i>-1;i--{
		res += string(str[i])
	}
	return res
}

func  main(){
	str := "hello"
	res := reverse(str)
	fmt.Println(res)
}