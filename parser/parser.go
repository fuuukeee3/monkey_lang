package parser

import (
	"github.com/fuuukeee3/monkey_lang/ast"
	"github.com/fuuukeee3/monkey_lang/lexer"
	"github.com/fuuukeee3/monkey_lang/token"
)

// Parser is パーサー構造体
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

// New is パーサーインスタンスの作成
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// curTokenとpeekTokenにセットするため2回読み込む
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken is トークンを読み進める
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram is
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
