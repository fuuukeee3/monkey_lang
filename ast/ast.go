package ast

import (
	"bytes"

	"github.com/fuuukeee3/monkey_lang/token"
)

// Node is ASTのインターフェース
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is 文のインターフェース
type Statement interface {
	Node
	statementNode()
}

// Expression is 式のインターフェース
type Expression interface {
	Node
	expressionNode()
}

// Program is ASTのルートノード
// プログラムの全ての文を配列で保持する
type Program struct {
	Statements []Statement
}

// TokenLiteral is ノードが関連づけられているトークンのリテラルを返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement is let文の構造体
type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier // 識別子
	Value Expression  // 値を生成する式
}

// statementNode is
func (ls *LetStatement) statementNode() {}

// TokenLiteral is ノードが関連づけられているトークンのリテラルを返す
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String is
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// Identifier is 識別子の構造体
type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode is
func (i *Identifier) expressionNode() {}

// TokenLiteral is ノードが関連づけられているトークンのリテラルを返す
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String is
func (i *Identifier) String() string { return i.Value }

// ReturnStatement is return文の構造体
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// statementNode is
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String is
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is 式の構造体
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// statementNode is
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral is
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String is
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
