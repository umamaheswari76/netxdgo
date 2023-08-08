package main

import (
    "fmt"
    "unicode"
    )

func Count_go(str string) (count int){
    count = 0
    rune_str := []rune(str)
    for i:=0;i<len(rune_str);i++{
        
        if string(unicode.ToUpper(rune_str[i]))=="G"{
            if i+1<len(rune_str) && string(unicode.ToUpper(rune_str[i+1]))=="O"{
                count++
                i++
            }
        }
    }
    return count
}

func main(){
    var inp_str string
    fmt.Scanln(&inp_str)
   fmt.Println(Count_go(inp_str))
    
}
