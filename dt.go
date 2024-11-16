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

	switch business.ParseEndpoint(business.SubcmdFlags) {
	case "dictionary":
		data := business.GetDefinition(business.ParseSubcmd(os.Args))
		presentation.Print(data, "dictionary")
	case "thesaurus":
		data := business.GetThesaurus(business.ParseSubcmd(os.Args))
		presentation.Print(data, "thesaurus")
	}
}
