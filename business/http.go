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
		fmt.Println(err)
	}
	fmt.Println("Unmarshalled data:", data)

	fmt.Println("test:", data)
}

func Search() {
	if LookupValue == "" {
		fmt.Fprintf(os.Stdout, "lookup value is nil")
		return
	}

	resp, _ := request()
	decoder := json.NewDecoder(resp.Body)

	// read open bracket
	_, err := decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	// While the array (json [{...}]) contains values
	for decoder.More() {
		// decode an array value
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", data)
	}

	// read closing brakcet
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}
}
