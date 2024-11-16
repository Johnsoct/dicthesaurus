package repository

type MWSense struct {
	// Sense number
	Sn string `json:"sn"`
	//
	Dt [][]any `json:"dt"`
}

type MWResult struct {
	Date string `json:"date"`
	// Headword information
	Def []struct {
		// Verb divider
		Vd string `json:"vd,omitempty"`
		// Group of all the sense sequences and verb dividers for each headword or defined run-on phrase
		// Sense sequence = contains series of senses and subsenses, ordered by sense numbers
		Sseq [][][]struct {
			string
			MWSense
		} `json:"sseq"`
	} `json:"def"`
	// Etymology
	Et [][]string `json:"et"`
	// Functional label - grammatical function of a headword
	Fl string `json:"fl"`
	// Headwords with identical spelling but distint meanings
	// Marked by an integer to distinguish between identically-spelled headwords
	Hom int `json:"hom"`
	Hwi struct {
		// Headword
		Hw string `json:"hw"`
		// Pronunciations
		Prs []struct {
			// Merriam-webster written format
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"hwi"`
	// Definition
	Meta struct {
		ID   string `json:"id"`
		UUID string `json:"uuid"`
		// 9-digit code used to sort entries if alphabetical order is needed
		Sort string `json:"sort"`
		// Internal data set for entry -- ignore
		Src string `json:"src"`
		// The section the entry belongs to (print, biographical, geographical, foreign words, phrases)
		Section string `json:"section"`
		// List of entry's headwords, variants, inflections, undefined entry words, and defined run-on phrases
		Stems []string `json:"stems"`
		// List []Synonyms
		Syns      [][]string `json:"syns"`
		Ants      [][]string `json:"ants"`
		Offensive bool       `json:"offensive"`
	} `json:"meta"`
	// Highly abridged version of the Definition section consisting of the first three senses
	// definitions
	Shortdef []string `json:"shortdef"`
}
