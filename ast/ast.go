package ast

import (
	"bytes"

	"github.com/kfess/go_interpreter/token"
)

type Node interface {
	// return the literal token to which the node is related,
	// for debug and test
	TokenLiteral() string
	String() string
}

// 文
type Statement interface {
	Node
	statementNode() // dammy method for Go Compiler
}

// 式
type Expression interface {
	Node
	expressionNode() // dammy method for Go Compiler
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stms := range p.Statements {
		out.WriteString(stms.String())
	}
	return out.String()
}

// vvv v   vv
// let x = 10
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // x の部分
	Value Expression  // 10 の部分
}

// dammy function for Go Compiler
func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ") // let
	out.WriteString(ls.Name.String())        // x
	out.WriteString(" = ")                   // =

	if ls.Value != nil {
		out.WriteString(ls.Value.String()) // 5
	}

	out.WriteString(";")

	return out.String()
}

// /   v   vv
// let x = 10
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

// dammy function for Go Compiler
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// vvvvvv vv
// return 10
type ReturnStatement struct {
	Token       token.Token // token.RETURN トークン
	ReturnValue Expression  // 10 の部分
}

// dammy function for Go Compiler
func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())

	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
