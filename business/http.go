package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

type MWResponse interface {
	[]repository.MWDResult | []repository.MWTResult
}

func getMeriamWebster[T MWResponse](word, endpoint string) T {
	API_KEY := "MERRIAM_WEBSTER_DICTIONARY_API_KEY"
	if endpoint == "thesaurus" {
		API_KEY = "MERRIAM_WEBSTER_THESAURUS_API_KEY"
	}

	resp, err := http.Get("https://www.dictionaryapi.com/api/v3/references/" + endpoint + "/json/" + word + "?key=" + os.Getenv(API_KEY))
	if err != nil {
		// In case of panicking goroutine: terminates request(), reports error
		panic(err)
	}
	defer resp.Body.Close()

	// Check for 404 in case of a word not being found
	if resp.Status == "404 Not Found" {
		fmt.Fprintf(os.Stderr, "Sorry, a definition for %s was not found. Feel free to try again.\n", word)
		os.Exit(1)
	}

	var data T

	json.NewDecoder(resp.Body).Decode(&data)

	return data
}

func GetDefinition(word string) []repository.MWDResult {
	return getMeriamWebster[[]repository.MWDResult](word, "collegiate")
}

func GetThesaurus(word string) []repository.MWTResult {
	return getMeriamWebster[[]repository.MWTResult](word, "thesaurus")
}
