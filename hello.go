package main

import (
	"fmt"
	"reflect"
)

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

	simpleName := "Douglas"
	systemVersion := 1.1

	fmt.Println("Olá Sr.", simpleName)
	fmt.Println("A versão do sistema é:", systemVersion)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	var inputCommand int
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
	fmt.Scan(&inputCommand)

	fmt.Println("O comando escolhido foi:", inputCommand)
	fmt.Println("O endereço da minha variável inputCommand é:", &inputCommand)
}
