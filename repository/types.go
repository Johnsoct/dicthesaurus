package repository

// Constants data structures
type Flag struct {
	Flag        string
	Description string
}

// API data structures

type MWDef struct {
	// Verb divider
	Vd string `json:"vd,omitempty"`
	// Group of all the sense sequences and verb dividers for each headword or defined run-on phrase
	// Sense sequence = contains series of senses and subsenses, ordered by sense numbers
	Sseq [][][]MWSenseSequence `json:"sseq"`
}

type MWDt struct {
	// Defining text of a particular sense
	Dt [][]string `json:"dt"`
}

type MWHwi struct {
	// Headword
	Hw string `json:"hw"`
	// Pronunciations
	Prs []MWPrs `json:"prs"`
}

type MWMeta struct {
	ID   string `json:"id"`
	UUID string `json:"uuid"`
	// 9-digit code used to sort entries if alphabetical order is needed
	Sort string `json:"sort"`
	// Internal data set for entry -- ignore
	Src string `json:"src"`
	// The section the entry belongs to (print, biographical, geographical, foreign words, phrases)
	Section string `json:"section"`
	// List of entry's headwords, variants, inflections, undefined entry words, and defined run-on phrases
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}

type MWPrs struct {
	// Merriam-webster written format
	Mw    string  `json:"mw"`
	Sound MWSound `json:"sound"`
}

type MWSense struct {
	// Sense number
	Sn string `json:"sn"`
	//
	Dt [][]any `json:"dt"`
}

type MWSenseSequence struct {
	string
	MWSense
}

type MWSound struct {
	Audio string `json:"audio"`
	Ref   string `json:"ref"`
	Stat  string `json:"stat"`
}

// Root Structs

type MWDResult struct {
	Date string `json:"date"`
	// Headword information
	Def []MWDef `json:"def"`
	// Etymology
	Et [][]string `json:"et"`
	// Functional label - grammatical function of a headword
	Fl string `json:"fl"`
	// Headwords with identical spelling but distint meanings
	// Marked by an integer to distinguish between identically-spelled headwords
	Hom int   `json:"hom"`
	Hwi MWHwi `json:"hwi"`
	// Definition
	Meta MWMeta `json:"meta"`
	// Highly abridged version of the Definition section consisting of the first three senses
	// definitions
	Shortdef []string `json:"shortdef"`
}

type MWTResult struct{}
