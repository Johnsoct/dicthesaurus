package business

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	HFlag      = flag.Bool("h", false, "Show help")
	wordSubCmd = flag.NewFlagSet("word", flag.ExitOnError)
	EFlag      = wordSubCmd.Bool("e", false, "Display the word in a sentence")
	TFlag      = wordSubCmd.Bool("t", false, "Display only thesaurus results")
	UDFlag     = wordSubCmd.Bool("ud", false, "Also query Urban Dictionary for results")

	lookupValue string
)

func cliUsageError() {
	base := filepath.Base(os.Args[0])

	fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s <word> [flags]\n\n", base)

	flag.PrintDefaults()

	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "Dicthesaurus requires only a single command: the word you want to search for.")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "\"dt linux\"")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "FLAGS")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-e   :  Include the word used in a sentence")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-t   :  Only return a word's result from the thesaurus")
	fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", "-ud  :  Return defintions from Urban Dictionary")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")

	os.Exit(1)
}

func GetLookupValue() string {
	return lookupValue
}

func overwriteFlagUsageDefault() {
	// Overwrite the default error output
	flag.Usage = cliUsageError
	wordSubCmd.Usage = cliUsageError
}

func parseFlags() {
	// Handle no subcommand or flag
	if len(os.Args) == 1 {
		cliUsageError()
	}

	// Handle the lack of a subcommand (word to search)
	if strings.HasPrefix(os.Args[1], "-") {
		cliUsageError()
	}

	// Handle the 3rd argument not being a flag
	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
		cliUsageError()
	}

	// Parse each flag set
	flag.Parse()
	wordSubCmd.Parse(os.Args[2:]) // everything after the subcommand
}

func subcommand() string {
	return os.Args[1]
}

func updateState() {
	lookupValue = subcommand()
}

func init() {
	overwriteFlagUsageDefault()
	parseFlags()
	updateState()
}
