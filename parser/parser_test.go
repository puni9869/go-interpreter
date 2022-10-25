package parser

import (
	"play/ast"
	"play/lexer"
	"testing"
)

func testLetStatement(t *testing.T, stm ast.Statement, name string) bool {
	if stm.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stm.TokenLiteral())
		return false
	}
	letStm, ok := stm.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt is not *ast.LetStatement. got %T", stm)
		return false
	}

	if letStm.Name.Value != name {
		t.Errorf("letStm.Name.Value not '%s'. got %s", name, letStm.Name.Value)
		return false
	}

	if letStm.Name.TokenLiteral() != name {
		t.Errorf("letStm.Name not '%s'. got %s", name, letStm.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = x + y; 
`
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statemensts. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

	}
}

func TestLetWithErrorStatements(t *testing.T) {
	input := `
	let  5;
	let x;
	let a+b ; 
`
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("Parse() returned nil")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 3;
	return 10;
	return 10000;
`
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statemensts. got=%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return'. got %q", returnStmt.TokenLiteral())
		}
	}

}
