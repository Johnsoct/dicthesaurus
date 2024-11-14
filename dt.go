// user/bin/go run $0 $@ ; exit

// English dictionary and thesaurus lookup CLI tool
package main

import (
	"log"

	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/presentation"
	"github.com/Johnsoct/dicthesaurus/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	data := business.GetDefinition(repository.SUBCOMMAND)
	presentation.Print(data)
}
