package presentation

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/repository"
)

type (
	Senses         []string
	SenseSequences map[int]Senses
)

type (
	VerbDivider map[string]SenseSequences
	Definitions map[string]VerbDivider
	Thesaurus   map[string][][]string
)

func boldText(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func underlineText(text string) string {
	return "\033[4m" + text + "\033[0m"
}

// Must be applied before any other console code because of strings.toUpper
func uppercaseText(text string) string {
	return "\033[1;4m" + strings.ToUpper(text) + "\033[0m"
}

func replaceTokens(text string) string {
	type Replace struct {
		submatch  string
		replaceFn func(text string) string
	}
	replaces := map[string]Replace{
		`\{a_link\|(\w+)\}`: {
			replaceFn: func(text string) string {
				return underlineText(text)
			},
			submatch: `\|(\w+)`,
		},
		`\{sx\|(\w+)\|*\}`: {
			replaceFn: func(text string) string {
				uppercased := uppercaseText(text)
				underlined := underlineText(uppercased)
				return underlined
			},
			submatch: `\|(\w+)`,
		},
	}

	replaced := text

	for regex, replace := range replaces {
		re := regexp.MustCompile(regex)
		subre := regexp.MustCompile(replace.submatch)
		submatches := 0

		replaced = re.ReplaceAllStringFunc(replaced, func(substring string) string {
			submatch := replace.replaceFn(subre.FindAllStringSubmatch(substring, -1)[0][1])

			if submatches == 0 {
				submatch = ": " + submatch
			}

			submatches++

			return submatch
		})
	}

	return replaced
}

func stripCrossReferenceGroupingTokens(text string) string {
	re := regexp.MustCompile(`\|+[0-9]+[a-z]*}`)
	return re.ReplaceAllLiteralString(text, "")
}

func stripTokens(text string) string {
	prefixes := []string{
		`{bc}`,
		`{sx|`,
	}
	suffixes := []string{
		`\|*}`,
		`}`,
		`\|+[0-9]+[a-z]*}`,
	}

	stripped := text

	for _, v := range prefixes {
		re := regexp.MustCompile(v)
		stripped = re.ReplaceAllString(stripped, "")
	}

	for _, v := range suffixes {
		re := regexp.MustCompile(v)
		stripped = re.ReplaceAllString(stripped, "")
	}

	return stripped
}

func formatAntsSyns(s []string) string {
	rowFormat := "\t" + strings.Repeat("%-15s", len(s)) + "\n\n"
	anys := make([]any, len(s))

	for i, v := range s {
		anys[i] = v
	}

	return fmt.Sprintf(rowFormat, anys...)
}

func formatHeader(h string) string {
	return fmt.Sprintf("\n\n\t%s\n\n", uppercaseText(h))
}

func formatSenseText(text string) string {
	// Replace before stripping... hehe (replace relies on the tokens)
	replaced := replaceTokens(text)
	stripped := stripTokens(replaced)

	return stripped
}

func formatSequence(sseqn int, sn int, text string) string {
	// Do not prent the sseq number for every sense
	if sn == 0 {
		return fmt.Sprintf("%d\t%s", sseqn+1, formatSenseText(text))
	}
	return fmt.Sprintf("\t%s", formatSenseText(text))
}

func prepareAntonyms(data []repository.MWTResult) [][]string {
	antonyms := make([][]string, len(data))

	for i, v := range data {
		for _, ants := range v.Meta.Ants {
			antonyms[i] = ants
		}
	}

	return antonyms
}

func prepareDefinitions(data []repository.MWDResult) Definitions {
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

			if *business.SSFlag {
				definitions[v.Fl][verbDivider][0] = v.Shortdef
			} else {
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
	}

	return definitions
}

func prepareSynonyms(data []repository.MWTResult) Thesaurus {
	synonyms := make(Thesaurus)

	for _, v := range data {
		// Ignore all the synonyms for stems off of the SUBCOMMAND
		if i := slices.Compare(v.Meta.Stems, []string{repository.SUBCOMMAND}); i != 0 {
			continue
		}

		// Do not overwrite synonyms[v.Fl]
		if _, ok := synonyms[v.Fl]; !ok {
			synonyms[v.Fl] = make([][]string, len(v.Meta.Syns))
		}

		for i, syns := range v.Meta.Syns {
			synonyms[v.Fl][i] = syns
		}
	}

	return synonyms
}

func printDictionary(data []repository.MWDResult) {
	// Example of verbs: definitions["verb"]["intransitive verb"]
	definitions := prepareDefinitions(data)

	for fl := range definitions {
		for divider, value := range definitions[fl] {
			fmt.Println(formatHeader(divider))

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

func printThesaurus(data []repository.MWTResult) {
	synonyms := prepareSynonyms(data)
	// antonyms := prepareAntonyms(data)

	for fl, rows := range synonyms {
		fmt.Println(formatHeader(fl))

		for _, row := range rows {
			fmt.Println(formatAntsSyns(row))
		}
	}

	// for i, fl := range antonyms {
	// 	fmt.Println(fl)
	// 	fmt.Println(formatAntsSyns(antonyms[fl]))
	// }
}

func Print[T repository.APIData](data any) {
	if data == nil {
		fmt.Println("The API response was empty. Please check your API keys.")
		os.Exit(1)
	}

	switch data := data.(type) {
	case []repository.MWDResult:
		printDictionary(data)
	case []repository.MWTResult:
		printThesaurus(data)
	}
}
