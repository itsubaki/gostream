package lexer

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	literal_begin
	IDENT
	STRING
	INT
	FLOAT
	literal_end

	operator_begin
	ASTERISK  // *
	DOT       // .
	COMMA     // ,
	COLON     // :
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LARGER    // >
	LESS      // <
	operator_end

	keyword_begin
	SELECT       // SELECT
	FROM         // FROM
	TIME         // TIME
	LENGTH       // LENGTH
	TIME_BATCH   // TIME_BATCH
	LENGTH_BATCH // LENGTH_BATCH
	SEC          // SEC
	MIN          // MIN
	WHERE        // WHERE
	keyword_end
)

var Tokens = [...]string{
	// Specials
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	WHITESPACE: "WHITESPACE",

	// Literals
	IDENT:  "IDENT",
	STRING: "STRING",
	INT:    "INT",
	FLOAT:  "FLOAT",

	// Operators
	ASTERISK:  "*",
	DOT:       ".",
	COMMA:     ",",
	COLON:     ":",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	LARGER:    ">",
	LESS:      "<",

	// Keywords
	SELECT:       "SELECT",
	FROM:         "FROM",
	TIME:         "TIME",
	LENGTH:       "LENGTH",
	TIME_BATCH:   "TIME_BATCH",
	LENGTH_BATCH: "LENGTH_BATCH",
	SEC:          "SEC",
	MIN:          "MIN",
	WHERE:        "WHERE",
}

func IsBasicLit(token Token) bool {
	if token == IDENT || token == STRING || token == INT || token == FLOAT {
		return true
	}

	return false
}