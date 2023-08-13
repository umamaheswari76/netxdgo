//
package main
  
import (
	"fmt"
	"strconv"
)

// main function
func main() {
	var str_num string
	fmt.Println("Enter a number to convert: ")
	fmt.Scanln(&str_num)
	int_form,_ := strconv.ParseInt(str_num,10,64)
	// int_form := int(str_num)
	fmt.Println("int format is: ",int_form+100)
	i := 10
	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)
	fmt.Printf("%v, %v\n", s1, s2)
}