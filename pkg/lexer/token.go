package lexer

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	literal_begin
	IDENTIFIER
	STRING
	INT
	FLOAT
	literal_end

	operator_begin
	ASTERISK  // *
	DOT       // .
	COMMA     // ,
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
	COUNT        // COUNT
	SUM          // SUM
	AVG          // AVG
	MAX          // MAX
	MED          // MED
	FROM         // FROM
	TIME         // TIME
	LENGTH       // LENGTH
	TIME_BATCH   // TIME_BATCH
	LENGTH_BATCH // LENGTH_BATCH
	SEC          // SEC
	MIN          // MIN
	WHERE        // WHERE
	AND          // AND
	OR           // OR
	keyword_end
)

var tokens = [...]string{
	// Specials
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	WHITESPACE: "WHITESPACE",

	// Literals
	IDENTIFIER: "IDENTIFIER",
	STRING:     "STRING",
	INT:        "INT",
	FLOAT:      "FLOAT",

	// Operators
	DOT:       ".",
	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	LARGER:    ">",
	LESS:      "<",

	// Keywords
	SELECT:       "SELECT",
	ASTERISK:     "*",
	COUNT:        "COUNT",
	SUM:          "SUM",
	AVG:          "AVG",
	MAX:          "MAX",
	MED:          "MED",
	FROM:         "FROM",
	TIME:         "TIME",
	LENGTH:       "LENGTH",
	TIME_BATCH:   "TIME_BATCH",
	LENGTH_BATCH: "LENGTH_BATCH",
	SEC:          "SEC",
	MIN:          "MIN",
	WHERE:        "WHERE",
	AND:          "AND",
	OR:           "OR",
}
