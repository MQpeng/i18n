package command

import "flag"

var (
	File = flag.String("f", "", "source file path")
	Language = flag.String("l", "", "language file path")
	NewLanguage = flag.String("nl", "", "language file path")
	Out = flag.String("o", "", "new language file path")
	Write = flag.Bool("w", false, "write func")
	Regex = flag.String("r", "['\"]([^{]*)translate", "Regex pattern")
)

func Parse(){
	flag.Parse()
}