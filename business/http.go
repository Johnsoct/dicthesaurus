package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

func handle404Error(word string) {
	fmt.Fprintf(os.Stderr, "Sorry, a definition for %s was not found. Feel free to try again.\n", word)
	os.Exit(1)
}

func handleDecodingError(err error) {
	// TODO: handle Decode not liking decoding a tuple([string, {...}]) into MWResult.Def.Sseq
	// Most likely need a custom unmarshal function
	// fmt.Printf("error decoding JSON from Meriam-Webster %v", decodeErr)
}

func handleGetErrors(err error) {
	// In case of panicking goroutine: terminates request(), reports error
	panic(err)
}

func get(word, endpoint, key string) []repository.MWResult {
	resp, respErr := http.Get("https://www.dictionaryapi.com/api/v3/references/" + endpoint + "/json/" + word + "?key=" + os.Getenv(key))
	if respErr != nil {
		handleGetErrors(respErr)
	}
	defer resp.Body.Close()

	// Check for 404 in case of a word not being found
	if resp.Status == "404 Not Found" {
		handle404Error(word)
	}

	var data []repository.MWResult

	decodeErr := json.NewDecoder(resp.Body).Decode(&data)
	if decodeErr != nil {
		handleDecodingError(decodeErr)
	}

	return data
}

func GetDefinition(word string) []repository.MWResult {
	return get(word, "collegiate", "MERRIAM_WEBSTER_DICTIONARY_API_KEY")
}

func GetThesaurus(word string) []repository.MWResult {
	return get(word, "thesaurus", "MERRIAM_WEBSTER_THESAURUS_API_KEY")
}
