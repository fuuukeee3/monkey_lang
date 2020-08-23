package lexer

import (
	"github.com/fuuukeee3/monkey_lang/token"
)

// Lexer is 字句解析機
type Lexer struct {
	// 解析対象の文字列
	input string
	// 現在の文字の位置
	position int
	// 次に読み込む文字の位置
	readPosition int
	// 現在検査中の文字
	ch byte
}

// New is lexerの作成
// @params input[String] 字句解析対象の文字列
// @return Lexerインスタンス
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar is input文字列の現在位置を進める
// @params
// @return
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// input終端に達した場合
		l.ch = 0
	} else {
		// 終端では無い場合,chに次の文字をセット
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken is 検査中の文字に応じたTokenを返す
// @params
// @return Tokenインスタンス
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 空白はスキップする
	l.skipWhitespace()

	switch l.ch {
	case '=':
		// 「=」 or 「==」で処理を分岐
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// 「=」 or 「!=」で処理を分岐
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 識別子・キーワードの処理
			// 英字の場合は切れ目まで読み進める
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// 数値の処理
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken is tokenの作成
// @params tokenType トークン種別
// @params ch 検査中の文字
// @return Tokenインスタンス
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier is 識別子の読み込み
// 英字を非英字まで読み進める
// @params
// @return [Strring] 識別子
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter is 英字か非英字かを判別する
// @params [Byte] 検査対象文字
// @return [Boolean]
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skipWhitespace is 空白を読み進める
// @params
// @return
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readNumber is 数値を読み込む
// @params
// @return [String]
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit is 数値かどうか判定する
// @params ch [Byte] 検査対象バイト
// @return [Boolean]
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar is 次の位置の文字を返す。読み込み位置は進めない。
// @params
// @return [Byte] 次の位置の文字
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
