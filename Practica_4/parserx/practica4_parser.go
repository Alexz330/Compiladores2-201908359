// Code generated from Practica4.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Practica4

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"Practica4/Node"

	"math"
	"strings"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 31, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 26,
	10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 2, 2, 29, 2,
	6, 3, 2, 2, 2, 4, 10, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8,
	9, 8, 2, 1, 2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1, 2, 11, 12, 7, 3, 2, 2,
	12, 13, 7, 6, 2, 2, 13, 14, 7, 5, 2, 2, 14, 15, 7, 6, 2, 2, 15, 16, 7,
	4, 2, 2, 16, 17, 8, 3, 1, 2, 17, 27, 3, 2, 2, 2, 18, 19, 12, 4, 2, 2, 19,
	20, 7, 3, 2, 2, 20, 21, 7, 6, 2, 2, 21, 22, 7, 5, 2, 2, 22, 23, 7, 6, 2,
	2, 23, 24, 7, 4, 2, 2, 24, 26, 8, 3, 1, 2, 25, 18, 3, 2, 2, 2, 26, 29,
	3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2,
	29, 27, 3, 2, 2, 2, 3, 27,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "PARA", "PARC", "COMA", "NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "lista",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Practica4Parser struct {
	*antlr.BaseParser
}

func NewPractica4Parser(input antlr.TokenStream) *Practica4Parser {
	this := new(Practica4Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Practica4.g4"

	return this
}

func mostrarResultado(out string) {

	cadena := strings.Split(out, " ")

	var arreglito []float64
	for i := 0; i < len(cadena); i++ {

		flotante, _ := strconv.ParseFloat(cadena[i], 64)

		arreglito = append(arreglito, flotante)
	}
	var res float64 = 0
	for i := 0; i < len(arreglito)-1; i++ {
		if i+2 >= (len(arreglito) - 1) {
			break
		}

		op1 := math.Pow(arreglito[i+2]-arreglito[i], 2)
		op2 := math.Pow(arreglito[i+3]-arreglito[i+1], 2)
		res = res + math.Sqrt(op1+op2)

		i = i + 1
		if i >= (len(arreglito) - 1) {
			break
		}
	}

	fmt.Println("Distancia: ", fmt.Sprintf("%.2f", res))

}

// Practica4Parser tokens.
const (
	Practica4ParserEOF        = antlr.TokenEOF
	Practica4ParserPARA       = 1
	Practica4ParserPARC       = 2
	Practica4ParserCOMA       = 3
	Practica4ParserNUMBER     = 4
	Practica4ParserWHITESPACE = 5
)

// Practica4Parser rules.
const (
	Practica4ParserRULE_start = 0
	Practica4ParserRULE_lista = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUn returns the un rule contexts.
	GetUn() IListaContext

	// SetUn sets the un rule contexts.
	SetUn(IListaContext)

	// GetN returns the n attribute.
	GetN() node.Nodo

	// SetN sets the n attribute.
	SetN(node.Nodo)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      node.Nodo
	un     IListaContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica4ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetUn() IListaContext { return s.un }

func (s *StartContext) SetUn(v IListaContext) { s.un = v }

func (s *StartContext) GetN() node.Nodo { return s.n }

func (s *StartContext) SetN(v node.Nodo) { s.n = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(Practica4ParserEOF, 0)
}

func (s *StartContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Practica4Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Practica4ParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)

		var _x = p.lista(0)

		localctx.(*StartContext).un = _x
	}
	{
		p.SetState(5)
		p.Match(Practica4ParserEOF)
	}

	mostrarResultado(localctx.(*StartContext).GetUn().GetN().A + localctx.(*StartContext).GetUn().GetN().B)

	return localctx
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// GetNum2 returns the num2 token.
	GetNum2() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// SetNum2 sets the num2 token.
	SetNum2(antlr.Token)

	// GetNa returns the na rule contexts.
	GetNa() IListaContext

	// SetNa sets the na rule contexts.
	SetNa(IListaContext)

	// GetN returns the n attribute.
	GetN() node.Nodo

	// SetN sets the n attribute.
	SetN(node.Nodo)

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      node.Nodo
	na     IListaContext
	num    antlr.Token
	num2   antlr.Token
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica4ParserRULE_lista
	return p
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) GetNum() antlr.Token { return s.num }

