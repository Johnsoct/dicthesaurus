package business

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// type JSON interface {
//         get()
//         TODO: return???()
// }
//
// TODO: string method?

type dictionaryapi struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	} `json:"phonetics"`
	Origin   string `json:"origin"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
		} `json:"definitions"`
	} `json:"meanings"`
}

// dictionaryapi.dev returns an array of *dictionaryapi
// type JSON map[string]dictionaryapi

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
	fmt.Println("test:", data.Word])
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
