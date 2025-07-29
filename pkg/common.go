package pkg

import (
	"fmt"
	"os"
)


func ExitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}