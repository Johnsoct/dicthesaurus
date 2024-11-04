package repository

type Definitions struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example,omitempty"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definitions `json:"definitions"`
}

type phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio,omitempty"`
}

type DictionaryAPIFound struct {
	Word      string      `json:"word"`
	Phonetic  string      `json:"phonetic"`
	Phonetics []phonetics `json:"phonetics"`
	Origin    string      `json:"origin"`
	Meanings  []Meanings  `json:"meanings"`
}