func (s *ListaContext) GetNum2() antlr.Token { return s.num2 }

func (s *ListaContext) SetNum(v antlr.Token) { s.num = v }

func (s *ListaContext) SetNum2(v antlr.Token) { s.num2 = v }

func (s *ListaContext) GetNa() IListaContext { return s.na }

func (s *ListaContext) SetNa(v IListaContext) { s.na = v }

func (s *ListaContext) GetN() node.Nodo { return s.n }

func (s *ListaContext) SetN(v node.Nodo) { s.n = v }

func (s *ListaContext) PARA() antlr.TerminalNode {
	return s.GetToken(Practica4ParserPARA, 0)
}

func (s *ListaContext) COMA() antlr.TerminalNode {
	return s.GetToken(Practica4ParserCOMA, 0)
}

func (s *ListaContext) PARC() antlr.TerminalNode {
	return s.GetToken(Practica4ParserPARC, 0)
}

func (s *ListaContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(Practica4ParserNUMBER)
}

func (s *ListaContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(Practica4ParserNUMBER, i)
}

func (s *ListaContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterLista(s)
	}
}

func (s *ListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitLista(s)
	}
}

func (p *Practica4Parser) Lista() (localctx IListaContext) {
	return p.lista(0)
}

func (p *Practica4Parser) lista(_p int) (localctx IListaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, Practica4ParserRULE_lista, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(9)
		p.Match(Practica4ParserPARA)
	}
	{
		p.SetState(10)

		var _m = p.Match(Practica4ParserNUMBER)

		localctx.(*ListaContext).num = _m
	}
	{
		p.SetState(11)
		p.Match(Practica4ParserCOMA)
	}
	{
		p.SetState(12)

		var _m = p.Match(Practica4ParserNUMBER)

		localctx.(*ListaContext).num2 = _m
	}
	{
		p.SetState(13)
		p.Match(Practica4ParserPARC)
	}

	x := (func() string {
		if localctx.(*ListaContext).GetNum() == nil {
			return ""
		} else {
			return localctx.(*ListaContext).GetNum().GetText()
		}
	}()) + " " + (func() string {
		if localctx.(*ListaContext).GetNum2() == nil {
			return ""
		} else {
			return localctx.(*ListaContext).GetNum2().GetText()
		}
	}())
	y := ""
	localctx.(*ListaContext).n = node.NewNodo(x, y)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaContext(p, _parentctx, _parentState)
			localctx.(*ListaContext).na = _prevctx
			p.PushNewRecursionContext(localctx, _startState, Practica4ParserRULE_lista)
			p.SetState(16)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(17)
				p.Match(Practica4ParserPARA)
			}
			{
				p.SetState(18)

				var _m = p.Match(Practica4ParserNUMBER)

				localctx.(*ListaContext).num = _m
			}
			{
				p.SetState(19)
				p.Match(Practica4ParserCOMA)
			}
			{
				p.SetState(20)

				var _m = p.Match(Practica4ParserNUMBER)

				localctx.(*ListaContext).num2 = _m
			}
			{
				p.SetState(21)
				p.Match(Practica4ParserPARC)
			}

			x := localctx.(*ListaContext).GetNa().GetN().A + " " + (func() string {
				if localctx.(*ListaContext).GetNum() == nil {
					return ""
				} else {
					return localctx.(*ListaContext).GetNum().GetText()
				}
			}()) + " " + (func() string {
				if localctx.(*ListaContext).GetNum2() == nil {
					return ""
				} else {
					return localctx.(*ListaContext).GetNum2().GetText()
				}
			}())
			y := localctx.(*ListaContext).GetNa().GetN().B
			localctx.(*ListaContext).n = node.NewNodo(x, y)

		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Practica4Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ListaContext = nil
		if localctx != nil {
			t = localctx.(*ListaContext)
		}
		return p.Lista_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Practica4Parser) Lista_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
