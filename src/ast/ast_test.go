package ast

import (
	"monkey_interpreter/src/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{Statements: []Statement{
		&LetStatement{
			Token: token.Token{Type: token.Let, Literal: "let"},
			Name: &Indentifier{
				Token: token.Token{Type: token.Identifier, Literal: "myVar"},
				Value: "myVar",
			},

			Value: &Indentifier{
				Token: token.Token{Type: token.Identifier, Literal: "anotherVar"},
				Value: "anotherVar",
			},
		},
	},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. go=%q", program.String())
	}
}
