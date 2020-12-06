/*
2020/10/31 已经签了火星公司了 准备考完研就去给张一鸣打福报工了
            ^~^
随着go tour的简易教程 粗浅学习go语言啦
*/
package helloworld

import (
	"fmt"
	"runtime"
)

//==========常量声明==========
//常量可以是字符、字符串、布尔值或数值。常量不能用 := 语法声明。
const const1 = 3.14
const (
	const2 = 1 << 100 //数值常量是高精度的 值。
	cons3  = 1 >> 100
)



//=======函数定义==============
// 显式返回的函数定义
func function1(x, y int, z string) (string, string) {
	//第一个括号里面写参数，参数名在前，参数类型在后.连续的多参数类型相同时可以只写最后一个
	//第二个括号里面写返回值类型 返回值可以有多个。如果只有一个返回值不需要写括号
	return z + "x+y is ", fmt.Sprintf("%d", x+y) //显式指出了返回的是什么
}

//隐式返回 函数定义
func function2() (x int) { //函数返回值处直接定义的是有名字的返回值变量
	x = 105 //给这个变量赋值
	return  //返回语句依然不可避免，但是这里无需指明返回值
}

//==========变量声明=============
func function3() {
	//golang的局部变量声明之后必须使用否则编译不过（error）
	var string1, string2 string = function1(1, 102, "") //变量声明的第一种方式，包含了类型和初始值
	var string3 = "\nHello"                             //第二种方式 未声明类型，自动判断
	//以上两种方式是可以在函数作用域和包作用域里
	string4, string5 := " World", "" //第三种方式 连var都省略了，也不需要（也不能）写类型
	//这种方式不能用在包作用域

	fmt.Println(string1 + string2 + string3 + string4 + string5)
	fmt.Println(function2())
}

/*
Go 的基本类型有
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名表示一个 Unicode 码点
float32 float64
complex64 complex128
*/

/*golang中，未明确初始化的变量会被赋予零值
零值是：
数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/
//=======类型转换===================
func function4() {
	var i float64
	i = float64(34)
	//Go 在不同类型的项之间赋值时需要显式转换
	//表达式 T(v) 将值 v 转换为类型 T
	fmt.Println(i + 0.1)

}

//================循环========================
func function5() {
	//golang只有for循环；golang的for循环没有括号；但大括号是必须的,不支持C++的逗号表达式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//在初始化语句里定义变量的话，作用域是for之内
	}
	//while循环用法,只是while关键字替换成了for
	//for后面什么都不写就是无限循环
	var i int = 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

//================条件跳转 if========================
func function6() {
	var i int = 10
	//表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
	if i == 10 {
		fmt.Println("yes")
	} else { //卧槽这个else不能写到下一行否则就编译不过，我真是日了
		fmt.Println("no")
	}
	if j := i + 11; j == 11 {
		fmt.Println("yes")
		//if 语句可以在条件表达式前执行一个简单的语句
	} else {
		fmt.Println(j) //在这里可以使用if里面定义的j
	}
}

//================条件跳转 switch========================
func function7() {
	//switch语句也可以在条件表达式前执行一个简单的语句
	switch os := runtime.GOOS; os {
	//switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。没执行到不会求值
	case "darwin": //case 无需为常量，且取值不必为整数，甚至可以是函数，也可以是表达式。下面会有示例
		fmt.Println("OS X.")
		fallthrough
		//每个分支结束之后并不需要写break；如果需要继续向下执行应当写fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	var os string = runtime.GOOS

	switch { //可以不写条件，不写就是true，因为case里可以是表达式，条件可以挪到case里面写
	case os == "windows": //其实就是更清晰的if-else联合
		fmt.Println("微软windows")
	case os == "linux":
		fmt.Println("linus的linux")
	}
}

func function8() {
	defer fmt.Println("world")
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会__立即求值__，但直到外层函数返回前该函数都不会被调用
	//被推迟的函数会按照后进先出的顺序调用(栈)
	fmt.Println("hello")
}

//==============指针================
//golang也有指针，类型是原类型！！前面加*
//golang有&,*运算，但是没有C的指针运算
func swap(x *int, y *int) {
	var tmp int
	tmp = *x
	var tmp2 *int = &tmp
	*x = *y
	*y = *tmp2
}

//结构体定义，type 结构体名 struct
type vertex struct {
	x int
	y int
}

func function9() {
	var v vertex = vertex{10, 2} //列表初始化
	fmt.Println(v.x)
	var v2 vertex = vertex{x: 5} //可以使用这种方法只初始化一部分
	fmt.Println(v2.x)
	//==========结构体的指针===============================
	//!!!!在golang里面没有->符号，而是还是.代替，例如下面
	var x *vertex = &vertex{-100, -101}
	fmt.Println(x.y) //->还是.
	v2.y = 10
	var v3 = v2
	v3.y = 101
	fmt.Println(v2, v3) //这说明结构体和C++一样，不是java里的引用

}

//==================数组================
//数组类型声明为[100]int诸如此类。显然，数组的大小在初始化时就确定了
//切片类型声明为[]int之类。其实切片就是浏览器
func function10() {
	var a [20]string
	a[0] = "soviet union"
	a[1] = "america"
	a[2] = "prc"
	var aslice []string = a[1:3]
	//标号包括第一个元素，但排除最后一个元素
	//切片切出来访问的时候下标从0开始而不是以前的下标
	for i := 0; i < 2; i++ {
		fmt.Println(aslice[i])
		aslice[i] += "."
	}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
		//更改切片的元素会修改其底层数组中对应的元素
	}

	//下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
	var tmp = []bool{true, true, false}
	fmt.Println(tmp)
	/*
		在进行切片时，你可以利用它的默认行为来忽略上下界。
		切片下界的默认值为 0，上界则是该切片的长度。
		对于数组 var a [10]int 来说，以下切片是等价的：
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	fmt.Println(len(tmp), cap(tmp)) //切片拥有 长度 和 容量。切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
}

func function11() {
	//使用make来创建一个已经自带底层数组的切片
	a := make([]int, 5)    //只有两个参数，传入的是长度5
	b := make([]int, 0, 5) //有三个参数，前者是长度，后者是容量
	//只要具有足够的容量，你就可以通过重新切片来扩展或缩减一个切片
	fmt.Println(a)
	b = b[0:4] //把一个长度为0的切片扩展为长度为4。如果超过容量会有runtime error
	fmt.Println(b, len(b), cap(b))
	//对于切片类型，如果只定义不初始化，那他的值是nil，长度和容量都是0

	//对于切片来说，容量不够不用虚，可以有append
	c := append(a, 10, 10, 10, 10)
	a = append(a, 11, 11, 11) //相当于扩容了
	fmt.Println(a, c)
}

func function12() {
	//for 循环的 range 形式可遍历切片或映射。
	//当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*可以将下标或值赋予 _ 来忽略它
		for _, value := range pow
		for i, _ := range pow
	若你只需要索引，忽略第二个变量即可。
		for i := range pow
	*/
}

//=========映射map==============
//类型是map[键类名]值类名。和数组一样，可以通过赋值初始化也可make
//golang似乎没有C++那种不用等号直接初始化的方法
func function13() {
	var m map[string]vertex = map[string]vertex{
		"usa": {1, 5},
	}
	m["ussr"] = vertex{1, 1} //增加/查找元素
	fmt.Println(m["usa"])
	delete(m, "ussr") //delete(m, key)删除元素
	fmt.Println(m)
	elem, ok := m["ussr"] //双赋值可以检测变量存不存在
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("苏联解体了")
	}
}

