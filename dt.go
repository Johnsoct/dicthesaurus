// user/bin/go run $0 $@ ; exit

// English dictionary and thesaurus lookup CLI tool
package main

import (
	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/presentation"
)

func main() {
	lookupValue := business.GetSubCommand()
	response := business.GetDefintionInBytes(lookupValue)
	data := business.UnmarshaledJSON(response)
	presentation.Print(data)
}
