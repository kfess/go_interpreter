package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kfess/go_interpreter/lexar"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexar.New(line)
		for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
