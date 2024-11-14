package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	CMD   = filepath.Base(os.Args[0])
	Flags = []Flag{
		{"e", "Display the word in a sentence"},
		{"ss", "Display a short and sweet version of the definition"},
		{"t", "Display only thesaurus results"},
	}
	SUBCOMMAND    = os.Args[1]
	UsageExample  = fmt.Sprintf("$ %s linux [FLAGS]", CMD)
	UsageHeadline = "Dicthesaurus requires only a single command: the word you want to search for."
)
