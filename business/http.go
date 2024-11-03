package business

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

var data []dictionaryapi

// TODO: better error handling
func request() (*http.Response, []byte) {
	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + LookupValue)
	if err != nil {
		// TODO: wtf does panic do
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resp, body
}

// TODO: better error handling
func Unmarshal() {
	_, body := request()

	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decoding the data: %v", err)
	}

	fmt.Println("Unmarshaled data:", data[0])
}
