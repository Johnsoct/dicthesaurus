// English dictionary and thesaurus lookup CLI tool
package main

import (
	"github.com/Johnsoct/dicthesaurus/business"
	"github.com/Johnsoct/dicthesaurus/presentation"
)

//"github.com/Johnsoct/dicthesaurus/presentation"
// "github.com/Johnsoct/dicthesaurus/repository"

func main() {
	business.Unmarshal()
	presentation.Output()
}
