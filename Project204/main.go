package main

import "fmt"

func main(){
	var tmp []func()()=make([]func()(),0,3)
	for i:=0;i<3;i++{
		var b int=0
		var c func()=func()(){
			fmt.Println(b)
		}
		tmp=append(tmp,c)
		fmt.Printf("%x  %x\n",&i,&b)
	}
	panic("")
	

}