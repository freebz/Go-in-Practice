// Listing 11.5  Determine whether a type implements an interface

package main

import (
	"fmt"
	"io"
	"reflect"
)

type Name struct {
	First, Last string
}

func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {
	n := &Name{First: "Inogo", Last: "Montoya"}

	stringer :=
		(*fmt.Stringer)(nil)
	implements(n, stringer)

	writer := (*io.Writer)(nil)
	implements(n, writer)
}

func implements(concrete interface{}, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	t := v.Type()

	if t.Implements(iface) {
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("%T is not a %s\n", concrete, iface.Name())
	return false
}
