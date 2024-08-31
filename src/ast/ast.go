package ast

import "monkey_interpreter/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() string
}

type Expression interface {
	Node
	expressionNode() string
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET Token
	Name *Indentifier
	Value Expression
}

type Indentifier struct {
	Token token.Token
	Value string
}

func (i *Indentifier) expressionNode() {}

func (i *Indentifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


