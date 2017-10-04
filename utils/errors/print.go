package errors

import (
	"fmt"
	"os"
)

func printError(err error) {
	fmt.Println("Error", err)
	os.Exit(1)
}
