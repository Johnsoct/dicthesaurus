package business

import (
	"fmt"
	"os"
)

func Search() {
	if LookupValue == "" {
		fmt.Fprintf(os.Stdout, "lookup value is nil")
		return
	}

	fmt.Fprintf(os.Stdout, "Searching via net/http")
}
