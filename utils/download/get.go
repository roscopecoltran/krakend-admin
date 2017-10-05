package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetFile(filepath string, url string) error {

	if url == "" {
		return errors.New("empty remote url.")
	}

	if filepath == "" {
		return errors.New("no destination path provided.")
	}

	prefixPath, err := GetPrefixPath(filepath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(prefixPath, 0700); err != nil {
		return err
	}

	if cached := FileExists(filepath); cached == true {
		fmt.Printf("skipping content to download as already present locally, url: \"%s\", dest: \"%s\" \n", url, filepath)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("successfully downloaded content, url: \"%s\", dest: \"%s\" \n", url, filepath)

	return nil
}
