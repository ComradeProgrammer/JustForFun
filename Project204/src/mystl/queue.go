package mystl

import (
	"bytes"
	"fmt"
)

type Queue struct {
	v []interface{}
}

//any other methods should call this function at the beginning
//in order to make sure this struct is properly initiated
func (this *Queue) init() {
	if this.v == nil {
		this.v = make([]interface{}, 0, 1)
	}
}

func (this *Queue) Push(x ...interface{}) {
	this.init()
	for _, tmp := range x {
		this.v = append(this.v, tmp)
	}
}

func (this *Queue) Pop() {
	this.init()
	if len(this.v) == 0 {
		panic("front: size of vector is 0\n")
	}
	this.v = this.v[1:]
}

func (this *Queue) front() interface{} {
	this.init()
	if len(this.v) == 0 {
		panic("front: size of vector is 0\n")
	}
	return this.v[0]
}

func (this *Queue) back() interface{} {
	this.init()
	if len(this.v) == 0 {
		panic("back: size of vector is 0\n")
	}
	return this.v[len(this.v)-1]
}

func (this *Queue) empty() bool {
	this.init()
	return len(this.v) == 0
}

func (this *Queue) size() int {
	this.init()
	return len(this.v)
}

func (this *Queue) String() string {
	this.init()
	var buf *bytes.Buffer = new(bytes.Buffer)
	buf.WriteString("[")
	for index, obj := range this.v {
		if index != len(this.v)-1 {
			if tmp, ok := obj.(fmt.Stringer); ok {
				fmt.Fprintf(buf, "%s,", tmp)
			} else {
				fmt.Fprintf(buf, "%v,", obj)
			}

		} else {
			if tmp, ok := obj.(fmt.Stringer); ok {
				fmt.Fprintf(buf, "%s]", tmp)
			} else {
				fmt.Fprintf(buf, "%v]", obj)
			}
		}
	}
	return buf.String()
}
