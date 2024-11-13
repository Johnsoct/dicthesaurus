package presentation

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Johnsoct/dicthesaurus/repository"
)

type (
	Senses         []string
	SenseSequences map[int]Senses
)

type (
	VerbDivider map[string]SenseSequences
	Definitions map[string]VerbDivider
)

func boldText(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func headerText(text string) string {
	return "\033[47;30m" + text + "\033[0m"
}

func upperCaseText(text string) string {
	return "" + text + "k"
}

func stripBoldColonTokens(text string) string {
	return strings.ReplaceAll(text, "{bc}", "")
}

func stripCrossReferenceGroupingTokens(text string) string {
	re := regexp.MustCompile(`\|+[0-9]+[a-z]*}`)
	return re.ReplaceAllLiteralString(text, "")
}

func stripCrossReferenceTokens(text string) string {
	firstHalf := strings.ReplaceAll(text, "{sx|", "")
	return strings.ReplaceAll(firstHalf, "||}", "")
}

func formatSenseText(text string) string {
	t := stripBoldColonTokens(text)
	t = stripCrossReferenceTokens(t)
	t = stripCrossReferenceGroupingTokens(t)
	return t
}

func formatSequence(sseqn int, sn int, text string) string {
	// Do not prent the sseq number for every sense
	if sn == 0 {
		return fmt.Sprintf("%d\t%d : %s", sseqn+1, sn+1, formatSenseText(text))
	}
	return fmt.Sprintf("\t%d : %s", sn+1, formatSenseText(text))
}

func prepareSenseSequences(data []repository.MWDResult) Definitions {
	definitions := make(Definitions)

	// data could have multiple results
	for _, v := range data {
		// If the data object doesn't have the property "hom,"
		// it's not an identical spelling as the searched word
		if v.Def == nil {
			continue
		}

		// Do not overwrite definitions[v.Fl]; will later append to definitions[v.Fl][verbDivider][sn]
		if _, ok := definitions[v.Fl]; !ok {
			definitions[v.Fl] = make(VerbDivider)
		}

		// each data object could have multiple definitions
		for _, def := range v.Def {
			verbDivider := def.Vd

			// Verbs have dividers; nouns, adjectives do not
			// Use Fl as key name for these cases to make the
			// parsing below much more simple and readable
			if verbDivider == "" {
				verbDivider = v.Fl
			}

			// Do not overwrite verb dividers; will later append to definitions[v.Fl][verbDivider][sn]
			if _, ok := definitions[v.Fl][verbDivider]; !ok {
				definitions[v.Fl][verbDivider] = make(SenseSequences)
			}

			// Each index of sseq (sn - sense number) is a group of senses
			// At each index is an array of the senses within that group
			for sn, sseq := range def.Sseq {

				// Do not overwrite sense sequences; will later append to
				if _, ok := definitions[v.Fl][verbDivider][sn]; !ok {
					definitions[v.Fl][verbDivider][sn] = make(Senses, 0)
				}

				for _, sense := range sseq {
					// Ignore useless defining texts
					if sense[1].Dt == nil {
						continue
					}
					if sense[1].Dt[0][0] != "text" {
						continue
					}

					// Only capturing the first index of dt (other indexes are not immediate definitions)
					dt := sense[1].Dt[0][1]

					// Make sure dt is a string value (definition)
					if d, ok := dt.(string); ok {
						definitions[v.Fl][verbDivider][sn] = append(definitions[v.Fl][verbDivider][sn], d)
					}
				}
			}
		}
	}

	return definitions
}

func Print(data []repository.MWDResult) {
	if data == nil {
		fmt.Println("I'm not sure how you got here, but something is wrong. Sorry, try again.")
		os.Exit(1)
	}

	// Example of verbs: definitions["verb"]["intransitive verb"]
	definitions := prepareSenseSequences(data)

	for fl := range definitions {
		for divider, value := range definitions[fl] {
			fmt.Printf("\n\n\t%s\n\n", headerText(strings.ToUpper(divider)))

			// Iterate through the sequences in order (matches Merriam-Webster's results)
			for i := range len(value) {
				for sn, sense := range value[i] {
					fmt.Println(formatSequence(i, sn, sense))
				}

				fmt.Println()
			}
		}
	}
}
