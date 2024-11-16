package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

func getMeriamWebster(word, endpoint string) []repository.MWResult {
	API_KEY := "MERRIAM_WEBSTER_DICTIONARY_API_KEY"
	if endpoint == "thesaurus" {
		API_KEY = "MERRIAM_WEBSTER_THESAURUS_API_KEY"
	}

	resp, respErr := http.Get("https://www.dictionaryapi.com/api/v3/references/" + endpoint + "/json/" + word + "?key=" + os.Getenv(API_KEY))
	if respErr != nil {
		// In case of panicking goroutine: terminates request(), reports error
		panic(respErr)
	}
	defer resp.Body.Close()

	// Check for 404 in case of a word not being found
	if resp.Status == "404 Not Found" {
		fmt.Fprintf(os.Stderr, "Sorry, a definition for %s was not found. Feel free to try again.\n", word)
		os.Exit(1)
	}

	var data []repository.MWResult

	decodeErr := json.NewDecoder(resp.Body).Decode(&data)
	if decodeErr != nil {
		// TODO: handle Decode not liking decoding a tuple([string, {...}]) into MWResult.Def.Sseq
		// Most likely need a custom unmarshal function
		// fmt.Printf("error decoding JSON from Meriam-Webster %v", decodeErr)
	}

	return data
}

func GetDefinition(word string) []repository.MWResult {
	return getMeriamWebster(word, "collegiate")
}

func GetThesaurus(word string) []repository.MWResult {
	return getMeriamWebster(word, "thesaurus")
}
