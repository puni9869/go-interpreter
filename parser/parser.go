package parser

import (
	"play/ast"
	"play/lexer"
	"play/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expectedPeek(typ token.TokenType) bool {
	if p.peekTokenIs(typ) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) curTokenIs(typ token.TokenType) bool {
	return p.curToken.Type == typ
}

func (p *Parser) peekTokenIs(typ token.TokenType) bool {
	return p.peekToken.Type == typ
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() *ast.LetStatement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Reads current and peek token
	p.nextToken()
	p.nextToken()
	return p
}
