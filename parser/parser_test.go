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

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = x + y; 
`
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()
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
