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
		{"t", "Display only thesaurus results"},
		{"ud", "Also query Urban Dictionary for results"},
	}
	SUBCOMMAND   = os.Args[1]
	UsageExample = fmt.Sprintf("$ %s linux [FLAGS]", CMD)
)

const (
	UsageHeadline = "Dicthesaurus requires only a single command: the word you want to search for."
)
