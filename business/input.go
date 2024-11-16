package business

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Johnsoct/dicthesaurus/repository"
	"github.com/Johnsoct/dicthesaurus/utils"
)

var (
	// flags
	HFlag       = flag.Bool("h", false, "Show help")
	subcmdFlags = flag.NewFlagSet("word", flag.ExitOnError)
	EFlag       = subcmdFlags.Bool("e", false, "Display the word in a sentence")
	SSFlag      = subcmdFlags.Bool("ss", false, "Display a short and sweet version of the definition")
	TFlag       = subcmdFlags.Bool("t", false, "Display only thesaurus results")

	// "Constants"
	flags = []repository.Flag{
		{Flag: "e", Description: "Display the word in a sentence"},
		{Flag: "ss", Description: "Display a short and sweet version of the definition"},
		{Flag: "t", Description: "Display thesaurus instead of dictionary results"},
	}
	usageExample  = fmt.Sprintf("$ %s linux [FLAGS]", parseCmd(os.Args))
	usageHeadline = "Dicthesaurus requires only a single command: the word you want to search for."
)

// ERROR HANDLING
// ERROR HANDLING
// ERROR HANDLING

func cliUsageError(headline, headlineExample string, flags []repository.Flag) {
	fmt.Fprintf(flag.CommandLine.Output(), "\n%s\n\n", utils.BoldText(headline))
	fmt.Fprintf(flag.CommandLine.Output(), "%s", headlineExample)
	fmt.Fprintf(flag.CommandLine.Output(), "%s", utils.FormatHeader("FLAGS"))

	for _, v := range flags {
		fmt.Fprintf(flag.CommandLine.Output(), "\t-%s   :  %s\n", v.Flag, v.Description)
	}

	fmt.Fprintf(flag.CommandLine.Output(), "\n")

	os.Exit(1)
}

func overwriteFlagUsageDefault() {
	// Overwrite the our flagsets error output (flag, subcmdFlags)
	flag.Usage = func() { cliUsageError(usageHeadline, usageExample, flags) }
	subcmdFlags.Usage = func() { cliUsageError(usageHeadline, usageExample, flags) }
}

// FLAGS
// FLAGS
// FLAGS

func hasInvalidArgsAfterSubcmd(args []string) bool {
	// Handle the 3rd+ argument not being a flag
	if len(args) > 2 && !strings.HasPrefix(args[2], "-") {
		return true
	}
	return false
}

func parseFlags(args []string) {
	// Parse each flag set
	flag.Parse()
	subcmdFlags.Parse(args[2:]) // everything after the subcommand
}

// COMMAND
// COMMAND
// COMMAND
func parseCmd(args []string) string {
	return filepath.Base(args[0])
}

// SUBCOMMAND
// SUBCOMMAND
// SUBCOMMAND

func hasInvalidSubcmd(args []string) bool {
	// Handle the lack of a subcmd (word to search)
	if len(args) < 2 {
		return true
	}
	return strings.HasPrefix(ParseSubcmd(args), "-")
}

func ParseSubcmd(args []string) string {
	return args[1]
}

// INIT FUNCTIONS
// INIT FUNCTIONS
// INIT FUNCTIONS

func initMessage(thesaurus bool, args []string) {
	searchingFor := "dictionary"

	if thesaurus {
		searchingFor = "thesaurus"
	}

	fmt.Printf("\n\tSearching %s for \"%s\" ...\n", searchingFor, ParseSubcmd(args))
}

func init() {
	// Handle missing subcommand or invalid 3rd+ arguments
	if hasInvalidSubcmd(os.Args) || hasInvalidArgsAfterSubcmd(os.Args) {
		cliUsageError(usageHeadline, usageExample, flags)
	}

	overwriteFlagUsageDefault()
	parseFlags(os.Args)
	initMessage(*TFlag, os.Args)
}
