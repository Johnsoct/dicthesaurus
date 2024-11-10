package presentation

import (
	"fmt"
	"os"

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

func formatSenseText(text string) string {
	return text
}

func prepareSenseSequences(data []repository.MWDResult) Definitions {
	definitions := make(Definitions)

	// data could have multiple results
	for _, v := range data {
		// If the data object doesn't have the property "hom,"
		// it's not an identical spelling as the searched word
		if v.Hom == 0 {
			continue
		}

		// Do not overwrite definitions[v.Fl]
		if _, ok := definitions[v.Fl]; !ok {
			definitions[v.Fl] = make(VerbDivider)
		}

		// each data object could have multiple definitions
		for _, def := range v.Def {
			verbDivider := def.Vd

			// Verbs have dividers; nouns, adjectives do not
			// Use "dummy" as key name for these cases to make the
			// parsing below much more simple and readable
			if verbDivider == "" {
				verbDivider = "dummy"
			}

			// Do not overwrite verb dividers
			if _, ok := definitions[v.Fl][verbDivider]; !ok {
				definitions[v.Fl][verbDivider] = make(SenseSequences)
			}

			for sn, sseq := range def.Sseq {
				// Each index of sseq is a group of senses
				// At each index is an array of the senses within that group

				// Do not overwrite sense sequences
				if _, ok := definitions[v.Fl][verbDivider][sn]; !ok {
					definitions[v.Fl][verbDivider][sn] = make(Senses, 3)
				}

				for _, sense := range sseq {
					if sense[1].Dt == nil {
						continue
					}
					if sense[1].Dt[0][0] != "text" {
						continue
					}

					// Only capturing the first index of dt, which technically can have many results
					dt := sense[1].Dt[0][1]
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
	nouns := definitions["noun"]

	for _, sequence := range nouns["dummy"] {
		for _, sense := range sequence {
			fmt.Println(formatSenseText(sense))
		}
	}
}
