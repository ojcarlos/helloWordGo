package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	exibeIntroducao()
	lerArquivo()

	for {
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
	sites := []string{"https://httpbin.org/status/404",
		"https://www.alura.com.br", "https://www.caelum.com.br"}
	// site com URL inexistente
	for j := 0; j < 5; j++ {
		fmt.Println("INICIANDO BATERIA DE TESTES ", j)
		for i, site := range sites {
			fmt.Println("TESTE SITE ", i, ":", site)
			testarSite(site)
		}
		time.Sleep(5 * time.Second)
	}

}
func testarSite(site string) {
	fmt.Println(site)
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}

func lerArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	fmt.Println(sites)
	return sites
}
