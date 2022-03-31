package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	for {
		var nome string
		var check bool

		fmt.Println("Digite o nome do item: ")
		fmt.Scan(&nome)

		if nome == "sair" {
			imprimeLista()
			return
		}

		fmt.Println("Deu certo?")
		fmt.Scan(&check)

		fmt.Println("Item registrado!")
		registrarLista(nome, check)
	}
}

func registrarLista(item string, status bool) {
	arquivo, erro := os.OpenFile("listaCompras.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	data := time.Now().Format("02/01/2006 15:04:05")

	if erro != nil {
		fmt.Println("Ocorreu um erro: ", erro)
	}

	arquivo.WriteString("registro as " + data + " ")
	arquivo.WriteString(item + " deu " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLista() {
	arquivo, erro := ioutil.ReadFile("listaCompras.txt")

	if erro != nil {
		fmt.Println("Ocorreu um erro: ", erro)
	}

	fmt.Println(string(arquivo))
}
