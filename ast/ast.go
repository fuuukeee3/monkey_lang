package ast

import (
	"bytes"
	"strings"

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

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}
