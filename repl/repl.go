package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/weakphish/monkey-interpreter/lexer"
	"github.com/weakphish/monkey-interpreter/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.TokenType(token.EOF); tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
