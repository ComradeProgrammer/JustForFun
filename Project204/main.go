package main

import "fmt"

func main(){
	/*var a *mystl.Vector=new(mystl.Vector)
	a.PushBack(0,1,2,3)
	fmt.Println(a)
	a.Insert(4,5)
	fmt.Println(a)*/
	var a []int=[]int{1,2,3,4,5}
	copy(a[1:],a[0:])
	fmt.Println(a)


}