package ast

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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
