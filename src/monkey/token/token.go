package token

type TokenType string

var keywords = map[string]TokenType{
    "fn"  : FUNCTION,
    "let" : LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL   = "ILLEGAL"
    EOF       = "EOF"

    // Identifieres and literals
    IDENT     = "IDENT"
    INT       = "INT"

    // Operators
    ASSIGN    = "="
    PLUS      = "+"

    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN  = ")"
    LBRACE  = "{"
    RBRACE  = "}"

    FUNCTION  = "FUNCTION"
    LET       = "LET"
)