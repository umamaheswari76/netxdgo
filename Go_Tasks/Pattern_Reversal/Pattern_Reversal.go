package main

import "fmt"


func rev(str string) string {
    c := 0
    rune_str := []rune(str)
    for i:=0;i<len(rune_str);i++ {
        if i>1 && c==-2 {
            c++
            if i+1 < len(rune_str) {
                c++
                i++
                temp := rune_str[i]
                rune_str[i] = rune_str[i-1]
                rune_str[i-1] = temp
            }
        }else{
            c--
        }
    }
    return string(rune_str)
    
}

func main(){
     s := "ABCDEFGHIJKL"
    fmt.Printf(rev(s))
}
