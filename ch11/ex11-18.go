// Listing 11.18  Simple queue

package main

type MyTypeQueue struct {
	q []ByType
}

func NewMyTypeQueue() *MyTypeQueue {
	return &MyTypeQueue{
		q: []ByType{},
	}
}

func (o *MyTypeQueue) Insert(v MyType) {
	o.q = append(o.q, v)
}

func (o *MyTypeQueue) Remove() MyType {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
