// user/bin/go run $0 $@ ; exit

// English dictionary and thesaurus lookup CLI tool
package main

import (
	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/presentation"
	"github.com/Johnsoct/dicthesaurus/repository"
)

func main() {
	data := business.GetDefinition(repository.SUBCOMMAND)
	presentation.Print(data)
}
