package main

import (
	"fmt"
)

func main() {
	nome := "Carlos"
	versao := 1.1
	fmt.Println("Olá, ", nome)
	fmt.Println("este programa está na versão", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("o Comando escolhido foi: ", comando)
	fmt.Println("O endereço do comando foi: ", &comando)
}
