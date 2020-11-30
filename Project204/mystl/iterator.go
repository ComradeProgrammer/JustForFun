package mystl
type Operation int
const(
	THIS Operation=iota
	NEXT
	BACK
)

//implement iterator interaface with closure
type Iterator func(_ Operation)interface{}

func (this Iterator)This()interface{}{
	return this(THIS)
}

func (this Iterator)Next()interface{}{
	return this(NEXT)
}

func (this Iterator)Back()interface{}{
	return this(BACK)
}
