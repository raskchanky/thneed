package repl

import (
	"bufio"
	"fmt"
	"os"
)

func Read(in string) string {
	return in
}

func Eval(in string) string {
	return in
}

func Print(in string) string {
	return in
}

func Rep(in string) string {
	return Print(Eval(Read(in)))
}

func MainLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("user> ")
		line, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			} else {
				fmt.Printf("err: %q\n", err)
				os.Exit(1)
			}
		}

		result := Rep(line)
		fmt.Println(result)
	}
}
