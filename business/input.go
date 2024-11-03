package business

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	HFlag      = flag.Bool("h", false, "Show help")
	WordSubCmd = flag.NewFlagSet("word", flag.ExitOnError)
	DFlag      = WordSubCmd.Bool("d", false, "Display only dictionary results")
	SSFlag     = WordSubCmd.Bool("ss", false, "Display format is 'sweet and simple'")
	TFlag      = WordSubCmd.Bool("t", false, "Display only thesaurus results")
	UDFlag     = WordSubCmd.Bool("ud", false, "Also query Urban Dictionary for results")

	LookupValue string
)

func manual() {
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "Dicthesaurus requires only a single command: the word you want to search for.")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "\"dt linux\"")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "FLAGS")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-d   :  Only return a word's result from the dictionary")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-t   :  Only return a word's result from the thesaurus")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-ss  :  Keep the output sweet and simple")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-ud  :  Return defintions from Urban Dictionary")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
}

func subcommand() string {
	return os.Args[1]
}

func usage() {
	base := filepath.Base(os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s <word> [flags]\n\n", base)
	flag.PrintDefaults()
	manual()
}

func init() {
	// Overwrite the default error output
	flag.Usage = usage
	WordSubCmd.Usage = usage

	// Parse each flag set
	flag.Parse()
	WordSubCmd.Parse(os.Args[2:]) // everything after the subcommand

	// Globally store the word to look up
	LookupValue = subcommand()
}
