package presentation

import (
	"fmt"
	"os"
	"strings"

	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/repository"
	"github.com/Johnsoct/dicthesaurus/utils"
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

func formatAntsSyns(s []string) string {
	rowFormat := utils.FormatTableRow(s)
	return fmt.Sprintf(rowFormat, utils.ConvertSliceBuiltInTypeToSliceAny(s)...)
}

func formatSenseText(text string) string {
	// Replace before stripping... hehe (replace relies on the tokens)
	replaced := utils.FormatValueBetweenTokens(text)
	stripped := utils.StripTokens(replaced)

	return stripped
}

func formatSequence(sseqn int, sn int, text string) string {
	// Do not prent the sseq number for every sense
	if sn == 0 {
		return fmt.Sprintf("\n%d\t%s", sseqn+1, formatSenseText(text))
	}
	return fmt.Sprintf("\t%s", utils.FormatRow(formatSenseText(text)))
}

// DEFINITIONS
// DEFINITIONS
// DEFINITIONS

func defineVd(vd string, fl string) string {
	verbDivider := vd

	// Verbs have dividers; nouns, adjectives do not
	// Use Fl as key name for these cases to make the
	// parsing below much more simple and readable
	if verbDivider == "" {
		verbDivider = fl
	}

	return verbDivider
}

func excludeDefinitions(v repository.MWResult) bool {
	if excludeEmptyDefinition(v) || excludeStemMatch(v) {
		return true
	}

	return false
}

func excludeEmptyDefinition(v repository.MWResult) bool {
	// If the data object doesn't have the property "hom,"
	// it's not an identical spelling as the searched word
	if v.Def == nil {
		return true
	}

	return false
}

func excludeStemMatch(v repository.MWResult) bool {
	// If Meta ID does not == [word] or [word:#], it's a stem off of [word]
	directMatch := v.Meta.ID == strings.ToLower(business.ParseSubcmd(os.Args))
	prefixMatch := strings.HasPrefix(v.Meta.ID, strings.ToLower(business.ParseSubcmd(os.Args)+":"))
	if !directMatch && !prefixMatch {
		return true
	}

	return false
}

func setDefVd(definitions VerbDivider, verb string) VerbDivider {
	vd := definitions

	// Do not overwrite verb dividers; will later append to definitions[v.Fl][verbDivider][sn]
	if _, ok := vd[verb]; !ok {
		vd[verb] = make(SenseSequences)
	}

	return vd
}

func setDefFl(definitions Definitions, fl string) Definitions {
	def := definitions

	// Do not overwrite definitions[v.Fl]; will later append to definitions[v.Fl][verbDivider][sn]
	if _, ok := def[fl]; !ok {
		def[fl] = make(VerbDivider)
	}

	return def
}

func setVdSseq(definitions SenseSequences, v repository.MWResult, def repository.MWDef) SenseSequences {
	sseq := definitions

	// Short and sweet response
	if *business.SSFlag {
		sseq[0] = v.Shortdef
	} else {
		sseq = setSseq(definitions, def)
	}

	return sseq
}

func setSseq(vd SenseSequences, def repository.MWDef) SenseSequences {
	sns := vd

	// Each index of sseq is a seq, which contains the actual senses
	// "sn" is simply the index to keep track of the number of seq's for a tiered output
	for sn, sseq := range def.Sseq {
		// Do not overwrite sense sequences; will later append to
		if _, ok := sns[sn]; !ok {
			sns[sn] = make(Senses, 0)
		}

		sns[sn] = setSseqSns(sseq)

	}

	return sns
}

func setSseqSns(sseq [][]repository.MWSseq) Senses {
	senses := make(Senses, 0)

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
			senses = append(senses, d)
		}
	}

	return senses
}

