package main

import (
	"context"
	"fmt"
)

func main() {
	Reachable()
}

func init() {
	s := myString{Value: "hello"}
	s.Reachable()
}

func Reachable() {
	fmt.Println("reachable")
}

func Unreachable() {
	fmt.Println("unreachable")
}

func UnreachableGeneric[T comparable](a T) {
	fmt.Println(a)
}

func ReachableByTest() {
	fmt.Println("reachableByTest")
}

func UnusedImportStatementRemoval() context.Context {
	return context.TODO() // Test unused import statement removal
}

var _ fmt.Stringer = myString{}

type myString struct {
	Value string
}

// NOTE: This function is unreachable but it is neccessary to implement the fmt.Stringer interface
func (s myString) String() string {
	return s.Value
}

func (s myString) Reachable() {
	return
}

func (s *myString) Unreachable() {
	return
}

type myGenericStruct[T any] struct {
	value T
}

func (s myGenericStruct[T]) Value() T {
	return s.value
}

func (s *myGenericStruct[T]) Unreachable() {
	return
}
