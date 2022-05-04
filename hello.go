package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const secondInterval = 2
const monitoringTimes = 5

func main() {
	/*
		Função com letra maiúsculina - Convenção da linguagem para
		informar que é uma função de um pacote/modulo externo
	*/
	fmt.Println("Olá Mundo, Douglas!")

	var name string = "Douglas Lira"
	var version float32 = 1.1
	var age int = 26

	/*
		- Variáveis declaradas sem atribuição de valor assumem
			um valor padrão: 0 para inteiros, 0.0 para float, string
			vazia para string...
		- Em Go, ocorre erro de compilação para variáveis declaradas
			na qual não estão sendo utilizadas
	*/
	// var height float32

	fmt.Println("Olá, Sr.", name, ", sua idade é", age)
	fmt.Println("Este programa está na versão", version)

	/*
		- Inferência de tipo - Ao atribuir um valor explícito, não
			se torna obrigatório informar o tipo da variável
		- Operador de declaração curta de variáveis (:=)
			declara e atribui um valor à variável (short syntax)
	*/
	var quantity = 21
	height := 1.68
	fmt.Println("O tipo da variável é:", reflect.TypeOf(quantity))
	fmt.Println("O tipo da variável é:", reflect.TypeOf(height))

	/*
		Capturando input do terminal
	*/

	// simpleName := "Douglas"
	// systemVersion := 1.1

	// fmt.Println("Olá Sr.", simpleName)
	// fmt.Println("A versão do sistema é:", systemVersion)

	// fmt.Println("1 - Iniciar Monitoramento")
	// fmt.Println("2 - Exibir Logs")
	// fmt.Println("0 - Sair")

	// var inputCommand int
	/*
		O "&" indica o endereço da variável na qual está atrelado.
		Neste caso, está indicando o endereço de memória da variável
		"inputCommand" para que o valor inserido no terminal seja atribuído
		à variável.
		A função "Scanf" espera como primeiro argumento, a especificação
		do tipo do input para a variável
	*/
	// fmt.Scanf("%d", &inputCommand)

	/*
		Na função "Scan" não necessita especificar o tipo de input.
		Caso seja informado um input diferente do tipo da variável,
		neste caso "inputCommand", este valor não será interpretado
		e a variável irá assumir seu valor padrão, no caso de int será 0
	*/
	// fmt.Scan(&inputCommand)

	// fmt.Println("O comando escolhido foi:", inputCommand)
	// fmt.Println("O endereço da minha variável inputCommand é:", &inputCommand)

	/**
	No IF é obrigatório a utilização de uma expressão na qual retorne um boolean
	Não é possível validar tipos diferentes de boolean na expressão
	*/
	// if inputCommand == 1 {
	// 	fmt.Println("Exibindo logs...")
	// } else if inputCommand == 2 {
	// 	fmt.Println("IF/ELSE> Monitorando...")
	// } else if inputCommand == 0 {
	// 	fmt.Println("IF/ELSE> Saindo do programa...")
	// } else {
	// 	fmt.Println("IF/ELSE> Não conheço este comando!")
	// }

	// switch inputCommand {
	// case 1:
	// 	fmt.Println("SWITCH> Exibindo logs...")
	// case 2:
	// 	fmt.Println("SWITCH> Monitorando...")
	// case 0:
	// 	fmt.Println("SWITCH> Saindo do programa...")
	// default:
	// 	fmt.Println("SWITCH> Não conheço este comando!")
	// }

	/**
	Trabalhando com funções
	*/
	displayIntro()
	displayMenu()
	userChoiceCommand := readCommand()

	switch userChoiceCommand {
	case 1:
		initializeMonitoring()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Ocorreu um erro inesperado")
		os.Exit(-1)
	}
}

func displayIntro() {
	fmt.Println("###################################")
	fmt.Println("###### Seja muito bem-vindo! ######")
	fmt.Println("###################################")
}

func displayMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func readCommand() int {
	var userChoiceCommand int
	fmt.Scan(&userChoiceCommand)

	return userChoiceCommand
}

func initializeMonitoring() {

	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	for i := 0; i < monitoringTimes; i++ {
		for index, site := range sites {
			message, _ := requestMonitoring(site)
			fmt.Println(index+1, "-", message)
		}
		time.Sleep(time.Duration(secondInterval) * time.Second)
	}
	// for i := 0; i < len(sites); i++ {
	// 	site := sites[i]
	// 	message, _ := requestMonitoring(site)
	// 	fmt.Println(message)
	// 	time.Sleep(time.Duration(secondInterval) * time.Second)
	// }
}

func requestMonitoring(site string) (string, *http.Response) {
	/*
		Quando uma função retorna múltiplos valores, podemos ignorar
		os valores utilizando o operador underline (_)
	*/
	var message string
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		message = fmt.Sprintf("Site: %s foi carregado com sucesso!!", site)
	} else {
		message = fmt.Sprintf("Site: %s está com problema. Status Code: %s", site, response.Status)
	}

	return message, response
}
