package token

// TookenTypeの種類
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子, リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// TokenType is トークンの種類を保持するフィールド
// 整数(INTEGER), 識別子(IDENTIFER)など
type TokenType string

// Token is トークンの型
type Token struct {
	// トークンの種別
	Type TokenType

	// リテラル値
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent is 引数に対応するTokenTypeを返却する
// keywordsに定義されていれば定義されたType、定義されていない場合は識別子(IDENT)
// @params ident [String]
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
