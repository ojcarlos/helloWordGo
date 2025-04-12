package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Carlos"
	versao := 1.1
	idade := 28
	fmt.Println("hello word,", nome)
	fmt.Println("este programa está na versão", versao)
	fmt.Println("sua idade é", idade)
	fmt.Println("o tipo da variável nome é:", reflect.TypeOf(nome))
}
