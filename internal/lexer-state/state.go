package lexerstate

import (
	"github.com/MonsieurTa/go-lexer"
)

const (
	IDENT = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
)

func IdentState(l lexer.Lexer) lexer.StateFn {
	r := l.Peek()
	if r == lexer.EOFRune {
		return nil
	} else if r == '#' {
		return CommentState
	} else if !l.Accept(IDENT) {
		return l.Errorf("expected ident got %q", r)
	}

	l.AcceptRun(IDENT)
	l.Emit(IdentToken)

	r = l.Peek()
	if r == ':' {
		return ColonState
	} else if r == ';' {
		return SemicolonState
	}
	return RParState
}

func CommentState(l lexer.Lexer) lexer.StateFn {
	l.Accept("#")
	var r rune
	for r != lexer.EOFRune && r != '\n' { // skip line
		r = l.Next()
	}
	if r == lexer.EOFRune {
		l.Ignore()
		l.Emit(EOF)
		return nil
	}
	l.Ignore()
	return IdentState
}

func ColonState(l lexer.Lexer) lexer.StateFn {
	if !l.Accept(":") {
		if l.Peek() == lexer.EOFRune {
			return l.Errorf("unexpected EOF")
		}
		return l.Errorf("expected ':' got %q", l.Peek())
	}

	l.Emit(ColonToken)

	r := l.Peek()
	if r == '(' {
		return LParState
	} else if r >= '0' && r <= '9' {
		return NumberState
	} else if r == ':' {
		return ColonState
	}
	return l.Errorf("expected '(' or number got %q", r)
}

func LParState(l lexer.Lexer) lexer.StateFn {
	if !l.Accept("(") {
		return l.Errorf("expected '(' got %q", l.Peek())
	}
	l.Ignore()
	return IdentState
}

func RParState(l lexer.Lexer) lexer.StateFn {
	if !l.Accept(")") {
		return l.Errorf("expected ')' got %q", string(l.Peek()))
	}
	l.Ignore()
	r := l.Next()
	if r == lexer.EOFRune {
		return nil
	} else if r == ':' {
		l.Backup()
		return ColonState
	} else if r == '\n' {
		l.Ignore()
		return IdentState
	}
	return l.Errorf("expected semicolon or EOF got %q", string(r))
}

func NumberState(l lexer.Lexer) lexer.StateFn {
	digits := "0123456789"
	if !l.Accept(digits) {
		return l.Errorf("expected digit got %q", l.Peek())
	}
	l.AcceptRun("0123456789")
	l.Emit(IntToken)

	r := l.Next()
	if r == lexer.EOFRune {
		return nil
	} else if r == '\n' {
		l.Ignore()
		return IdentState
	} else if r == ')' {
		l.Backup()
		return RParState
	} else if r == ';' {
		l.Backup()
		return SemicolonState
	}
	return l.Errorf("expected '\\n' or ')' got %q", r)
}

func SemicolonState(l lexer.Lexer) lexer.StateFn {
	if !l.Accept(";") {
		return l.Errorf("expected ';' got %q", string(l.Peek()))
	}
	l.Emit(SemicolonToken)
	r := l.Peek()
	if r == lexer.EOFRune {
		return l.Errorf("unexpected EOF")
	}
	return IdentState
}
