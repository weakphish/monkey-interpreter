package ast

import "github.com/weakphish/monkey-interpreter/token"

type Node interface {
	TokenLiteral() string // used for debugging / testing
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// "This Program node is going to be the root node of every AST our parser produces"
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Let Statements
type LetStatement struct {
	Token token.Token
	Name  *Identifier // Holds the identifier of the binding
	Value Expression  // The expression that produces the value
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
