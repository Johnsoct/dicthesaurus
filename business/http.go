package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type definitions struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
	Synonyms   []any  `json:"synonyms"`
	Antonyms   []any  `json:"antonyms"`
}

type meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []definitions `json:"definitions"`
}

type phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio,omitempty"`
}

type dictionaryapi struct {
	Word      string      `json:"word"`
	Phonetic  string      `json:"phonetic"`
	Phonetics []phonetics `json:"phonetics"`
	Origin    string      `json:"origin"`
	Meanings  []meanings  `json:"meanings"`
}

var Data []dictionaryapi

// TODO: handle case where no results are found
func request() []byte {
	fmt.Fprintf(os.Stdout, "\nSearching for \"%s\" ... \n\n", LookupValue)

	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + LookupValue)
	if err != nil {
		// In case of panicking goroutine: terminates request(), reports error
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body into a []byte, err (JSON is all one line)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed reading response body into []byte: %v", err)
		os.Exit(1)
	}

	return body
}

func Unmarshal() {
	body := request()

	err := json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decoding the data: %v", err)
		os.Exit(1)
	}
}
