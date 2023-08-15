package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var opcao int
	opcao = InputInt(`Escolha uma opção 
	1 - Soma
	2 - Tabuada
	3 - Range
	4 - Calcular area retangulo
	5 - Alterar pessoa por endereço
	0 - Sair`)

	switch opcao {
	case 1:
		Soma()
	case 2:
		Tabuada()
	case 3:
		PercorrerVetor()
	case 4:
		AreaRetangulo()
	case 5:
		AlterarPessoaPorEndereco()
	default:
		break
	}
}

func AlterarPessoaPorEndereco() {
	var nome = "Amanda"
	var idade = 25
	fmt.Printf("Antes:\n  Nome: %s\n  Idade: %d\n", nome, idade)
	AlterarPessoa(&nome, &idade)
	fmt.Printf("Depois:\n  Nome: %s\n  Idade: %d\n", nome, idade)
}

func AlterarPessoa(nome *string, idade *int) {
	*nome = *nome + " " + InputStr("Informe o sobrenome")
	*idade += 5
}

func AreaRetangulo() {
	var base, altura, area, parametro int
	base = InputInt("Informe a base")
	altura = InputInt("Informe a altura")
	area, parametro = GetAreaRetangulo(base, altura)
	fmt.Printf("Área: %d\nParametro: %d", area, parametro)
}

func GetAreaRetangulo(base, altura int) (area, parametro int) {
	parametro = 2 * (altura + base)
	area = altura * base
	return
}

func PercorrerVetor() {
	var cpfs = []string{"12347203842", "45348264821", "96402047222", "98660638393"}
	var nomes = []string{"Alberto", "Marcelo", "Roberto", "Patricia"}
	for i, cpf := range cpfs {
		fmt.Printf("cpf: %s nome: %s\n", cpf, nomes[i])
	}
}

func Tabuada() {
	var valor, limite int
	valor = InputInt("Informe o valor da tabuada")
	limite = InputInt("Informe o limite")
	for i := 1; i <= limite; i++ {
		fmt.Printf("%d X %d = %d\n", valor, i, i*valor)
	}
}

func Soma() {
	var valor1, valor2 float64
	valor1 = InputFloat("Informe um valor")
	valor2 = InputFloat("Informe outro valor")
	fmt.Printf("%.2f + %.2f é igual a %.2f", valor1, valor2, valor1+valor2)
}

func InputStr(msg string) string {
	var result string
	fmt.Println(msg)
	fmt.Scanln(&result)
	return result
}

func InputFloat(msg string) float64 {
	var erro error
	var result float64
	result, erro = strconv.ParseFloat(strings.Replace(InputStr(msg), ",", ".", -1), 64)

	if erro != nil {
		fmt.Println("Valor inválido")
		return InputFloat(msg)
	} else {
		return result
	}
}

func InputInt(msg string) int {
	var erro error
	var result int
	result, erro = strconv.Atoi(strings.Replace(InputStr(msg), ",", ".", -1))

	if erro != nil {
		fmt.Println("Valor inválido")
		return InputInt(msg)
	} else {
		return result
	}
}
