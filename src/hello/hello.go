package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	fmt.Println("")
	exibeIntroducao()
	for {
		exibeMenu()
		fmt.Println("")
		comando := leComando()

		switch comando {   //Não precisa do break, Go para o case automaticamente
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}



func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu(){
	fmt.Println("1 - INICIAR MONITORAMENTO")
	fmt.Println("2 - EXIBIR LOGS")
	fmt.Println("0 - SAIR DO PROGRAMA")
}

func leComando() int{
	var comandoLido int            //inicia com zero
	//fmt.Scanf("%d", &comando)  //scanf tem que indicar o modificador, aponta para o endereço do comando
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")

	//var sites [4]string  //array tem tamanho fixo, melhor trabalhar com slices
	
	// sites := [] string {
	// 	"https://httpstat.us/Random/200,404",
	// 	"https://www.alura.com.br",
	// 	"https://www.google.com.br",
	// }
	sites := leArquivo()


	for i:=0; i<monitoramentos; i++ {
		for pos := range sites{
			testaSite(sites[pos])
		}	
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}


func testaSite(site string){
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

		if response.StatusCode == 200 {
			fmt.Println("Site: ", site, "foi carregado com sucesso")
		} else {
			fmt.Println("Site: ", site, "está com problemas. Status Code: ", response.StatusCode)
		}
} 


func leArquivo() []string {  //retorna uma slice de strings
	
	var sites []string
	
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(arquivo)
	return sites
}