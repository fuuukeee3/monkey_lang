package ast

import (
	"github.com/fuuukeee3/monkey_lang/token"
)

// Node is ASTのインターフェース
type Node interface {
	TokenLiteral() string
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

// Identifier is 識別子の構造体
type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode is
func (i *Identifier) expressionNode() {}

// TokenLiteral is ノードが関連づけられているトークンのリテラルを返す
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement is return文の構造体
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// statementNode is
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
