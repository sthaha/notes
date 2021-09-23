package lexer

import "unicode"

type tokenizer struct {
	input string
	index int
	len   int
}

func newTokenizer(input string) *tokenizer {
	return &tokenizer{
		input: input,
		index: 0,
		len:   len(input),
	}
}

func (t *tokenizer) Next() *token {
	ch, eof := t.read()
	if eof {
		return &token{Eof, ""}
	}

	switch ch {
	case "(":
		return &token{LeftParan, "("}
	case ")":
		return &token{RightParan, ")"}
	}
	return &token{Let, ""}
}

func (t *tokenizer) read() (string, bool) {
	if t.index >= t.len {
		return "", true
	}

	for t.index < t.len {
		t.index++

		if unicode.IsSpace(rune(t.input[t.index])) {
			continue
		}
	}
}
