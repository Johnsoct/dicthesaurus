package business

import (
	"fmt"
	"os"

	input "github.com/Johnsoct/dicthesaurus/business/input"
)

func Search() {
	if input.LookupValue == "" {
		fmt.Fprintf(os.Stdout, "lookup value is nil")
		return
	}

	fmt.Fprintf(os.Stdout, "Searching via net/http")
}
