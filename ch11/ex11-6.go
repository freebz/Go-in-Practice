// Listing 11.6  Types to examine

package main

import (
	"fmt"
	"reflect"

	"strings"
)

type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Nmae struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}
