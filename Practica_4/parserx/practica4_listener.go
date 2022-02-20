// Code generated from Practica4.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Practica4

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Practica4Listener is a complete listener for a parse tree produced by Practica4Parser.
type Practica4Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLista is called when entering the lista production.
	EnterLista(c *ListaContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLista is called when exiting the lista production.
	ExitLista(c *ListaContext)
}
