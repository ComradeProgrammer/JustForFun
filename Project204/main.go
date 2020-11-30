package main

import (
	"./mystl"
	"fmt"
)

type fuck struct{

}


func main(){
	var a *mystl.Vector=new(mystl.Vector)
	a.PushBack(0,1,2,3,4,5,6,7,8,9)
	fmt.Println(a)
	var itr mystl.Iterator =a.Iterator(5)
	fmt.Println(itr.This())//5
	fmt.Println(itr.Next())//6
	fmt.Println(itr.This())//6
	fmt.Println(itr.Back())//5
	fmt.Println(itr.Back())//4
	fmt.Printf("%s\n",5)

}