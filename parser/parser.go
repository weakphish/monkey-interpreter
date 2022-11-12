package parser

import (
	"github.com/weakphish/monkey-interpreter/ast"
	"github.com/weakphish/monkey-interpreter/lexer"
	"github.com/weakphish/monkey-interpreter/token"
)

/*
 * The basic idea of recursive-descent parsing is that there is an entry point,
 * parseProgram, that constructs the root node of the AST. "It then builds
 * child nodes, the statements, by calling other functions that know which
 * AST node to construct based on the current token."
 */

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil // TODO
}
