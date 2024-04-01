package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					//Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					//Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;\n" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}