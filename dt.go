// user/bin/go run $0 $@ ; exit

// English dictionary and thesaurus lookup CLI tool
package main

import (
	"log"
	"os"

	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/presentation"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	if v := *business.TFlag; v {
		data := business.GetThesaurus(business.ParseSubcmd(os.Args))
		presentation.Print(data, "thesaurus")
	} else {
		data := business.GetDefinition(business.ParseSubcmd(os.Args))
		presentation.Print(data, "dictionary")
	}
}
