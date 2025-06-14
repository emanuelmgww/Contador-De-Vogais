package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("----- Contador de Vogais -----")

	reader := bufio.NewReader(os.Stdin)

	var stringUsuario string

	for {
		fmt.Print("\nDigite qualquer coisa: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("\nErro ao ler a entrada! Tente novamente")
			continue
		} 

		stringUsuario = strings.TrimSpace(input)

		if stringUsuario == "" {
			fmt.Println("\nEntrada inválida! Tente novamente")
			continue
		}

		if !apenasLetras(stringUsuario) { // Diferente de true
			fmt.Println("\nEntrada inválida! Digite apenas letras.")
			continue
		}

		break
	}

	fmt.Println("\nQuantidade de vogais na frase digitada:", somaVogais(stringUsuario))

}

func apenasLetras(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func somaVogais(stringUsuario string) int {
	stringUsuario = strings.ToLower(stringUsuario)

	var soma int
	for _, r := range stringUsuario {
		if strings.ContainsRune("aeiou", r) {
			soma++
		}
	}
	return soma
}
