package main

import (
	"bufio"
	"strings"
)

type Command int
const (
	C_NULL       Command = 0 // empty command
	C_ARITHMETIC Command = 1
	C_PUSH       Command = 2
	C_POP        Command = 3
	C_LABEL      Command = 4
	C_GOTO       Command = 5
	C_IF         Command = 6
	C_FUNCTION   Command = 7
	C_RETURN     Command = 8
	C_CALL       Command = 9
)

type Parser struct {
	scanner *bufio.Scanner
	text string
	arg1 string
	arg2 int
}

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() Command
	Arg1() string
	Arg2() int
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", "", ""}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {
	p.text = p.scanner.Text()
	if len(p.text) <= 0 {
		return
	}
	// normalization: remove comments
	tokens := strings.SplitN(p.text, "//", 2)
	if len(tokens) > 0 {
		p.text = tokens[0]
	}
	// normalization: remove spaces and tabs
	p.text = strings.TrimSpace(p.text)
}

func (p *Parser) CommandType() Command {
	p.arg1 = ""
	p.arg2 = 0
	if len(p.text) == 0 {
		// empty token
		return C_NULL
	}
	return C_NULL
}

func (p *Parser) Arg1() string {
	return p.arg1
}

func (p *Parser) Arg2() int {
	return p.arg2
}

