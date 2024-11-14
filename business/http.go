package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

func GetDefinition(word string) []repository.MWDResult {
	resp, err := http.Get("https://www.dictionaryapi.com/api/v3/references/collegiate/json/" + word + "?key=" + os.Getenv("MERRIAM_WEBSTER_DICTIONARY_API_KEY"))
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

	var data []repository.MWDResult
	json.NewDecoder(resp.Body).Decode(&data)

	return data
}
