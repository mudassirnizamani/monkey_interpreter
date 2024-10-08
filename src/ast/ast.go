package ast

import "monkey_interpreter/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET Token
	Name  *Indentifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	RetrunValue Expression
}

func (rs *ReturnStatement) statementNode()      {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

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

