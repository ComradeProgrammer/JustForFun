package mystl

import (
	"bytes"
	"fmt"
)

type Vector struct{
	v []interface{}
}


//any other methods should call this function at the beginning
//in order to make sure this struct is properly initiated
func (this *Vector)init(){
	if this.v==nil{
		this.v=make([]interface{},0,1)
	}
}


//=====================INSERT=================================

/*counterpart for void push_back(const T& x) */
func (this *Vector)PushBack(x ...interface{}){
	this.init()
	for _,tmp:=range x{
		this.v=append(this.v,tmp)
	}
}

func(this *Vector)Insert(pos int,x interface{}){
	this.init()
	if pos<0||pos>len(this.v){
		panic(fmt.Sprintf("at:index out of bounds:%d\n",pos))
	}
	this.v=append(this.v,x)
	for i:=len(this.v)-2;i>=pos;i--{
		this.v[i+1]=this.v[i]
	}
	this.v[pos]=x
}



func (this *Vector)String()string{
	this.init()
	var buf *bytes.Buffer=new(bytes.Buffer)
	buf.WriteString("[")
	for index,obj:=range this.v{
		if index!=len(this.v)-1{
			if tmp,ok:=obj.(fmt.Stringer);ok{
				fmt.Fprintf(buf,"%s,",tmp)
			} else{
				fmt.Fprintf(buf,"%v,",obj)
			}

		}else{
			if tmp,ok:=obj.(fmt.Stringer);ok{
				fmt.Fprintf(buf,"%s]",tmp)
			} else{
				fmt.Fprintf(buf,"%v]",obj)
			}
		}
	}
	return buf.String()
}

//=====================QUERY=================================
func (this *Vector)at(pos int)interface{}{
	this.init()
	if pos<0||pos>=len(this.v){
		panic(fmt.Sprintf("at:index out of bounds:%d\n",pos))
	}
	return this.v[pos]
}

func (this *Vector)front()interface{}{
	this.init()
	if len(this.v)==0{
		panic("front: size of vector is 0\n")
	}
	return this.v[0]
}

func (this *Vector)back()interface{}{
	this.init()
	if len(this.v)==0{
		panic("back: size of vector is 0\n")
	}
	return this.v[len(this.v)-1]
}

func (this *Vector)empty()bool{
	this.init()
	return len(this.v)==0
}

func (this *Vector)size()int{
	this.init()
	return len(this.v)
}
//=========================update====================

func (this *Vector)set(pos int,x interface{}){
	this.init()
	if pos<0||pos>=len(this.v){
		panic(fmt.Sprintf("at:index out of bounds:%d\n",pos))
	}
	this.v[pos]=x
}

//========================delete=====================
func (this *Vector)erase(pos int){
	this.init()
	if pos<0||pos>=len(this.v){
		panic(fmt.Sprintf("at:index out of bounds:%d\n",pos))
	}
	if(pos==0){
		this.v=this.v[1:len(this.v)]
	}else if(pos==len(this.v)-1){
		this.v=this.v[:len(this.v)-1]
	}else{
		copy(this.v[pos:],this.v[pos+1:])
		this.v= this.v[0:len(this.v)-1]
	}

}

func (this *Vector)PopBack(){
	this.init()
	if len(this.v)==0{
		panic("back: size of vector is 0\n")
	}
	this.v=this.v[:len(this.v)-1]
}

func(this *Vector) clear(){
	this.init()
	this.v=this.v[0:0]
}











