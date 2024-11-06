package business

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Johnsoct/dicthesaurus/repository"
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
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", repository.UsageHeadline)
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", repository.UsageExample)
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "FLAGS")

	for _, v := range repository.Flags {
		fmt.Fprintf(flag.CommandLine.Output(), "\t-%s   :  %s\n", v.Flag, v.Description)
	}

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

func updateState() {
	lookupValue = repository.SUBCOMMAND
}

func init() {
	overwriteFlagUsageDefault()
	parseFlags()
	updateState()

	fmt.Printf("\nSearching for \"%s\" ... \n\n", repository.SUBCOMMAND)
}
