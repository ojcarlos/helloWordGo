package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	exibirMenu()
	comando := lerComando()

	if comando == 1 {
		iniciarMonitoramento()

	} else if comando == 2 {
		fmt.Println("exibindo log...")

	} else if comando == 0 {
		fmt.Println("Fechando programa")
		os.Exit(0)

	} else {
		fmt.Println("Não existe tal comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Carlos"
	versao := 1.1
	fmt.Println("Olá, ", nome)
	fmt.Println("este programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("o Comando escolhido foi: ", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
