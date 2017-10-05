package errors

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	fmt.Println("Error", err)
	os.Exit(1)
}
