package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
)

func main(){
	a := bufio.NewScanner(os.Stdin)
	a.Scan()
	fmt.Printf("%T",a)
	// inp := a.Text()

}