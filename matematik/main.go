package main

import (
	"fmt"
	"github.com/go-turk/konya/matematik/matematikLib"
)

func main() {
	mat := matematikLib.NewMatematik()
	fmt.Println(mat.Topla(10, 2))
	fmt.Println(mat.Cikar(10, 2))
	fmt.Println(mat.Carp(10, 2))
	fmt.Println(mat.Bol(10, 2))

}
