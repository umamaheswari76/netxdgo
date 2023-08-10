package main

import (
    "fmt"
    "math")


var n int = 4


func print(dist[4] int){
    fmt.Println("Shortest distance from source")
    for i:=0; i<n;i++{
        fmt.Println(i,"\t",dist[i])
    }
}
func min_index(dist [4]int, sptset [4]bool) int{
    var  min_value = math.MaxInt
    var min_index int
    for i:=0;i<n;i++{
        if !sptset[i] && dist[i]<=min_value{
            min_value = dist[i]
            min_index = i
        }
    }
    return min_index
}


func djikstra(graph [4][4]int, src int){
    var sptset [4]bool
    var dist [4]int


    for i:=0;i<n;i++{
        sptset[i]=false
        dist[i]=math.MaxInt
    }
    dist[src] = 0


    for count:=0;count<n;count++{
        element := min_index(dist,sptset)
        sptset[element] = true


        for con_ind:=0; con_ind<n; con_ind++{
            if (!sptset[con_ind] && graph[element][con_ind]!=0 && dist[element] != math.MaxInt && (dist[element]+graph[element][con_ind]) <dist[con_ind]){
                dist[con_ind] = dist[element]+graph[element][con_ind]
            }
        }
    }


    print(dist)
   
}

func main(){
	var graph = [4][4] int {{0,5,8,0},
	{5,0,9,2},
	{8,9,0,6},
	{0,2,6,0}}
    var src int
    fmt.Scanln(&src)
    djikstra(graph, src)
}

