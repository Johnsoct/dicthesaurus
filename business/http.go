package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Johnsoct/dicthesaurus/repository"
)

var Data []repository.DictionaryAPIFound

func request() []byte {
	fmt.Fprintf(os.Stdout, "\nSearching for \"%s\" ... \n\n", LookupValue)

	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + LookupValue)
	if err != nil {
		// In case of panicking goroutine: terminates request(), reports error
		panic(err)
	}
	defer resp.Body.Close()

	// Check for 404 in case of a word not being found
	if resp.Status == "404 Not Found" {
		fmt.Fprintf(os.Stderr, "Sorry, a definition for %s was not found. Feel free to try again.\n", LookupValue)
		os.Exit(1)
	}

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
