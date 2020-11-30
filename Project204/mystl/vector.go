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

//return a iterator of the vector
//if the vector is modified,then the use of iterator would be UNSAFE
func (this *Vector)Iterator(pos int)Iterator{
	var x=pos
	var ptr *[]interface{}=&(this.v)
	return func(op Operation)interface{}{
		switch(op){
		case NEXT:
			x++
		case BACK:
			x--
		}
		return (*ptr)[x]
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

//todo implement
/*iterator insert(iterator it,const T& x)*/
func (this *Vector) PushBackSingleByIterator(){

}

//todo implement
func(this *Vector)PushBackMultipulByIterator(){

}

func (this *Vector)String()string{
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







