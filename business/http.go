package business

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Not sure what to expect from the response's JSON data
type JSON map[string]interface{}

func Search() {
	if LookupValue == "" {
		fmt.Fprintf(os.Stdout, "lookup value is nil")
		return
	}

	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + LookupValue)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Body)

	var data JSON
	decoder := json.NewDecoder(resp.Body)

	// read open bracket
	t, err := decoder.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// While the array contains values
	for decoder.More() {
		// decode an array value
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", data)
	}

	// read closing brakcet
	t, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T: %v\n", t, t)

	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
	// 	fmt.Println(scanner.Bytes())
	// 	err = json.Unmarshal(scanner.Bytes(), &data)
	// }
	//
	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }

	fmt.Println(data)
}
