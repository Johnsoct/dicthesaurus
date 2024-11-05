package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

func convertResponseToBytes(response *http.Response) []byte {
	// Read the response body into a []byte, err (JSON is all one line)
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading response body into []byte: %v", err)
		os.Exit(1)
	}

	return bytes
}

func GetDefinition(word string) []byte {
	fmt.Fprintf(os.Stdout, "\nSearching for \"%s\" ... \n\n", word)

	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
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

	return convertResponseToBytes(resp)
}

func UnmarshalResponse(jsonBytes []byte) []repository.DictionaryAPIFound {
	var data []repository.DictionaryAPIFound

	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decoding the data: %v", err)
		os.Exit(1)
	}

	return data
}
