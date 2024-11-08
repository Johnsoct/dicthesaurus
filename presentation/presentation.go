package presentation

import (
	"fmt"
	"os"
	"strings"

	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/repository"
)

func boldText(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func prepareDefinition(definition repository.Definitions) string {
	output := fmt.Sprintf("%s\n", definition.Definition) // Defaults to just the definition declaration

	if *business.EFlag && definition.Example != "" {
		output += fmt.Sprintf("Used in a sentence: %s\n", definition.Example)
	}

	if len(definition.Synonyms) != 0 {
		output += fmt.Sprintf("Synonyms: %s\n", strings.Join(definition.Synonyms, " "))
	}

	if len(definition.Antonyms) != 0 {
		output += fmt.Sprintf("Antonyms: %s", strings.Join(definition.Antonyms, " "))
	}

	return fmt.Sprintf("%s\n", output)
}

func prepareMeaning(meaning repository.Meanings) string {
	partOfSpeech := meaning.PartOfSpeech
	definitions := make([]string, 5)

	for _, definition := range meaning.Definitions {
		definitions = append(definitions, prepareDefinition(definition))
	}

	return fmt.Sprintf(
		"%s\n\n%s",
		boldText(strings.ToUpper(partOfSpeech)),
		strings.Join(definitions, ""),
	)
}

func prepareMeanings(data []repository.DictionaryAPIFound) []string {
	meanings := make([]string, 5)

	// data could have multiple results
	for _, v := range data {
		// each data object could have multiple meanings
		for _, meaning := range v.Meanings {
			meanings = append(meanings, prepareMeaning(meaning))
		}
	}

	return meanings
}

func Print(data []repository.DictionaryAPIFound) {
	if data == nil {
		fmt.Println("I'm not sure how you got here, but something is wrong. Sorry, try again.")
		os.Exit(1)
	}

	fmt.Printf("%v", strings.Join(prepareMeanings(data), ""))
}
