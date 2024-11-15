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
	SSFlag     = wordSubCmd.Bool("ss", false, "Display a short and sweet version of the definition")
	TFlag      = wordSubCmd.Bool("t", false, "Display only thesaurus results")
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

func overwriteFlagUsageDefault() {
	// Overwrite the default error output
	flag.Usage = cliUsageError
	wordSubCmd.Usage = cliUsageError
}

func parseCommand() {
	// Handle missing subcommand or flag
	if len(os.Args) == 1 {
		cliUsageError()
	}
}

func parseFlags() {
	// Handle the 3rd argument not being a flag
	if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
		cliUsageError()
	}
}

func parseSubcommand(subcommand string) {
	// Handle the lack of a subcommand (word to search)
	if strings.HasPrefix(subcommand, "-") {
		cliUsageError()
	}
}

func parse() {
	// Cover the loopholes in "flag" parsing
	parseCommand()
	parseSubcommand(repository.SUBCOMMAND)
	parseFlags()

	// Parse each flag set
	flag.Parse()
	wordSubCmd.Parse(os.Args[2:]) // everything after the subcommand
}

func init() {
	overwriteFlagUsageDefault()
	parse()

	searchingFor := "dictionary"

	if *TFlag {
		searchingFor = "thesaurus"
	}

	fmt.Printf("\n\tSearching %s for \"%s\" ... \n\n", searchingFor, repository.SUBCOMMAND)
}
