package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	listaDeCompras := leListaDoArquivos()

	fmt.Println(listaDeCompras)
}

func leListaDoArquivos() []string {
	var listaCompras []string
	arquivo, err := os.Open("lista.txt")

	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		listaCompras = append(listaCompras, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return listaCompras
}
