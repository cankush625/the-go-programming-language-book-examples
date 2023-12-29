package main

import "fmt"

type Ptr struct {
	X, Y float64
}

func (p Ptr) Add(q Ptr) Ptr {
	return Ptr{p.X + q.X, p.Y + q.Y}
}

func (p Ptr) Sub(q Ptr) Ptr {
	return Ptr{p.X - q.X, p.Y - q.Y}
}

type Path_ []Ptr

func (path Path_) TranslateBy(offset Ptr, add bool) {
	var op func(p, q Ptr) Ptr
	if add {
		op = Ptr.Add // method expression
	} else {
		op = Ptr.Sub
	}
	for i := range path {
		// When calling a method expression, the first parameter
		// of the method expression always takes the place of
		// receiver.
		path[i] = op(path[i], offset)
	}
}

func main() {
	path := Path_{{2, 3}, {4, 5}}
	offset_ptr := Ptr{2, 2}
	path.TranslateBy(offset_ptr, true)
	fmt.Println(path) // "[{4 5} {6 7}]"
}
