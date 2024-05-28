package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Inicializando o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Definindo o número de lotes e o número de produtos por lote
	numLotes := 5
	numProdutos := 4

	// Criando uma matriz para armazenar os resultados da verificação de qualidade
	var matriz [5][4]bool

	// Preenchendo a matriz com resultados aleatórios de verificação de qualidade
	for i := 0; i < numLotes; i++ {
		for j := 0; j < numProdutos; j++ {
			// Gerando um resultado aleatório de verificação de qualidade
			// Aqui, 'true' significa que o produto passou na verificação de qualidade e 'false' significa que falhou
			matriz[i][j] = rand.Intn(2) == 0
		}
	}

	// Abrindo o arquivo HTML para escrita
	f, err := os.Create("resultados.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Escrevendo o início do arquivo HTML
	f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<style>\nbody {\n\ttext-align: center;\n\tfont-family: Arial, sans-serif;\n}\n.passed {\n\tbackground-color: green;\n\tcolor: white;\n\tmargin: 1em;\n\tpadding: 1em;\n}\n.failed {\n\tbackground-color: red;\n\tcolor: white;\n\tmargin: 1em;\n\tpadding: 1em;\n}\n.container {\n\tdisplay: flex;\n\tjustify-content: center;\n}\n.left {\n\tmargin-right: auto;\n}\n.right {\n\tmargin-left: auto;\n}\n</style>\n</head>\n<body>\n<h1>Resultados da Verificação de Qualidade</h1>\n<div class=\"container\">\n<div class=\"left\">\n<h2>Passou na Verificação de Qualidade</h2>\n")

	// Escrevendo os produtos que passaram na verificação de qualidade
	passou := 0
	for i := 0; i < numLotes; i++ {
		for j := 0; j < numProdutos; j++ {
			if matriz[i][j] {
				f.WriteString(fmt.Sprintf("<div class=\"passed\">Lote %d, Produto %d: Passou na verificação de qualidade</div>\n", i+1, j+1))
				passou++
			}
		}
	}

	// Escrevendo o meio do arquivo HTML
	f.WriteString("</div>\n<div class=\"right\">\n<h2>Falhou na Verificação de Qualidade</h2>\n")

	// Escrevendo os produtos que falharam na verificação de qualidade
	falhou := 0
	for i := 0; i < numLotes; i++ {
		for j := 0; j < numProdutos; j++ {
			if !matriz[i][j] {
				f.WriteString(fmt.Sprintf("<div class=\"failed\">Lote %d, Produto %d: Falhou na verificação de qualidade</div>\n", i+1, j+1))
				falhou++
			}
		}
	}

	// Escrevendo o final do arquivo HTML
	f.WriteString(fmt.Sprintf("</div>\n</div>\n<h2>Total de Verificações</h2>\n<p>Passou na Verificação de Qualidade: %d</p>\n<p>Falhou na Verificação de Qualidade: %d</p>\n</body>\n</html>", passou, falhou))
}
