 
package main

import (
	parser "Practica4/parserx"
	"bufio"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Syntax struct {
	*parser.BasePractica4Listener
}

func NewSyntax() *Syntax {
	return new(Syntax)
}
func interpretar(in string) {
	is := antlr.NewInputStream(in)

	//crear analizador Léxico
	lexer := parser.NewPractica4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//crear analizador sintáctico
	p := parser.NewPractica4Parser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewSyntax(), p.Start())

}
func main() {
	file, err := os.Open("entrada.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	text := ""

	for fileScanner.Scan() {
		text += fileScanner.Text()
	}

	interpretar(text)
}
