/* Excesso de comentários não é uma boa prática, 
   mas como aprendizagem, estou deixando alguns conceitos, 
	 afim de consultas posteriores. */

/*Package, são formas de organizar o código Go.
  Link para as bibliotecas do go: https://pkg.go.dev/std
	O Go precisa desse package "main" para executar uma aplicação */
package main

//Declaração dos packages para uso na aplicação.
import (
	"fmt"
	"time"
)

/* Variável com escopo de package 
	 está disponível para qualquer 
	 arquivo ou função em escopo de package
*/
var titulo string
var YELLOW_BACKGROUND = "\u001B[43m"
var BLUE, GREENCOLOR, WHITE_COLOR, RESET_COLOR = "\u001B[34m", "\u001B[32m", "\u001B[37m", "\u001B[0m"

/* Função necessária para executar uma aplicação */
func main() {
	formatarCabecalho("Aprendendo a linguagem:")
	apresentarLogo()
}

func formatarCabecalho(titulo string) {
	fmt.Println("")
	fmt.Println(GREENCOLOR, "************", YELLOW_BACKGROUND, WHITE_COLOR, titulo, RESET_COLOR, 
	GREENCOLOR, "**************", RESET_COLOR)
}

func apresentarLogo() {
	var bi = "\\" //Barra invertida

	fmt.Println(BLUE)
	fmt.Println("  _________          __ ")
	timer()
	fmt.Println(" /  ______/  ____   |  |   ____      ____     ____")
	timer()
	fmt.Printf("/   %s  __   /    %s  |  |  %s__   %s   /    %s   /  ___%s \n", bi,bi, bi, bi, bi, bi)
	timer()
	fmt.Printf("%s    %s___%s (  <_> ) |__|__ / __  %s |  |   %s /  / _/ >\n" ,bi, bi, bi, bi, bi)
	timer()
	fmt.Printf(" %s______  / %s____/  |______(___   /___|   /%s____   /\n", bi,bi,bi)
	timer()
	fmt.Printf("        %s/                      %s/      %s//_______/\n", bi,bi,bi)
	fmt.Println()
	fmt.Println(GREENCOLOR, "********************************************************", RESET_COLOR)
}

func timer() {
	newtimer := time.NewTimer(1 * time.Second)
	<-newtimer.C
}