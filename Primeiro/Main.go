package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Absoluto interface {
	Abs() float64
}

type Localizacao struct {
	Latitude, Longitude float64
}

var frutas = []string{"Mamão", "Abacaxi", "Uva", "Banana", "Morango", "Melão", "Maçã", "Pera", "Limão"}

func main() {
	var opcao int
	opcao = InputInt(`Escolha uma opção 
	1 - Soma
	2 - Tabuada
	3 - Range
	4 - Calcular area retangulo
	5 - Alterar pessoa por endereço
	6 - Pegar parte das frutas
	7 - Montar mapa
	8 - Exibir localização
	9 - Abrir arquivo de texto
	10 - Gerar erro
	11 - Assert Type
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
	case 6:
		PegarParteDasFrutas()
	case 7:
		MontarMapa()
	case 8:
		ExibirLocalizacao()
	case 9:
		AbrirArquivoTexto()
	case 10:
		GerarErro()
	case 11:
		AssertType()
	case 12:
		DoAbs()
	default:
		break
	}
}

func DoAbs() {
	var a Absoluto
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	//a = v //v � Vertex e n�o *Vertex e n�o implementa Absoluto
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AssertType() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //panic
	fmt.Println(f)
}

func GerarErro() {
	var err = GetErro()
	if err != nil {
		// log.Fatal(err)
		panic("Panicking!")
		// fmt.Println(err)
	}
	defer fmt.Println("Defer depois")
}

func GetErro() error {
	if 1 < 2 {
		return fmt.Errorf("Deu um erro: %d é menor que %d", 1, 2)
		// return errors.New("Erro: Um definitivamente é menor que 2")
	}
	return errors.New("Erro: Não deu erro?")
}

func AbrirArquivoTexto() {
	var caminho = InputStr("Informe o caminho do arquivo de texto")
	if strings.ToLower(filepath.Ext(caminho)) != ".txt" {
		fmt.Println("Formato de arquivo inválido")
		return
	}

	var arquivo, erroOpen = os.Open(caminho)
	if erroOpen != nil {
		log.Fatal(erroOpen)
		// fmt.Printf("Erro ao abrir o arquivo: %s\n", erroOpen)
		return
	}
	defer arquivo.Close() //defer é semelhante ao finnaly

	var conteudo, erroRead = io.ReadAll(arquivo)
	if erroRead != nil {
		log.Fatal(erroRead)
		// fmt.Printf("Erro ao ler o arquivo: %s\n", erroRead)
		return
	}

	fmt.Printf("Conteúdo do arquivo:\n%s", string(conteudo))
}

func ExibirLocalizacao() {
	var latitude, longitude = InputFloat("Informe a latitude"), InputFloat("Informe a longitude")
	fmt.Println(Localizacao{latitude, longitude})
}

func MontarMapa() {
	var timeZones = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
	}
	var zona = InputStr("Informe a zona desejada (UTC, EST, CST)")
	fmt.Println(timeZones[zona])
}

func PegarParteDasFrutas() {
	var qtdFrutas = InputInt("Informe a quantidade de frutas")
	for _, fruta := range frutas[0:qtdFrutas] {
		fmt.Printf("%s\n", fruta)
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
	var base, altura = InputInt("Informe a base"), InputInt("Informe a altura")
	var area, parametro int
	area, parametro = GetAreaRetangulo(base, altura)
	fmt.Printf("Ýrea: %d\nParametro: %d", area, parametro)
}

func GetAreaRetangulo(base, altura int) (area, parametro int) {
	parametro = 2 * (altura + base)
	area = altura * base
	return
}

func PercorrerVetor() {
	var cpfs = [4]string{"12347203842", "45348264821", "96402047222", "98660638393"}
	var nomes = [4]string{"Alberto", "Marcelo", "Roberto", "Patricia"}
	for i, cpf := range cpfs {
		fmt.Printf("cpf: %s nome: %s\n", cpf, nomes[i])
	}
}

func Tabuada() {
	var valor, limite = InputInt("Informe o valor da tabuada"), InputInt("Informe o limite")
	for i := 1; i <= limite; i++ {
		fmt.Printf("%d X %d = %d\n", valor, i, i*valor)
	}
}

func Soma() {
	var valor1, valor2 = InputFloat("Informe um valor"), InputFloat("Informe outro valor")
	fmt.Printf("%.2f + %.2f é igual a %.2f", valor1, valor2, valor1+valor2)
}

func InputStr(msg string) string {
	var result string
	fmt.Println(msg)
	fmt.Scanln(&result)
	return result
}

func InputFloat(msg string) float64 {
	var result, err = strconv.ParseFloat(strings.Replace(InputStr(msg), ",", ".", -1), 64)

	if err != nil {
		fmt.Println("Valor inválido")
		return InputFloat(msg)
	} else {
		return result
	}
}

func InputInt(msg string) int {
	var result, err = strconv.Atoi(strings.Replace(InputStr(msg), ",", ".", -1))

	if err != nil {
		fmt.Println("Valor inválido")
		return InputInt(msg)
	} else {
		return result
	}
}
