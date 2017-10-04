package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func getFilePathFromURL(remoteURL string) (string, string, string) {
	u, err := url.Parse(remoteURL)
	if err != nil {
		fmt.Printf("Invalid URL %s, skipping.\n", remoteURL)
		return "", "", ""
	}
	ext := path.Ext(u.Path)
	// fmt.Printf("ext: %s\n", ext)
	if *isSanitizeUrl {
		sanitizedName := sanitizeUrl(u.Scheme + u.Host + trimSuffix(u.Path, ext) + "?" + u.RawQuery)
		fmt.Printf("sanitizedName: %s\n", sanitizedName)
	}
	filePath, err := GetPrefixPath(u.Path)
	if err != nil {
		fmt.Println("Error while getting the destination file prefix path, skipping:", err)
		return "", "", ""
	}

	fileName := strings.TrimRight(u.Path, "\n") // filePath + ext
	fullPath := strings.TrimRight(fmt.Sprintf("%s%s", *downloadDir, filePath), "\n")
	fullName := strings.TrimRight(fmt.Sprintf("%s%s", *downloadDir, fileName), "\n") // parentDir + fileName
	return fileName, fullPath, fullName
}

func GetPrefixPath(filePath string) (string, error) {
	prefixPath, err := filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return "", err
	}
	return prefixPath, nil
}

// Exists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func downloadFile(filepath string, url string) error {

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

func worker(work chan string, wg *sync.WaitGroup) {
	for workUrl := range work {
		downloadUrl(workUrl)
	}
	wg.Done()
}

func downloadUrl(workUrl string) {
	u, err := url.Parse(workUrl)
	if err != nil {
		fmt.Printf("Invalid URL %s, skipping.\n", workUrl)
		return
	}
	//parentDir := downloadDir
	ext := path.Ext(u.Path)
	// fmt.Printf("ext: %s\n", ext)
	if *isSanitizeUrl {
		sanitizedName := sanitizeUrl(u.Scheme + u.Host + trimSuffix(u.Path, ext) + "?" + u.RawQuery)
		fmt.Printf("sanitizedName: %s\n", sanitizedName)
	}
	// filePath := u.Path + "/"

	filePath, err := GetPrefixPath(u.Path)
	if err != nil {
		fmt.Println("Error while getting the destination file prefix path, skipping:", err)
		return
	}

	fileName := u.Path // filePath + ext
	fullPath := fmt.Sprintf("%v%s", *downloadDir, filePath)
	fullName := fmt.Sprintf("%v%s", *downloadDir, fileName) // parentDir + fileName

	err = os.MkdirAll(fullPath, 0700)
	if err != nil {
		fmt.Println("Error creating directory, skipping:", err)
		return
	}

	_, err = os.Stat(fullName)
	if err == nil {
		fmt.Printf("File %s already exists, skipping\n", fileName)
		return
	}

	out, err := os.Create(fullName)
	if err != nil {
		fmt.Println("Error creating file, skipping:", err)
		return
	}
	defer out.Close()

	resp, err := http.Get(workUrl)
	if err != nil {
		fmt.Println("Error getting URL, skipping:", err)
		return
	}
	defer resp.Body.Close()

	bytes, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing file, skipping:", err)
		return
	}

	fmt.Printf("%d bytes: %s\n", bytes, fileName)
}

func sanitizeUrl(url string) string {
	return regexp.MustCompile("\\s+").ReplaceAllLiteralString(stripChars(url), "-")
}

func stripChars(str string) string {
	toStrip := [...]string{
		"~", "`", "!", "@", "#", "$", "%", "^", "&", "*",
		"(", ")", "_", "=", "+", "[", "{", "]", "}", "\\",
		"|", ";", ":", "\"", "'", "&#8216;", "&#8217;",
		"&#8220;", "&#8221;", "&#8211;", "&#8212;", "—",
		"–", ",", "<", ".", ">", "/", "?",
	}

	for _, char := range toStrip {
		str = strings.Replace(str, char, "", -1)
	}

	return str
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}

	return s
}

func addWorkFromCsv(work chan string, filePath string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		work <- line[0]
	}
}