func prepareDefinitions(data []repository.MWResult) Definitions {
	definitions := make(Definitions)

	// data could have multiple results
	for _, v := range data {
		if excl := excludeDefinitions(v); excl {
			continue
		}

		definitions = setDefFl(definitions, v.Fl)

		// each data object could have multiple definitions
		for _, def := range v.Def {
			verbDivider := defineVd(def.Vd, v.Fl)
			definitions[v.Fl] = setDefVd(definitions[v.Fl], verbDivider)
			definitions[v.Fl][verbDivider] = setVdSseq(definitions[v.Fl][verbDivider], v, def)

		}
	}

	return definitions
}

func printDictionaryFl(definitions Definitions, fl string) {
	for divider, value := range definitions[fl] {
		fmt.Println(utils.FormatHeading(divider))
		printDictionarySeq(value)
	}
}

func printDictionarySeq(value SenseSequences) {
	// Iterate through the sequences in order (matches Merriam-Webster's results)
	for i := range len(value) {
		for sn, sense := range value[i] {
			fmt.Println(formatSequence(i, sn, sense))
		}
	}
}

func printDictionary(data []repository.MWResult) {
	// Example of verbs: definitions["verb"]["intransitive verb"]
	definitions := prepareDefinitions(data)

	for fl := range definitions {
		printDictionaryFl(definitions, fl)
	}
}

// THESAURUS
// THESAURUS
// THESAURUS

func defineThValues(v repository.MWResult, which string) [][]string {
	values := v.Meta.Syns
	if which == "antonyms" {
		values = v.Meta.Ants
	}

	return values
}

func filterThStems(data []repository.MWResult) []repository.MWResult {
	filteredData := make([]repository.MWResult, len(data))

	for i, v := range data {
		// Ignore all the nyms for stems branching away from the SUBCOMMAND
		if v.Meta.ID == business.ParseSubcmd(os.Args) {
			filteredData = append(filteredData, data[i])
		}
	}

	return filteredData
}

func setThFl(nyms Thesaurus, fl string, values [][]string) Thesaurus {
	nymsFl := nyms

	// Do not overwrite nyms[v.Fl]
	if _, ok := nymsFl[fl]; !ok {
		nymsFl[fl] = make([][]string, len(values))
	}

	return nymsFl
}

func setThFlValues(nyms [][]string, values [][]string) [][]string {
	nymsFl := nyms

	for i, val := range values {
		nymsFl[i] = val
	}

	return nymsFl
}

func prepareTh(data []repository.MWResult, whichType string) Thesaurus {
	filteredData := filterThStems(data)
	nyms := make(Thesaurus)

	for _, v := range filteredData {
		values := defineThValues(v, whichType)

		if len(values) == 0 {
			continue
		}

		nyms = setThFl(nyms, v.Fl, values)
		nyms[v.Fl] = setThFlValues(nyms[v.Fl], values)
	}

	return nyms
}

func printThRows(row []string) {
	fmt.Println(utils.FormatRow(formatAntsSyns(row)))
}

func printThNyms(fl string, nyms [][]string) {
	if len(nyms) == 0 {
		return
	}

	fmt.Println(utils.FormatHeading(fl))

	for _, row := range nyms {
		printThRows(row)
	}
}

func printThesaurus(data []repository.MWResult) {
	synonyms := prepareTh(data, "synonyms")
	antonyms := prepareTh(data, "antonyms")

	if len(synonyms) == 0 && len(antonyms) == 0 {
		fmt.Println(utils.FormatRow("There are no results for that word. Most likely, the word is a stem (version) of a word and has too many possible matches"))
	}

	if len(synonyms) > 0 {
		fmt.Println(utils.FormatHeader("synonyms"))
	}
	for fl, nyms := range synonyms {
		printThNyms(fl, nyms)
	}

	if len(antonyms) > 0 {
		fmt.Println(utils.FormatHeader("antonyms"))
	}
	for fl, nyms := range antonyms {
		printThNyms(fl, nyms)
	}
}

func Print(data []repository.MWResult, which string) {
	if data == nil {
		fmt.Printf("\n\t%s\n\n", "The API response was empty. Please check your API keys.")
		os.Exit(1)
	}

	switch which {
	case "dictionary":
		printDictionary(data)
	case "thesaurus":
		printThesaurus(data)
	}

	fmt.Println()
}
