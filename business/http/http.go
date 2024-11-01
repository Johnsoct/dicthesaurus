package business

import (
	"fmt"
	"os"

	business "github.com/Johnsoct/dicthesaurus/business/input"
)

func Search() {
	if business.LookupValue == "" {
		fmt.Fprintf(os.Stdout, "lookup value is nil")
		return
	}

	fmt.Fprintf(os.Stdout, "Searching via net/http")
}
