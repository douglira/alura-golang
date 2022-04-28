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
}
