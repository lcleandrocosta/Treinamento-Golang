/* Excesso de comentários não é uma boa prática, 
   mas como aprendizagem, estou deixando alguns conceitos, 
	 afim de consultas posteriores. */

/*Package, são formas de organizar o código Go.
  Link para as bibliotecas do go: https://pkg.go.dev/std
	O Go precisa desse package "main" para executar uma aplicação */
package main

//Declaração dos packages para uso na aplicação.
import "fmt"

/* Variável com escopo de package 
	 está disponível para qualquer 
	 arquivo ou função em escopo de package
*/
var titulo, subtitulo string
var WHITE_BACKGROUND, GREENCOLOR = "\u001B[47m", "\u001B[32m"
var YELLOW_COLOR, RESET_COLOR = "\u001B[33m", "\u001B[0m"

/* Função necessária para executar uma aplicação */
func main() {
	formatarCabecalho("Aprendendo Golang", "Funções")
}

func formatarCabecalho(titulo, subtitulo string) {
	fmt.Println("")
	fmt.Println(GREENCOLOR, "************", WHITE_BACKGROUND, titulo,RESET_COLOR, GREENCOLOR, "**************", RESET_COLOR)
	fmt.Println(YELLOW_COLOR,"# Assunto:", subtitulo, RESET_COLOR)
	fmt.Println("")
}