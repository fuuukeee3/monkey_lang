package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fuuukeee3/monkey_lang/evalutor"
	"github.com/fuuukeee3/monkey_lang/lexer"
	"github.com/fuuukeee3/monkey_lang/parser"
)

// PROMPT is プロンプト
const PROMPT = ">> "

// Start is REPLの開始
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluted := evalutor.Eval(program)
		if evaluted != nil {
			io.WriteString(out, evaluted.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
