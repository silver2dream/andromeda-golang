package main

import (
	"fmt"
	"github.com/silver2dream/andromeda/m31"
	_ "github.com/silver2dream/andromeda/m31"
)

func main() {
	Example1()

	Example2()
}

func Example1() {
	n1 := m31.NewFInt(1)
	n2 := m31.NewFInt(0.8)

	//fmt.Println(n1.LeftShift(1).ToString())
	fmt.Println(n1.Add(n2).ToString())
	fmt.Println(n1.Neg().ToString())
}

func Example2() {
	n1 := m31.NewFInt(1)
	n2 := m31.NewFInt(1.5)
	fmt.Println(n1.Mul(n2).ToString())

	n3 := m31.NewFInt(2)
	n4 := m31.NewFInt(0.5)

	fmt.Println(n3.Div(n4).ToString())
}
